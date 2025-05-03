package main

import (
	"fmt"
)

// Structs

type EventType int

type Event struct {
	ID   EventType
	Data map[string]interface{}
}

type commandType int

type command struct {
	typ       commandType
	topic     string
	event     Event
	replyChan chan interface{} // we use interface{} to multiplex return values
}

// consts

const (
	cmdNewTopic commandType = iota
	cmdSubscribe
	cmdPublish
)

// The event bus itself

type EventBus struct {
	cmdChan chan command
}

func NewEventBus() *EventBus {
	e := &EventBus{
		cmdChan: make(chan command),
	}
	go e.run()
	return e
}

func (e *EventBus) run() {
	topics := make(map[string][]chan Event)
	for cmd := range e.cmdChan {

		switch cmd.typ {

		case cmdNewTopic:
			if _, ok := topics[cmd.topic]; !ok {
				topics[cmd.topic] = make([]chan Event, 0)
				cmd.replyChan <- nil
			} else {
				cmd.replyChan <- fmt.Errorf("topic already exists")
			}

		case cmdSubscribe:
			if _, ok := topics[cmd.topic]; !ok {
				cmd.replyChan <- fmt.Errorf("topic does not exist")
			}
			ch := make(chan Event, 10)
			topics[cmd.topic] = append(topics[cmd.topic], ch)
			cmd.replyChan <- ch

		case cmdPublish:
			if channels, ok := topics[cmd.topic]; ok {
				for _, ch := range channels {
					select {
					case ch <- cmd.event:
					default:
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
