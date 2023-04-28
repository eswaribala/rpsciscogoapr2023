package main

import (
	"CiscoApr2023/day1/models"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	topic         = "bankingqueue"
	brokerAddress = "localhost:9092"
)

//var wg sync.WaitGroup

func main() {
	// create a new context
	//wg.Add(10)
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	produce(ctx)
	//wg.Wait()
	fmt.Println("Returning to main....")

}

func produce(ctx context.Context) {
	//defer wg.Done()
	// initialize a counter

	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})
	i := 0
	for {
		customer := models.Customer{
			AccountNo: int64(rand.Int31n(1000000)),
			FirstName: "Parameswari", LastName: "Bala" + strconv.Itoa(int(rand.Int31n(10000))),
			Address: models.Address{DoorNo: "13A",
				StreetName: "Rajaji Nagar", City: "Bangalore", State: "KN"},
			DOB: models.Date{
				Day: 1 + int(rand.Int31n(26)), Month: 1 + int(rand.Int31n(11)),
				Year: 1900 + int(rand.Int31n(122))},
			Email: "Param@gmail.com", Active: true}
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// create an arbitrary message payload for the value

			Value: []byte(customer.FirstName),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
	}
}
