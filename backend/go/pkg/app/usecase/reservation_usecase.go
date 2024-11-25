package usecase

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"path/filepath"
	"slices"
	"time"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/auster-kaki/auster-mono/pkg/util/austerid"
	"github.com/auster-kaki/auster-mono/pkg/util/austerstorage"
)

type ReservationUseCase struct {
	repository repository.Repository
}

func NewReservationUseCase(r repository.Repository) *ReservationUseCase {
	return &ReservationUseCase{repository: r}
}

type CreateReservationInput struct {
	UserID            entity.UserID
	TravelSpotID      entity.TravelSpotID
	TravelSpotDiaryID entity.TravelSpotDiaryID
	From              time.Time
	To                time.Time
}

func (u *ReservationUseCase) Create(ctx context.Context, input *CreateReservationInput) (entity.ReservationID, error) {
	id := austerid.Generate[entity.ReservationID]()
	if err := u.repository.Reservation().Create(ctx, &entity.Reservation{
		ID:                id,
		UserID:            input.UserID,
		TravelSpotID:      input.TravelSpotID,
		TravelSpotDiaryID: input.TravelSpotDiaryID,
		FromDate:          input.From,
		ToDate:            input.To,
	}); err != nil {
		return "", err
	}
	return id, nil
}

type ListOutputInput struct {
	UserID entity.UserID
	Status string
}

type ListOutput struct {
	Reservations   entity.Reservations
	TravelSpotByID map[entity.TravelSpotID]*entity.TravelSpot
	DiaryByID      map[entity.TravelSpotDiaryID]*entity.TravelSpotDiary
}

func (u *ReservationUseCase) List(ctx context.Context, input ListOutputInput) (*ListOutput, error) {
	var (
		reservations entity.Reservations
		err          error
	)
	switch input.Status {
	case "yet":
		reservations, err = u.repository.Reservation().GetUpcomingReservations(ctx, input.UserID)
	case "done":
		reservations, err = u.repository.Reservation().GetEndedReservations(ctx, input.UserID)
	default:
		return nil, errors.New(fmt.Sprintf("invalid status: %s", input.Status))
	}
	if err != nil {
		return nil, err
	}

	travelSpots, err := u.repository.TravelSpot().GetByIDs(ctx, reservations.TravelSpotIDs())
	if err != nil {
		return nil, err
	}

	travelSpotByID := make(map[entity.TravelSpotID]*entity.TravelSpot, len(travelSpots))
	for _, travelSpot := range travelSpots {
		travelSpotByID[travelSpot.ID] = travelSpot
	}

	diaries, err := u.repository.TravelSpotDiary().GetByIDs(ctx, reservations.TravelSpotDiaryIDs())
	if err != nil {
		return nil, err
	}
	diaryByID := make(map[entity.TravelSpotDiaryID]*entity.TravelSpotDiary, len(diaries))
	for _, diary := range diaries {
		diaryByID[diary.ID] = diary
	}

	return &ListOutput{
		Reservations:   reservations,
		TravelSpotByID: travelSpotByID,
		DiaryByID:      diaryByID,
	}, nil
}

type GetReservationOutput struct {
	Reservation                                     *entity.Reservation
	TravelSpot                                      *entity.TravelSpot
	TravelSpotItineraries                           entity.TravelSpotItineraries
	TravelSpotItineraryItemsByTravelSpotItineraryID map[entity.TravelSpotItineraryID]entity.TravelSpotItineraryItems
	TravelSpotDiary                                 *entity.TravelSpotDiary
}

func (u *ReservationUseCase) GetReservation(ctx context.Context, id entity.ReservationID) (*GetReservationOutput, error) {
	reservation, err := u.repository.Reservation().FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var (
		travelSpot                                      = &entity.TravelSpot{}
		travelSpotItineraries                           = entity.TravelSpotItineraries{}
		travelSpotItineraryItems                        = entity.TravelSpotItineraryItems{}
		travelSpotItineraryItemsByTravelSpotItineraryID = make(map[entity.TravelSpotItineraryID]entity.TravelSpotItineraryItems, len(travelSpotItineraries))
	)
	if !reservation.IsOffer {
		travelSpot, err = u.repository.TravelSpot().FindByID(ctx, reservation.TravelSpotID)
		if err != nil {
			return nil, err
		}

		travelSpotItineraries, err = u.repository.TravelSpotItinerary().GetByTravelSpotID(ctx, reservation.TravelSpotID)
		if err != nil {
			return nil, err
		}

		travelSpotItineraryItems, err = u.repository.TravelSpotItineraryItem().GetByTravelSpotItineraryIDs(ctx, travelSpotItineraries.IDs())
		if err != nil {
			return nil, err
		}

		for _, item := range travelSpotItineraryItems {
			travelSpotItineraryItemsByTravelSpotItineraryID[item.TravelSpotItineraryID] = append(travelSpotItineraryItemsByTravelSpotItineraryID[item.TravelSpotItineraryID], item)
		}
	}

	travelSpotDiary, err := u.repository.TravelSpotDiary().FindByID(ctx, reservation.TravelSpotDiaryID)
	if err != nil {
		return nil, err
	}

	return &GetReservationOutput{
		Reservation:           reservation,
		TravelSpot:            travelSpot,
		TravelSpotItineraries: travelSpotItineraries,
		TravelSpotItineraryItemsByTravelSpotItineraryID: travelSpotItineraryItemsByTravelSpotItineraryID,
		TravelSpotDiary: travelSpotDiary,
	}, nil
}

//go:embed img/fishing.jpeg
var fishingImg []byte

//go:embed img/camp.jpg
var campingImg []byte

type UpdateDiaryPhotoInput struct {
	ID     entity.ReservationID
	UserID entity.UserID
	Photo  Photo
}

