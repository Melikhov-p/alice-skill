package main

import (
	"fmt"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(fmt.Sprintf(`Unhandled error %v`, err.Error()))
	}
}

func run() error {
	return http.ListenAndServe(`localhost:8080`, http.HandlerFunc(WebHook))
}

func WebHook(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	writer.Header().Set(`Content-Type`, `application/json`)

	_, _ = writer.Write([]byte(`{
        "response": {
          "text": "Извините, я пока ничего не умею"
        },
        "version": "1.0"
      }`))
}
