package eventbus

import (
	"fmt"
)

type EventType int

type Event struct {
	ID   EventType
	Data map[string]interface{}
}

type cmdType int

const (
	cmdNewTopic cmdType = iota
	cmdSubscribe
	cmdPublish
)

type command struct {
	typ       cmdType
	topic     string
	event     Event
	replyChan chan interface{} // we use interface{} to multiplex return values
}

type EventBus struct {
	cmdChan chan command
}

func NewEventBus() *EventBus {
	b := &EventBus{
		cmdChan: make(chan command),
	}
	go b.run()
	return b
}

func (e *EventBus) run() {
	topics := make(map[string][]chan Event)
	for cmd := range e.cmdChan {

		switch cmd.typ {

		case cmdSubscribe:

		case cmdNewTopic:

		case cmdPublish:
			if channels, ok := topics[topic]; ok {
				for _, ch := range channels {
					select {
					case ch <- event:
					default:
						// Optionally drop or log if channel is full
					}
				}
			}

		}
	}
}

func (e *EventBus) NewTopic(name string) error {
	resp := make(chan interface{})
	e.cmdChan <- command{typ: cmdNewTopic, topic: name, replyChan: resp}
	res := <-resp
	if err, ok := res.(error); ok {
		return fmt.Errorf("error creating topic: %s", err)
	}
	return nil
}

func (e *EventBus) Subscribe(topic string) (chan Event, error) {
	resp := make(chan interface{})
	e.cmdChan <- command{typ: cmdSubscribe, topic: topic, replyChan: resp}
	res := <-resp
	if ch, ok := res.(chan Event); ok {
		return ch, nil
	}
	return nil, fmt.Errorf("subscribe failed for topic %q", topic)
}

func (e *EventBus) Publish(topic string, event Event) {
	e.cmdChan <- command{typ: cmdPublish, topic: topic, event: event}
}
