package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"
)

type fakeMessageBroker struct {
	mock.Mock
}

func (mock *fakeMessageBroker) send(topicID string, msg interface{}) error {
	args := mock.Called(topicID, msg)
	log.Printf("args: %+v", args)
	return args.Error(0)
}

func TestMessageBroker_Send(t *testing.T) {
	mb := &fakeMessageBroker{}
	mb.On("send", "topic", "message").Return(nil)

	service := &messageAPI{
		API: mb,
	}

	err := service.publish("topic", "message")
	assert.NoError(t, err)
	mb.AssertExpectations(t)
}

func TestMessageBroker_Send_ErrorEmptyTopic(t *testing.T) {
	mb := &fakeMessageBroker{}
	mb.On("send", "", "message").Return(nil)

	service := &messageAPI{
		API: mb,
	}

	err := service.publish("", "message")
	assert.Error(t, err)
}

func TestMessageBroker_Send_ErrorEmptyMessage(t *testing.T) {
	mb := &fakeMessageBroker{}
	mb.On("send", "topic", nil).Return(nil)

	service := &messageAPI{
		API: mb,
	}

	err := service.publish("topic", nil)
	assert.Error(t, err)
}

func TestCreateGooglePubSub(t *testing.T) {
	gPubSub, err := createGooglePubSub("project-id")
	assert.NoError(t, err)

	assert.Contains(t, "project-id", gPubSub.ProjectID)
}

func TestCreateGooglePubSub_ErrorBlankProjectID(t *testing.T) {
	gPubSub, err := createGooglePubSub("")
	assert.Errorf(t, err, gPubSubEmptyProjectID)
	assert.Nil(t, gPubSub)
}
