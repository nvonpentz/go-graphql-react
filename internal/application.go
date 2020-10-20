package internal

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Application struct {
	Router chi.Router
}

func NewApplicationWithDefaults() (*Application, error) {
	// Setup router
	router, err := NewRouterWithDefaults()
	if err != nil {
		return &Application{}, err
	}

	return &Application{
		Router: router,
	}, nil
}

func (application *Application) Start() {
	log.Println("Listening on :5000")
	err := http.ListenAndServe(":5000", application.Router)
	if err != nil {
		log.Fatal(err)
	}
}
