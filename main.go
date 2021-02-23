package main

import (
	"flag"
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {

	keyPtr := flag.String("key", "", "Key")

	flag.Parse()

	if keyPtr == nil || *keyPtr == "" {
		fmt.Println("Invalid Key")
		return
	}

	message := sarama.ProducerMessage{
		Topic: "trip",
		Key:   sarama.StringEncoder(*keyPtr),
		Value: nil,
	}

	fmt.Println("Using default partitioner: Hash")

	partitioner := sarama.NewHashPartitioner("trip")
	partition, _ := partitioner.Partition(&message, 12)
	fmt.Println(fmt.Sprintf("Partition: %d", partition))

	return
}
