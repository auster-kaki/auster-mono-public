package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
)

// Recover は panic からrecoverしてサーバエラーを返す
type Recover struct{}

// NewRecover -
func NewRecover() *Recover { return &Recover{} }

// Handle -
func (*Recover) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				response.HandleError(r.Context(), w, fmt.Errorf("panic: %v\n%v", p, stackTrace()))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func stackTrace() string {
	st := ""
	for depth := 0; ; depth++ {
		_, file, line, ok := runtime.Caller(depth)
		if !ok {
			break
		}
		st += fmt.Sprintf("%d:%v:%d\n", depth, file, line)
	}
	return st
}
