package main

import (
	"flag"
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {

	topic := flag.String("key", "", "Key")
	numOfPartitions := flag.Int("partitions", 12, "Partitions")

	flag.Parse()

	if topic == nil || *topic == "" {
		fmt.Println("Invalid Key")
		return
	}

	message := sarama.ProducerMessage{
		Topic: "trip",
		Key:   sarama.StringEncoder(*topic),
		Value: nil,
	}

	fmt.Println("Using default partitioner: Hash")

	partitioner := sarama.NewHashPartitioner("trip")
	partition, _ := partitioner.Partition(&message, int32(*numOfPartitions))

	fmt.Println(fmt.Sprintf("Partition: %d", partition))

	return
}
