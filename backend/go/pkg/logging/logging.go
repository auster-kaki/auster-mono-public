package logging

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	w := os.Stdout
	slog.SetDefault(slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.MessageKey:
				return slog.Attr{Key: "message", Value: a.Value}
			case slog.SourceKey:
				const skip = 8
				_, file, line, ok := runtime.Caller(skip)
				if !ok {
					return a
				}
				v := fmt.Sprintf("%s:%d", filepath.Base(file), line)
				a.Value = slog.StringValue(v)
			}
			return a
		},
	})))
}

func Info(ctx context.Context, msg string) {
	slog.InfoContext(ctx, msg)
}

func Warn(ctx context.Context, msg string) {
	slog.WarnContext(ctx, msg)
}

func Error(ctx context.Context, msg string) {
	slog.ErrorContext(ctx, msg)
}
