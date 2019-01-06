package kafka_test

import (
	"fmt"
	"github.com/bxsmart/bxcore/kafka"
	"strings"
	"testing"
	"time"
)

type TestData struct {
	Msg       string
}

func TestProducer(t *testing.T) {
	brokers := strings.Split("192.168.10.119:9092", ",")
	producerWrapped := &kafka.MessageProducer{}
	err := producerWrapped.Initialize(brokers)

	if err != nil {
		fmt.Printf("Failed init producerWrapped %s", err.Error())
	}

	for i := 0; i < 10; i++ {
		//t := time.Now()
		data := &TestData{"msg"}
		partition, offset, sendRes := producerWrapped.SendMessage("deals", data, "1")
		if sendRes != nil {
			fmt.Errorf("failed to sendmsg : %s", err.Error())
		} else {
			fmt.Printf("Your data is stored with unique identifier important/%d/%d\n", partition, offset)
		}
		time.Sleep(time.Second * 3)
	}

	defer func() {
		producerWrapped.Close()
	}()

}
