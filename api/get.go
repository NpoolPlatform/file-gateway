package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/NpoolPlatform/file-gateway/common/servermux"
	get1 "github.com/NpoolPlatform/file-gateway/pkg/gw/get"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

func init() {
	mux := servermux.AppServerMux()
	mux.HandleFunc("/v1/images/", Image)
}

func Image(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	origin := r.Header.Get("Origin")
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Vary", "Origin")
	}

	var nonEmptyParts []string
	for _, part := range parts {
		if part != "" {
			nonEmptyParts = append(nonEmptyParts, part)
		}
	}

	minPathLength := 3
	if len(parts) < minPathLength {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	handler, err := get1.NewHandler(
		ctx,
		get1.WithFileID(nonEmptyParts[2], true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Image",
			"Error", err,
		)
		fmt.Fprintf(w, "%v", err.Error())
		return
	}

	info, err := handler.Get(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"Image",
			"Error", err,
		)
		fmt.Fprintf(w, "%v", err.Error())
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(string(info))
	if err != nil {
		decoded = info
	}

	if strings.HasPrefix(string(decoded), "<svg xmlns=") {
		w.Header().Set("Content-Type", "image/svg+xml")
	} else {
		w.Header().Set("Content-Type", "image/jpeg")
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(decoded)))
	w.Header().Set("Content-Disposition", "inline")

	_, err = w.Write(decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
