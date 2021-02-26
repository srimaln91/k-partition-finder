package main

import (
	"flag"
	"fmt"

	"github.com/Shopify/sarama"
	buildMeta "github.com/srimaln91/go-make"
)

func main() {

	// Print binary details and terminate the program when --version flag provided.
	buildMeta.CheckVersion()

	key := flag.String("key", "", "Key")
	numOfPartitions := flag.Int("partitions", 12, "Partitions")

	flag.Parse()

	if key == nil || *key == "" {
		fmt.Println("Invalid Key")
		return
	}

	message := sarama.ProducerMessage{
		Key:   sarama.StringEncoder(*key),
		Value: nil,
	}

	fmt.Println("Using default partitioner: Hash")

	partitioner := sarama.NewHashPartitioner("trip")
	partition, _ := partitioner.Partition(&message, int32(*numOfPartitions))

	fmt.Println(fmt.Sprintf("Partition: %d", partition))

	return
}
