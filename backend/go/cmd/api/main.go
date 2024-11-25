package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/auster-kaki/auster-mono/pkg/app/handler"
	"github.com/auster-kaki/auster-mono/pkg/app/handler/middleware"
	infraHTTP "github.com/auster-kaki/auster-mono/pkg/infrastructure/http"
	"github.com/auster-kaki/auster-mono/pkg/infrastructure/rdb"
	"github.com/auster-kaki/auster-mono/pkg/logging"

	"github.com/rs/cors"
)

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load location: %v\n", err)
		os.Exit(1)
	}
	time.Local = loc
}

func main() {
	if err := _main(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func _main() error {
	d, err := rdb.NewDB()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	for path, h := range handler.NewHandlerMap(d, infraHTTP.NewClient()) {
		mux.HandleFunc(path, h)
	}

	logging.Info(context.Background(), "server is running on :8080")
	if err := http.ListenAndServe(":8080",
		middleware.NewRecover().Handle(cors.AllowAll().Handler(mux)),
	); err != nil {
		return err
	}
	return nil
}
