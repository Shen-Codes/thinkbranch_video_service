package ResponseHandler

import {
	"net/http"
}

type Response struct {
	Message string `dynamodbav:"message" json: "message"`
	MessageId string `dynamodbav: "message_id" json: "messageId"`
	ResponseToId string `dynamodbav: "response_to_id" json: "responseToId"`
	ResponseFromId string `dynamodbav: "response_from_id" json: "responseFromId`
	
}

func ResponseGetter(dynaSvc dynamodb.Client, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func ResponsePoster(dynaSvc dynamodb.Client, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func ResponseUpdater(dynaSvc dynamodb.Client, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func ResponseDeleter(dynaSvc dynamodb.Client, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

