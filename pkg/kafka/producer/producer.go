package producer

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/user_grade_api/internal/entity"
)

type ProducerOptions struct {
	KafkaTopic  string
	BrokersList []string
}

type Producer struct {
	options  ProducerOptions
	producer sarama.AsyncProducer
}

func New(options *ProducerOptions) (producer Producer, err error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Producer.Return.Successes = true
	saramaProducer, err := sarama.NewAsyncProducer(options.BrokersList, config)
	if err != nil {
		return producer, err
	}
	go func() {
		for err := range saramaProducer.Errors() {
			logrus.Println("Failed to write message:", err)
		}
	}()
	return Producer{
		options:  *options,
		producer: saramaProducer,
	}, nil
}

func (p *Producer) SendMessage(user entity.UserGrade) error {
	msgBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	msg := sarama.ProducerMessage{
		Topic: p.options.KafkaTopic,
		Key:   sarama.StringEncoder(user.UserId),
		Value: sarama.ByteEncoder(msgBytes),
	}
	p.producer.Input() <- &msg
	successMsg := <-p.producer.Successes()
	logrus.Println("Succesful to write message, offset:", successMsg.Offset)
	if err != nil {
		return err
	}
	return nil
}
