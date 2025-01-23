package get

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/oss"
)

func (h *Handler) Get(ctx context.Context) ([]byte, error) {
	return oss.GetObject(ctx, h.FileID, true)
}
