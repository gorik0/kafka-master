package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Aqualung struct {
	Brand          string
	MaxDepth       int
	IsForColdWater bool
	Cost           int
}

var BRAND_LIST = strings.Split("AQua Nemo aqualung wateri oceanico romeshko", " ")
var BROKERS = []string{"localhost:9096", "localhost:9096"}

func makeAqua() *Aqualung {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &Aqualung{
		Brand:          BRAND_LIST[r.Intn(len(BRAND_LIST))],
		MaxDepth:       r.Intn(100),
		IsForColdWater: false,
		Cost:           r.Intn(100) * 100,
	}
}

func createProducerAsync() (sarama.AsyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner

	return sarama.NewAsyncProducer(BROKERS, cfg)
}
func createProducer() (sarama.SyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner

	return sarama.NewSyncProducer(BROKERS, cfg)
}

func prepareMsg(msgBody []byte) sarama.ProducerMessage {
	return sarama.ProducerMessage{
		Topic:     "songs",
		Value:     sarama.ByteEncoder(msgBody),
		Partition: -1,
	}
}

func main() {

	//	::; SETUP clients
	sync, err := createProducer()
	_ = sync
	HandlerErr(err, "producer creation failed")

	async, err := createProducerAsync()
	HandlerErr(err, "async producer creation failed")
	//	LISTEN to success\error chgannels

	go func() {

		for message := range async.Successes() {
			log.Println("--->>> Succes for %s", message.Value)

		}

	}()
	go func() {

		for e := range async.Errors() {
			log.Println("--->>> Error for %s", e.Error())

		}

	}()

	//	::: RUN sync/async producer
	for {

		msgBytes, _ := json.Marshal(makeAqua())
		msg := prepareMsg(msgBytes)
		async.Input() <- &msg

		time.Sleep(500 * time.Millisecond)
	}

}

func HandlerErr(err error, mes string) {
	if err != nil {
		log.Fatalf("Error while %s, :::: %w", mes, err)
	}
}
