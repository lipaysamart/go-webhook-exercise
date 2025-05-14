package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"

	"github.com/lipaysamart/go-webhook-exercise/internal/domain"
)

var (
	WxworkWebHookURL = os.Getenv("wxwork_webhook_url")
)

type wxworkUsecase struct {
	Logger *zerolog.Logger
	hc     *http.Client
}

func NewWxworkUsecase(logger *zerolog.Logger) domain.IWebhook {
	return &wxworkUsecase{
		Logger: logger,
		hc:     &http.Client{},
	}
}

func (u *wxworkUsecase) Receive(ctx context.Context, payload map[string]interface{}) error {
	_, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	u.Logger.Info().Any("payload", payload).Msg("print payload")
	return nil
}

func (u *wxworkUsecase) Send(ctx context.Context, payload *domain.WxworkPayload) error {
	_, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&payload); err != nil {
		u.Logger.Trace().Ctx(ctx).Any("payload", payload).Msg("failed to encode payload")
		return err
	}

	if err := u.makeRequest(ctx, WxworkWebHookURL, "POST", &buf); err != nil {
		return err
	}
	return nil
}

func (u *wxworkUsecase) makeRequest(ctx context.Context, url string, method string, body io.Reader) error {
	httpReq, _ := http.NewRequestWithContext(ctx, method, url, body)
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := u.hc.Do(httpReq)
	if httpResp.StatusCode != 200 || err != nil {
		u.Logger.Trace().Ctx(ctx).Err(err).Int("statusCode", httpResp.StatusCode).Msg("failed to request")
		return err
	}

	defer httpResp.Body.Close()
	return nil
}
