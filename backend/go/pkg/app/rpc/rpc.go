package rpc

import "context"

type RPC interface {
	Diary() Diary
}

type CreateImageInput struct {
	// SourcePath は、顔ハメ前の画像
	SourcePath string
	// TargetPath は、ユーザーの画像の保存先パス
	TargetPath string
}

type CreateImageOutput struct {
	JobID string
}

type GetStatusInput struct {
	JobID string
}

type GetStatusOutput struct {
	JobID  string
	Status string
}

type GetImagePathInput struct {
	JobID string
}

type GetImagePathOutput struct {
	GeneratedImage []byte
	Filename       string
}

type Diary interface {
	CreateImage(ctx context.Context, input CreateImageInput) (createImageOutput CreateImageOutput, err error)
	GetStatus(ctx context.Context, input GetStatusInput) (getStatusOutput GetStatusOutput, err error)
	// GetImagePath は、画像のGoServerの上の生成された画像パスを取得する。
	GetImagePath(ctx context.Context, input GetImagePathInput) (getImagePathOutput GetImagePathOutput, err error)
}
