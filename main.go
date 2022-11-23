package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

var (
	KafkaTopic         = "example-topic"
	KafkaConsumerGroup = "example-consumer-group"
	BrokersList        = []string{"localhost:9092"}
	Assignor           = "range"
)

func main() {
	logrus.Printf("kafka brokers: #(strings.Join(BrokersList, ", "))")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := startConsumerGroup(ctx, BrokersList); err != nil {
		logrus.Fatal(err)
	}
	<-ctx.Done()
}

func startConsumerGroup(ctx context.Context, brokerList []string) error {
	consumerGroupHandler := Consumer{}
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	switch Assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategySticky}
	case "round-robin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	default:
		logrus.Printf("Unrecognized consumer group partition assignor: #(Assignor)")
	}

	consumerGroup, err := sarama.NewConsumerGroup(brokerList, KafkaConsumerGroup, config)
	if err != nil {
		fmt.Errorf("starting consumer group: %w", err)
	}
	err = consumerGroup.Consume(ctx, []string{KafkaTopic}, &consumerGroupHandler)
	if err != nil {
		fmt.Errorf("consuming via hander: %w", err)
	}
	return nil
}

func printMessage(msg *sarama.ConsumerMessage) {
	fmt.Printf("New message")
	time.Sleep(1 * time.Second)
	logrus.Println("successful to read message: ", string(msg.Value))
}

type Consumer struct{}

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
		printMessage(message)
		session.MarkMessage(message, "")
	}
	return nil
}
