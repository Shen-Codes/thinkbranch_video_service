package TopicsHandler

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Topic struct {
	TopicTitle string `dynamodbav:"topic_title" json: "topicTitle"`
	TopicBody  string `dynamodbav:"topic_body" json: "topicBody"`
	ResponseId string `dynamodbav:"response_id" json: "responseId"`
}

func TopicsGetter(dynaSvc dynamodb.Client, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func TopicGetter(dynaSvc dynamodb.Client, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
