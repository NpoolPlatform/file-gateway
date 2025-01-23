package get

import (
	"context"
	"fmt"
)

type Handler struct {
	FileID string
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

func WithFileID(fileID string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if fileID == "" {
			return fmt.Errorf("invalid fileId")
		}
		h.FileID = fileID
		return nil
	}
}
