package upload

import (
	"context"
	"fmt"
)

type Handler struct {
	Payload string
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithPayload(payload string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if payload == "" {
			return fmt.Errorf("Invalid payload")
		}
		h.Payload = payload
		return nil
	}
}

