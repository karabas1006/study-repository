package http_server

import (
	"errors"
	"fmt"
	"net/http"
)

func StartHTTPServer() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Obrabotka zaprosa na patterne /ping")
		w.Write([]byte("Hello from Docker!\n"))
	})

	err := http.ListenAndServe(":5050", nil)

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}
