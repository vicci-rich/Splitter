package kafka

import (
	"github.com/Shopify/sarama"
	"strings"
	"time"
)

type Producer struct {
	kafkaConfig    *sarama.Config
	producers      []sarama.AsyncProducer
	messageChannel chan *Message
	errorChannel   chan *MessageError
	stopChannel    chan bool
}

type ProducerConfig struct {
	BrokerList       string
	BufferSize       int
	ProducerNum      int
	FlushMessages    int
	FlushFrequency   int
	FlushMaxMessages int
	Timeout          int
	ReturnErrors     bool
}

type Message struct {
	Topic     string
	Partition int32
	Offset    int64
	Key       []byte
	Data      []byte
}

type MessageError struct {
	Error     error
	Topic     string
	Timestamp time.Time
	Partition int32
	Offset    int64
	Metadata  interface{}
}

func NewProducer(cfg *ProducerConfig) (*Producer, error) {
	producer := new(Producer)
	producer.producers = make([]sarama.AsyncProducer, 0)
	producer.messageChannel = make(chan *Message, cfg.BufferSize)
	producer.errorChannel = make(chan *MessageError, cfg.BufferSize)
	producerNum := 1
	if cfg.ProducerNum > 1 {
		producerNum = cfg.ProducerNum
	}

	producer.stopChannel = make(chan bool, producerNum)
	for i := 0; i < producerNum; i++ {
		brokers := cfg.BrokerList
		producer.kafkaConfig = sarama.NewConfig()
		producer.kafkaConfig.Producer.RequiredAcks = sarama.WaitForLocal
		producer.kafkaConfig.Producer.Return.Errors = cfg.ReturnErrors
		producer.kafkaConfig.Producer.Return.Successes = false
		producer.kafkaConfig.Producer.Flush.Messages = cfg.FlushMessages
		producer.kafkaConfig.Producer.Flush.Frequency = time.Millisecond * time.Duration(cfg.FlushFrequency)
		producer.kafkaConfig.Producer.Flush.MaxMessages = cfg.FlushMaxMessages
		producer.kafkaConfig.Producer.Timeout = time.Millisecond * time.Duration(cfg.Timeout)
		p, err := sarama.NewAsyncProducer(strings.Split(brokers, ","), producer.kafkaConfig)
		if err != nil {
			return nil, err
		}
		producer.producers = append(producer.producers, p)
	}

	return producer, nil
}

func (p *Producer) MessageChannel() chan *Message {
	return p.messageChannel
}

func (p *Producer) ErrorChannel() <-chan *MessageError {
	return p.errorChannel
}

func (p *Producer) Start() {
	for _, producer := range p.producers {
		go p.receiveMessages(producer)
		if p.kafkaConfig.Producer.Return.Errors {
			go p.receiveErrors(producer)
		}
	}
}

func (p *Producer) Stop() {
	for i := 0; i < len(p.producers); i++ {
		p.stopChannel <- true
	}
}

func (p *Producer) receiveMessages(producer sarama.AsyncProducer) {
	for {
		select {
		case msg := <-p.messageChannel:
			producer.Input() <- &sarama.ProducerMessage{
				Topic: msg.Topic,
				Value: sarama.ByteEncoder(msg.Data),
			}
		case stop := <-p.stopChannel:
			if stop {
				return
			}
		}
	}
}

func (p *Producer) receiveErrors(producer sarama.AsyncProducer) {
	for {
		select {
		case msg := <-producer.Errors():
			p.errorChannel <- &MessageError{
				msg.Err,
				msg.Msg.Topic,
				msg.Msg.Timestamp,
				msg.Msg.Partition,
				msg.Msg.Offset,
				msg.Msg.Metadata,
			}
		}
	}
}

func (p *Producer) Send(topic string, data []byte) {
	msg := &Message{
		Topic: topic,
		Data:  data,
	}
	p.messageChannel <- msg
}