func (u *ReservationUseCase) UpdateDiaryPhoto(ctx context.Context, input UpdateDiaryPhotoInput) (string, error) {
	reservation, err := u.repository.Reservation().FindByID(ctx, input.ID)
	if err != nil {
		return "", err
	}

	travelSpotDiary, err := u.repository.TravelSpotDiary().FindByID(ctx, reservation.TravelSpotDiaryID)
	if err != nil {
		return "", err
	}

	path, err := austerstorage.Save(
		austerstorage.ContentType(input.Photo.ContentType),
		filepath.Join("travel_spot_diaries", string(travelSpotDiary.UserID), string(reservation.TravelSpotDiaryID), input.Photo.Filename),
		input.Photo.Body,
	)
	if err != nil {
		return "", err
	}

	travelSpotDiary.PhotoPath = path
	if err := u.repository.TravelSpotDiary().Update(ctx, travelSpotDiary); err != nil {
		return "", err
	}

	if err := u.createSpecialOffer(ctx, input.UserID); err != nil {
		return "", err
	}

	return path, nil
}

type UpdateDiaryDescriptionInput struct {
	ID          entity.ReservationID
	UserID      entity.UserID
	Description string
}

func (u *ReservationUseCase) UpdateDiaryDescription(ctx context.Context, input UpdateDiaryDescriptionInput) error {
	reservation, err := u.repository.Reservation().FindByID(ctx, input.ID)
	if err != nil {
		return err
	}

	travelSpotDiary, err := u.repository.TravelSpotDiary().FindByID(ctx, reservation.TravelSpotDiaryID)
	if err != nil {
		return err
	}

	travelSpotDiary.Description = input.Description
	if err := u.repository.TravelSpotDiary().Update(ctx, travelSpotDiary); err != nil {
		return err
	}

	if err := u.createSpecialOffer(ctx, input.UserID); err != nil {
		return err
	}
	return nil
}

func (u *ReservationUseCase) createSpecialOffer(ctx context.Context, userID entity.UserID) error {
	// スペシャルオファーが存在すれば何もしない
	offer, err := u.repository.Reservation().FindSpecialOfferByUserID(ctx, userID)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if offer != nil {
		return nil
	}

	userHobbies, err := u.repository.UserHobby().GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	hobbies, err := u.repository.Hobby().GetByIDs(ctx, userHobbies.HobbyIDs())
	if err != nil {
		return err
	}

	var (
		specialOffers  = entity.Reservations{}
		specialDiaries = entity.TravelSpotDiaries{}
	)
	if slices.Contains(hobbies.Names(), "釣り") {
		path, err := austerstorage.Save(austerstorage.JPEG, "/travel_spot_diaries/special/fishing.jpg", fishingImg)
		if err != nil {
			return err
		}
		diary := &entity.TravelSpotDiary{
			ID:           austerid.Generate[entity.TravelSpotDiaryID](),
			UserID:       userID,
			TravelSpotID: "",
			Title:        "漁師の想いを発信",
			Date:         time.Date(2024, 12, 22, 0, 0, 0, 0, time.Local),
			PhotoPath:    path,
			Description:  "銚子の漁業組合でマーケターとして働き始めた。漁師さんたちの真摯な仕事ぶりや、美味しい魚が獲れる銚子の魅力を、SNSやWebで発信する役割だ。今日は朝市場で撮影し、鮮度抜群の魚の様子や、活気ある競りの雰囲気を配信。漁師さんたちの誇りと情熱を、より多くの人に伝えていきたい。",
		}

		specialOffers = append(specialOffers, &entity.Reservation{
			ID:                austerid.Generate[entity.ReservationID](),
			UserID:            userID,
			TravelSpotID:      "",
			TravelSpotDiaryID: diary.ID,
			FromDate:          time.Date(2024, 12, 22, 0, 0, 0, 0, time.Local),
			ToDate:            time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local),
			IsOffer:           true,
		})
		specialDiaries = append(specialDiaries, diary)
	}

	if slices.Contains(hobbies.Names(), "キャンプ") {
		path, err := austerstorage.Save(austerstorage.JPEG, "/travel_spot_diaries/special/camp.jpg", campingImg)
		if err != nil {
			return err
		}

		diary := &entity.TravelSpotDiary{
			ID:           austerid.Generate[entity.TravelSpotDiaryID](),
			UserID:       userID,
			TravelSpotID: "",
			Title:        "自然と音楽の交差点",
			Date:         time.Date(2024, 12, 22, 0, 0, 0, 0, time.Local),
			PhotoPath:    path,
			Description:  "銚子のアウトドアショップで、イベント企画の仕事を始めた。キャンプ場でのカクテルバーや、アコースティックライブなど、自然の中での特別な体験を提供。今日は秋のイベントの打ち合わせで、スタッフと意見交換。夕暮れのキャンプサイトで音楽を楽しむ企画が好評で、次回は規模を拡大することに。趣味を仕事にできる喜びを実感。",
		}
		specialOffers = append(specialOffers, &entity.Reservation{
			ID:                austerid.Generate[entity.ReservationID](),
			UserID:            userID,
			TravelSpotID:      "",
			TravelSpotDiaryID: diary.ID,
			FromDate:          time.Date(2024, 12, 22, 0, 0, 0, 0, time.Local),
			ToDate:            time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local),
			IsOffer:           true,
		})
		specialDiaries = append(specialDiaries, diary)
	}

	if err := u.repository.Reservation().Create(ctx, specialOffers...); err != nil {
		return err
	}
	if err := u.repository.TravelSpotDiary().Create(ctx, specialDiaries...); err != nil {
		return err
	}
	return nil
}
