package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"lidiyatrf/user-actions-test-task/internal/config"
	"lidiyatrf/user-actions-test-task/internal/service"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "http server port")
	flag.Parse()

	s, err := service.New(config.Config{
		ActionsFilePath: "./sources/actions.json",
		UsersFilePath:   "./sources/users.json",
	})
	if err != nil {
		log.Fatalf("unable to create service: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", s.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}/actions/count", s.GetUserActionsCount).Methods("GET")
	r.HandleFunc("/users/{id}/actions/{actionType}/next", s.GetUserNextActions).Methods("GET")
	r.HandleFunc("/referralIndexes", s.GetReferralIndexes).Methods("GET")

	log.Println("server is running on port", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		log.Fatal(err)
	}
}
