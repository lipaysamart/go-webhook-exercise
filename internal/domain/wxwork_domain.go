package domain

import "context"

type WxworkPayload struct {
	MessageType string                 `json:"msgtype"`
	Mark        map[string]interface{} `json:"markdown"`
}

type IWebhook interface {
	Receive(ctx context.Context, payload map[string]interface{}) error
	Send(ctx context.Context, payload *WxworkPayload) error
}
