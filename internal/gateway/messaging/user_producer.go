package messaging

import (
	"go-clean-architecture-pzn/internal/model"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type UserProducer struct {
	Producer[*model.UserEvent]
}

func NewUserProducer(producer *kafka.Writer, log *logrus.Logger) *UserProducer {
	return &UserProducer{
		Producer: Producer[*model.UserEvent]{
			Producer: producer,
			Topic:    "users",
			Log:      log,
		},
	}
}
