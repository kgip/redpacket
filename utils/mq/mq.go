package mq

import (
	"errors"
	"log"
	"sync"
	"time"
)

var (
	QueueNotExist = errors.New("message queue does not exist")
)

type MqOperator interface {
	SendMessage(topic string, msg interface{}, expire time.Duration)
	RegistryMessageHandler(topic string, handler func(msg interface{})) (error error)
}

type LocalMQ struct {
	lock      *sync.Mutex
	msgGroups map[string]chan interface{}
}

func NewLocalMQ() *LocalMQ {
	return &LocalMQ{lock: &sync.Mutex{}, msgGroups: map[string]chan interface{}{}}
}

func (mq *LocalMQ) trySendMessage(topic string, msg interface{}, expire time.Duration) chan interface{} {
	if queue, ok := mq.msgGroups[topic]; ok && queue != nil {
		go func() {
			<-time.NewTimer(expire).C
			queue <- msg
		}()
		return queue
	}
	return nil
}

func (mq *LocalMQ) SendMessage(topic string, msg interface{}, expire time.Duration) {
	if queue := mq.trySendMessage(topic, msg, expire); queue == nil {
		mq.lock.Lock()
		defer mq.lock.Unlock()
		if queue = mq.trySendMessage(topic, msg, expire); queue == nil {
			queue = make(chan interface{})
			mq.msgGroups[topic] = queue
		}
	}
}

func (mq *LocalMQ) handleError(msg interface{}, handler func(msg interface{})) {
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
		}
	}()
	handler(msg)
}

func (mq *LocalMQ) RegistryMessageHandler(topic string, handler func(msg interface{})) (error error) {
	if queue, ok := mq.msgGroups[topic]; ok && queue != nil {
		go func() {
			for msg := range queue {
				mq.handleError(msg, handler)
			}
		}()
		return
	}
	return QueueNotExist
}
