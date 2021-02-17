package main

import (
	"context"
	"log"
	"net/http"
	"os"

	T "github.com/Shen-Codes/thinkbranch_video_service/TopicsHandler"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}

	dynaSvc := dynamodb.NewFromConfig(cfg)

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler(dynaSvc))
	router.HandleFunc("/topics", T.TopicsGetter(dynaSvc)).Methods("GET")
	router.HandleFunc("/topics/{topic}", T.TopicGetter(dynaSvc)).Methods("GET")
	router.HandleFunc("/topics/{topic}", ResponsePoster(dynaSvc)).Methods("POST")
	router.HandleFunc("/topics/{topic}/{message_id}", ResponseUpdater(dynaSvc)).Methods("PUT")
	router.HandleFunc("/topics/{topic}/{message_id}", ResponseDeleter(dynaSvc)).Methods("DELETE")

	http.Handle("/", router)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
