package api

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

var emptyTopicAndMsg = "topicID and msg are required"
var gPubSubEmptyProjectID = "ProjectID can not be blank"

// send a message to google pub sub
func (g googlePubSub) send(topicID string, msg interface{}) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, g.ProjectID)
	if err != nil {
		return errors.Wrap(err, "messageapi publish new client error:")
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "messageapi publish convert msg to data err:")
	}

	t := client.Topic(topicID)
	ctxto, cancel := context.WithTimeout(ctx, 14*time.Second)
	defer cancel()
	result := t.Publish(ctxto, &pubsub.Message{
		Data: data,
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	_, err = result.Get(ctx)
	if err != nil {
		return errors.Wrap(err, "messageapi could not get message id: ")
	}

	return nil
}

// publish a message to the current message broker
func (m messageAPI) publish(topicID string, msg interface{}) error {
	if topicID == "" || msg == nil {
		return errors.New(emptyTopicAndMsg)
	}

	return m.API.send(topicID, msg)
}

func createGooglePubSub(projectID string) (*googlePubSub, error) {
	if projectID == "" {
		return nil, errors.New(gPubSubEmptyProjectID)
	}

	return &googlePubSub{ProjectID: projectID}, nil
}

// PublishMessage a message to the message service
func PublishMessage(projectID, topicID string, msg interface{}) error {
	g, err := createGooglePubSub(projectID)
	if err != nil {
		return errors.Wrap(err, "PublishMessage:")
	}

	m := messageAPI{API: g}
	return m.publish(topicID, msg)
}
