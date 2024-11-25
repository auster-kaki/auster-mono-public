package usecase

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/app/rpc"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/auster-kaki/auster-mono/pkg/util/austerid"
	"github.com/auster-kaki/auster-mono/pkg/util/austerstorage"
)

type TravelSpotUseCase struct {
	repository repository.Repository
	rpc        rpc.RPC
}

func NewTravelSpotUseCase(r repository.Repository, rpc rpc.RPC) *TravelSpotUseCase {
	return &TravelSpotUseCase{repository: r, rpc: rpc}
}

func (u *TravelSpotUseCase) GetTravelSpots(ctx context.Context, userID entity.UserID, hobbyID entity.HobbyID) (entity.TravelSpots, error) {
	travelSpotHobbies, err := u.repository.TravelSpotHobby().GetByHobbyID(ctx, hobbyID)
	if err != nil {
		return nil, err
	}
	travelSpots, err := u.repository.TravelSpot().GetByIDs(ctx, travelSpotHobbies.TravelSpotIDs())
	if err != nil {
		return nil, err
	}
	rs, err := u.repository.Reservation().GetEndedReservations(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 体験のレベルに応じて表示する体験を絞り込む
	// レベル1: 0回以上の体験したことがある
	// レベル2: 3回以上の体験したことがある
	// レベル3: 5回以上の体験したことがある
	out := make(entity.TravelSpots, 0)
	for _, travelSpot := range travelSpots {
		// NOTE: オファー体験は表示しない
		if travelSpot.Level == entity.TravelSpotForOffer {
			continue
		}
		if (travelSpot.Level == entity.TravelSpotLevel1) ||
			(travelSpot.Level == entity.TravelSpotLevel2 && len(rs) >= 3) ||
			(travelSpot.Level == entity.TravelSpotLevel3 && len(rs) >= 5) {
			out = append(out, travelSpot)
		}
	}
	slices.SortFunc(out, func(a, b *entity.TravelSpot) int {
		return int(b.Level) - int(a.Level)
	})
	if len(out) > 4 {
		out = out[:4]
	}
	return out, nil
}

type CreateDiaryOutput struct {
	ID          entity.TravelSpotDiaryID
	Title       string
	PhotoPath   string
	Description string
}

// CreateDiary TODO: 画像・本文・タイトルの生成周りが未完成
func (u *TravelSpotUseCase) CreateDiary(ctx context.Context, userID entity.UserID, travelSpotID entity.TravelSpotID, date time.Time) (*CreateDiaryOutput, error) {
	user, err := u.repository.User().FindByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	travelSpot, err := u.repository.TravelSpot().FindByID(ctx, travelSpotID)
	if err != nil {
		return nil, fmt.Errorf("failed to find travel spot: %w", err)
	}
	/*
		日記を生成する条件
		1. 初めて体験を選択した場合
		2. 予約後、もう一度体験を選択した場合

		日記を生成しない条件
		1. 初めて体験を選択した後、一度予約をせずに戻り、もう一度同じ体験を選択した場合
	*/
	rs, err := u.repository.Reservation().FindByUserIDAndTravelSpotID(ctx, userID, travelSpotID)
	if err != nil {
		return nil, fmt.Errorf("failed to find reservation: %w", err)
	}

	travelSpotDiaries, dErr := u.repository.TravelSpotDiary().FindByUserIDAndTravelSpotID(ctx, userID, travelSpotID)
	if dErr != nil && !errors.Is(dErr, repository.ErrNotFound) {
		return nil, fmt.Errorf("failed to find travel spot diary: %w", dErr)
	}
	for _, travelSpotDiary := range travelSpotDiaries {
		// 日記を生成しない条件
		// 1. 初めて体験を選択した後、一度予約をせずに戻り、もう一度同じ体験を選択した場合
		if !slices.Contains(rs.TravelSpotDiaryIDs(), travelSpotDiary.ID) {
			return &CreateDiaryOutput{
				ID:          travelSpotDiary.ID,
				Title:       travelSpotDiary.Title,
				PhotoPath:   travelSpotDiary.PhotoPath,
				Description: travelSpotDiary.Description,
			}, nil
		}
	}

	gOut, err := u.generateDiary(ctx, user, travelSpot)
	if err != nil {
		fmt.Println("failed to generate diary: %w", err)
		//　画像生成に失敗した場合は元の体験画像をそのまま返す
		photo, err := austerstorage.Get(travelSpot.PhotoPath)
		if err != nil {
			return nil, fmt.Errorf("failed to get photo: %w", err)
		}
		paths := strings.Split(travelSpot.PhotoPath, "/")
		gOut = &rpc.GetImagePathOutput{
			GeneratedImage: photo,
			Filename:       paths[len(paths)-1],
		}
	}
	id := austerid.Generate[entity.TravelSpotDiaryID]()
	// goサーバにも画像を保存
	path, err := austerstorage.Save(
		austerstorage.PNG,
		filepath.Join("travel_spot_diaries", string(userID), string(id), gOut.Filename),
		gOut.GeneratedImage,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to save image: %w", err)
	}

	var (
		title       = "ダミータイトル１"
		description = "ダミー本文１"
	)

	// TODO: テキスト生成APIを呼び出す
	// 調整しましたが出力が安定しないのでダミーデータを使うことにしました
	// userHobbies, err := u.repository.UserHobby().GetByUserID(ctx, user.ID)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get user hobby: %w", err)
	// }

	// hobbies, err := u.repository.Hobby().GetByIDs(ctx, userHobbies.HobbyIDs())
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get hobbies: %w", err)
	// }
	// client := infraHTTP.NewPerplexityClient()
	// out, err := client.CreateDiaryText(&infraHTTP.CreateDiaryTextInput{
	// 	TravelSpotTitle:       travelSpot.Title,
	// 	TravelSpotDescription: travelSpot.Description,
	// 	Hobby:                 strings.Join(hobbies.Names(), "、"),
	// 	UserAge:               user.Age,
	// 	UserGender:            user.Gender,
	// })

	if d, ok := dummyDiaryTexts[travelSpotID]; ok {
		title = d.Title
		description = d.Description
	}

	if err := u.repository.TravelSpotDiary().Create(ctx, &entity.TravelSpotDiary{
		ID:           id,
		UserID:       userID,
		TravelSpotID: travelSpotID,
		Title:        title,
		Date:         date,
		PhotoPath:    path,
		Description:  description,
	}); err != nil {
		return nil, fmt.Errorf("failed to create travel spot diary: %w", err)
	}

	return &CreateDiaryOutput{
		ID:          id,
		Title:       title,
		PhotoPath:   path,
		Description: description,
	}, nil
}

func (u *TravelSpotUseCase) generateDiary(ctx context.Context, user *entity.User, travelSpot *entity.TravelSpot) (*rpc.GetImagePathOutput, error) {
	cOut, err := u.rpc.Diary().CreateImage(ctx, rpc.CreateImageInput{
		SourcePath: travelSpot.PhotoPath,
		TargetPath: user.ProfilePath,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create image: %w", err)
	}

	// ポーリング処理の設定
	var (
		maxRetries   = 90              // 最大リトライ回数
		pollInterval = time.Second * 2 // ポーリング間隔
	)
	// ジョブの完了を待つポーリング処理
	for i := 0; i < maxRetries; i++ {
		gOut, err := u.rpc.Diary().GetStatus(ctx, rpc.GetStatusInput{
			JobID: cOut.JobID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get job status: %w", err)
		}
		if gOut.Status == "ok" {
			// 処理完了
			fmt.Println("image generation completed. job id: ", cOut.JobID)
			break
		} else if gOut.Status == "error" {
			return nil, fmt.Errorf("image generation failed: %s", gOut.Status)
		}

		// 最大リトライ回数に達した場合
		if i == maxRetries-1 {
			return nil, fmt.Errorf("timeout: image generation did not complete within %d attempts", maxRetries)
		}

		// 次のポーリングまで待機
		time.Sleep(pollInterval)
	}

	// 処理が完了してから画像の準備が整うまで待機
	time.Sleep(time.Second * 10)
	gOut, err := u.rpc.Diary().GetImagePath(ctx, rpc.GetImagePathInput{
		JobID: cOut.JobID,
	})
	if err != nil {
		return nil, fmt.Errorf("u.rpc.Diary(). failed to get image path: %w", err)
	}
	return &gOut, nil
}

var dummyDiaryTexts = map[entity.TravelSpotID]entity.TravelSpotDiary{
	"1": {
		Title:       "大物ヒラマサとの出会い",
		Description: "今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！",
	},
	"2": {
		Title:       "大物シイラとの戦い",
		Description: "今日は銚子港から漁船に乗り、念願のシイラ釣りに挑戦。朝もやの中、期待と不安を胸に出港。3時間の格闘の末、体長2メートルを超える大物を釣り上げることができた！船長のアドバイスのおかげで、生涯最高の釣果。あの引きの強さは一生忘れられない。記念写真を撮って、達成感に浸った最高の1日だった。",
	},
	"3": {
		Title:       "美しい海を守る休日",
		Description: "銚子の海岸でゴミ拾いボランティアに参加。いつも釣りで楽しませてもらっている海への恩返しだ。2時間で大きなゴミ袋2つが一杯になった。プラスチックごみが特に目立ち、海の環境問題を実感。地元の漁師さんも「きれいな海を守りたい」と一緒に活動してくれて心強かった。次回は釣り仲間も誘って参加したい。",
	},
	"4": {
		Title:       "漁師の仕事を体験",
		Description: "銚子の漁船に乗せてもらい、本物の漁を体験。早朝3時の出港で眠かったけど、大量の魚を網で引き上げる瞬間の興奮で目が覚めた。漁師さんの手際の良さに感動。魚を選別する作業も手伝わせてもらい、プロの技を間近で見られた。海の上での仕事の大変さと、新鮮な魚が食卓に届くまでの過程を知る貴重な経験になった。",
	},
	"5": {
		Title:       "大物マダイとの出会い",
		Description: "銚子の漁船で初めての本格的な鯛釣りに挑戦。朝日が昇る中、船長から仕掛けの付け方や釣り方を丁寧に教わった。潮の流れを読みながら、じっくりと待つこと2時間。突然の大きな引きに興奮！必死の攻防の末、70センチを超える大物マダイを釣り上げることができた。記念撮影の時の達成感は格別だった。",
	},
	"6": {
		Title:       "銀色の太刀魚との対決",
		Description: "銚子の漁船でタチウオ釣りに挑戦。夜明け前の暗い海で、船長から独特の仕掛けと釣り方を教わる。水深100メートルの深場で待つこと1時間、竿が大きく曲がった！銀色に輝く全長1メートルのタチウオを釣り上げた時は感動で声が出なかった。船長も「良い型だ」と太鼓判を押してくれた。",
	},
	"7": {
		Title:       "漁師の想いを届ける仕事",
		Description: "銚子の漁業組合でマーケティングの仕事を始めた。今日は朝市場で漁師さんたちの仕事ぶりを撮影し、SNSで発信。魚のさばき方や、競りの様子、そして漁師さんたちの生の声を動画にまとめた。プレゼン資料作りの経験を活かしながら、漁業の魅力を分かりやすく伝える工夫を重ねている。この仕事で地域に貢献できる喜びを感じる。",
	},
	"10": {
		Title:       "原始の火を求めて",
		Description: "銚子のキャンプ場でブッシュクラフト体験。黄色い革手袋をはめて、火打ち石と火打ち金を使った伝統的な火起こしに挑戦。最初は火花が上手く飛ばず苦戦したが、コツを掴むと小さな火種が誕生！その瞬間の喜びは格別だった。自然の中で原始的な技術を学ぶ面白さを実感。次は自分の道具で挑戦してみたい。",
	},
	"11": {
		Title:       "駅巡りランニング",
		Description: "銚子電鉄の駅を巡るトレインランニングに挑戦。霧がかかる早朝、最初の駅を出発。海からの潮風を感じながら、駅と駅の間を走る。途中、地元の方から応援の声をかけていただき元気をもらう。各駅で記念撮影をしながら、銚子の町並みや自然を新しい視点で発見。汗を流しながらの駅巡りは格別な体験になった。",
	},
	"12": {
		Title:       "森のカクテルマスター",
		Description: "キャンプ場の特別イベントで一日バーテンダーを務めた。麦わら帽子をかぶり、オリジナルカクテルを提供。焚き火の香りとカクテルの香りが絶妙にマッチして、キャンパーたちにも好評だった。シェイカーを振る音が心地よく響き、夕暮れ時には多くの人が集まってきた。自然の中でのバーテンディングは新鮮で楽しい経験になった。",
	},
	"13": {
		Title:       "自然と人を繋ぐ企画",
		Description: "銚子のアウトドアショップで企画担当として初日を迎えた。音楽イベントやバーテンディング、トレッキングなど、様々なアクティビティを提案。スタッフと打ち合わせを重ね、自然の中での新しい体験価値を創造していく。キャンプ好きの経験を活かしながら、地域の魅力を引き出す企画作りに挑戦。充実感のある仕事の始まりだ。",
	},
	"14": {
		Title:       "音楽と自然の出会い",
		Description: "キャンプ場で音楽イベントの企画運営を始めた。地元ミュージシャンとキャンパーが交流できる場を作り、焚き火を囲んでの即興セッションも実現。音響機材のセッティングから、出演者との打ち合わせまで、充実した準備の日々。自然の中での音楽は格別な響きがあり、参加者の笑顔が印象的だった。新しい体験の創造に胸が躍る。",
	},
}

type GetItinerariesOutput struct {
	TravelSpotItineraries entity.TravelSpotItineraries
	Items                 entity.TravelSpotItineraryItems
}

func (u *TravelSpotUseCase) GetItineraries(ctx context.Context, userID entity.UserID, travelSpotID entity.TravelSpotID) (*GetItinerariesOutput, error) {
	travelSpotItineraries, err := u.repository.TravelSpotItinerary().GetByTravelSpotID(ctx, travelSpotID)
	if err != nil {
		return nil, err
	}
	slices.SortFunc(travelSpotItineraries, func(a, b *entity.TravelSpotItinerary) int {
		return a.Order - b.Order
	})

	travelSpotItineraryItems, err := u.repository.TravelSpotItineraryItem().GetByTravelSpotItineraryIDs(ctx, travelSpotItineraries.IDs())
	if err != nil {
		return nil, err
	}

	return &GetItinerariesOutput{
		TravelSpotItineraries: travelSpotItineraries,
		Items:                 travelSpotItineraryItems,
	}, nil
}
