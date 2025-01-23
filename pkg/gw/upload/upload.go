package upload

import (
	"context"
	"encoding/hex"

	"golang.org/x/crypto/sha3"

	"github.com/NpoolPlatform/go-service-framework/pkg/oss"
)

func (h *Handler) Upload(ctx context.Context) (string, error) {
	digest := sha3.Sum256([]byte(h.Payload))
	hash := hex.EncodeToString(digest[:])

	if err := oss.PutObject(ctx, hash, []byte(h.Payload), true); err != nil {
		return "", err
	}

	return hash, nil
}
