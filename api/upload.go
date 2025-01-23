//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	upload1 "github.com/NpoolPlatform/file-gateway/pkg/gw/upload"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/file/gw/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Upload(ctx context.Context, in *npool.UploadRequest) (*npool.UploadResponse, error) {
	handler, err := upload1.NewHandler(
		ctx,
		upload1.WithPayload(in.GetPayload(), false),
	)
	if err != nil {
		logger.Sugar().Errorw("Upload", "Error", err)
		return &npool.UploadResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := handler.Upload(ctx)
	if err != nil {
		logger.Sugar().Errorw("Upload", "Error", err)
		return &npool.UploadResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UploadResponse{
		FileId: info,
	}, nil
}
