package http

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load location: %v\n", err)
		os.Exit(1)
	}
	time.Local = loc
}

type PerplexityClient struct {
	client *http.Client
	apiKey string
}

func NewPerplexityClient() *PerplexityClient {
	return &PerplexityClient{
		client: &http.Client{},
		apiKey: cmp.Or(os.Getenv("DIARY_PERPLEXITY_APIKEY"), ""),
	}
}

type CreateDiaryTextInput struct {
	TravelSpotTitle       string
	TravelSpotDescription string
	Hobby                 string
	UserAge               int
	UserGender            string
}

type CreateDiaryTextOutput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (c *PerplexityClient) CreateDiaryText(input *CreateDiaryTextInput) (*CreateDiaryTextOutput, error) {
	messages := []Message{
		{
			Role:    "system",
			Content: fmt.Sprintf("あなたは日記作成を支援するAIアシスタントです。%d歳%s、趣味は%sのペルソナで書いてください。応答は必ず次のJSON形式で返してください：{\"title\":\"5~12文字以内のタイトル\",\"description\":\"150~170文字以内の本文\"}", input.UserAge, input.UserGender, input.Hobby),
		},
		{
			Role:    "user",
			Content: fmt.Sprintf("体験タイトル:%s,体験内容:%s,体験地域:銚子", input.TravelSpotTitle, input.TravelSpotDescription),
		},
	}

	res, err := c.perplexityRequest(messages)
	if err != nil {
		return nil, fmt.Errorf("failed to request perplexity: %w", err)
	}

	var out CreateDiaryTextOutput
	if err := json.Unmarshal([]byte(res.Choices[0].Message.Content), &out); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &out, nil
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func (c *PerplexityClient) perplexityRequest(messages []Message) (*PerplexityResponse, error) {
	var (
		url   = "https://api.perplexity.ai/chat/completions"
		model = "llama-3.1-sonar-huge-128k-online"
	)
	jsonBody, err := json.Marshal(Request{
		Model:    model,
		Messages: messages,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res PerplexityResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// PerplexityResponse はAPIレスポンス全体を表す構造体
type PerplexityResponse struct {
	ID        string       `json:"id"`
	Model     string       `json:"model"`
	Created   int64        `json:"created"`
	Usage     UsageInfo    `json:"usage"`
	Citations []string     `json:"citations"`
	Object    string       `json:"object"`
	Choices   []ChoiceInfo `json:"choices"`
}

// CreatedTime はCreatedフィールドをtime.Time型として返すメソッド
func (pr *PerplexityResponse) CreatedTime() time.Time {
	return time.Unix(pr.Created, 0)
}

// UsageInfo はトークン使用量の情報を表す構造体
type UsageInfo struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChoiceInfo は生成された選択肢の情報を表す構造体
type ChoiceInfo struct {
	Index        int         `json:"index"`
	FinishReason string      `json:"finish_reason"`
	Message      MessageInfo `json:"message"`
	Delta        DeltaInfo   `json:"delta"`
}

// MessageInfo は生成されたメッセージの情報を表す構造体
type MessageInfo struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// DeltaInfo はストリーミング応答時のデルタ情報を表す構造体
type DeltaInfo struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
