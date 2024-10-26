package main

import (
	"fmt"
	"net/http"
)

func main() {
	parseFlags()
	if err := run(); err != nil {
		panic(fmt.Sprintf(`Unhandled error %v`, err.Error()))
	}
}

func run() error {
	fmt.Println("Running server on " + flagRunAddr)
	return http.ListenAndServe(flagRunAddr, http.HandlerFunc(WebHook))
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
