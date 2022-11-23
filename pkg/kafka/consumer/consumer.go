package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/user_grade_api/internal/entity"
	"github.com/vladjong/user_grade_api/internal/storage"
)

type ConsumerOptions struct {
	KafkaTopic         string
	KafkaConsumerGroup string
	BrokersList        []string
	Assignor           string
}

type Consumer struct {
	Storage storage.UserStorager
}

func StartConsumerGroup(ctx context.Context, option ConsumerOptions, storage storage.UserStorager) error {
	consumerGroupHandler := Consumer{
		Storage: storage,
	}
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	switch option.Assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategySticky}
	case "round-robin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	default:
		logrus.Printf("Unrecognized consumer group partition assignor: #(Assignor)")
	}

	consumerGroup, err := sarama.NewConsumerGroup(option.BrokersList, option.KafkaConsumerGroup, config)
	if err != nil {
		return err
	}
	err = consumerGroup.Consume(ctx, []string{option.KafkaTopic}, &consumerGroupHandler)
	if err != nil {
		return err
	}
	return nil
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	fmt.Println("consumer - setup")
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Println("consumer - cleanup")
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var request entity.UserGrade
		if err := json.Unmarshal(message.Value, &request); err != nil {
			return err
		}
		c.Storage.SetUser(request)
		session.MarkMessage(message, "")
	}
	return nil
}
