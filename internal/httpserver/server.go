package httpserver

import (
	"fmt"
	"net/http"

	"github.com/rifkisururi/belajar-golang/internal/greeting"
	"github.com/rifkisururi/belajar-golang/internal/task"
)

func New() http.Handler {
	mux := http.NewServeMux()
	tasks := task.NewHandler(task.NewService(task.NewMemoryRepository())).Routes()
	mux.Handle("/tasks", tasks)
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("GET /greet", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = fmt.Fprintln(w, greeting.Message(r.URL.Query().Get("name")))
	})
	return mux
}
