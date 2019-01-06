package kafka_test

import (
	"fmt"
	"github.com/bxsmart/bxcore/kafka"
	"github.com/bxsmart/bxcore/log"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	log.InitializeTest()
	address := "127.0.0.1:9092"
	register := &kafka.ConsumerRegister{}
	brokerList := make([]string, 0)
	brokerList = append(brokerList, address)
	register.Initialize(brokerList)
	err := register.RegisterTopicAndHandler("test", "group1", TestData{}, func(data interface{}) error {
		dataValue := data.(*TestData)
		//fmt.Printf("Msg : %s, Timestamp : %s \n", dataValue.Msg, dataValue.Timestamp)
		fmt.Printf("Msg : %s\n", dataValue.Msg,)
		return nil
	})
	if err != nil {
		fmt.Errorf("Failed register")
		println(err)
	}
	time.Sleep(1000 * time.Second)

	defer func() {
		register.Close()
	}()
}
