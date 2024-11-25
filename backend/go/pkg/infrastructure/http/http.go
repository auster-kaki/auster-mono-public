package http

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	appRPC "github.com/auster-kaki/auster-mono/pkg/app/rpc"
)

type Client struct {
	client *http.Client
	host   string
	port   string
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
		host:   cmp.Or(os.Getenv("DIARY_HOST"), "localhost"),
		port:   cmp.Or(os.Getenv("DIARY_PORT"), ""),
	}
}

// URL -
func (c *Client) URL() string {
	return fmt.Sprintf("http://%s:%s", c.host, c.port)
}

func (c *Client) Diary() appRPC.Diary { return c }

func (c *Client) CreateImage(ctx context.Context, input appRPC.CreateImageInput) (appRPC.CreateImageOutput, error) {
	var (
		swapFacePath   = input.TargetPath
		swapBeforePath = input.SourcePath
	)
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	file, err := os.Open(swapFacePath)
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}
	defer file.Close()

	// ファイルを追加
	part, err := writer.CreateFormFile("source_path", filepath.Base(swapFacePath))
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}

	targetFile, err := os.Open(swapBeforePath)
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}
	defer targetFile.Close()

	targetPart, err := writer.CreateFormFile("target_path", filepath.Base(swapBeforePath))
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}
	_, err = io.Copy(targetPart, targetFile)
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL()+"/process", &buf)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}
	defer res.Body.Close()

	var response struct {
		JobID  string `json:"job_id"`
		Status string `json:"status"`
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return appRPC.CreateImageOutput{}, err
	}

	return appRPC.CreateImageOutput{
		JobID: response.JobID,
	}, nil
}

func (c *Client) GetStatus(ctx context.Context, input appRPC.GetStatusInput) (appRPC.GetStatusOutput, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.URL()+"/job_status/"+input.JobID, nil)
	if err != nil {
		return appRPC.GetStatusOutput{}, fmt.Errorf("failed to create request: %w", err)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return appRPC.GetStatusOutput{}, fmt.Errorf("failed to do request: %w", err)
	}
	defer res.Body.Close()

	var response struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return appRPC.GetStatusOutput{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return appRPC.GetStatusOutput{
		// NOTE:クソみたいな仕様です....
		JobID:  input.JobID,
		Status: response.Status,
	}, nil
}

// GetImagePath ファイル書き込みの設定をしてないので、未検証
func (c *Client) GetImagePath(ctx context.Context, input appRPC.GetImagePathInput) (appRPC.GetImagePathOutput, error) {
	var (
		filename = fmt.Sprintf("%s.png", input.JobID)
		url      = fmt.Sprintf(c.URL() + "/generated/" + filename)
	)
	res, err := http.Get(url)
	if err != nil {
		return appRPC.GetImagePathOutput{}, fmt.Errorf("failed to get image: %w", err)
	}
	defer res.Body.Close()

	// ステータスコードを確認
	if res.StatusCode != http.StatusOK {
		return appRPC.GetImagePathOutput{}, fmt.Errorf("server error: status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return appRPC.GetImagePathOutput{}, fmt.Errorf("failed to read body: %w", err)
	}

	return appRPC.GetImagePathOutput{
		GeneratedImage: body,
		Filename:       filename,
	}, nil
}
