package domain

import (
	"crypto/rand"
	"fmt"
	"sync"
)

/*
 * Subscriber is a struct that represents a subscriber.
 * Subscribers can subscribe to topics and receive messages
 */
type Subscriber struct {
	id       string          // id of subscriber
	messages chan *Message   // messages channel
	topics   map[string]bool // topics it is subscribed to.
	active   bool            // if given subscriber is active
	mutex    sync.RWMutex    // lock
}

/* NewSubscriber creates a new subscriber */
func NewSubscriber() (string, *Subscriber, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return "", nil, err
	}
	id := fmt.Sprintf("%X-%X", b[0:4], b[4:8])
	return id, &Subscriber{
		id:       id,
		messages: make(chan *Message),
		topics:   map[string]bool{},
		active:   true,
	}, nil
}

/* AddTopic adds a topic to the subscriber */
func (s *Subscriber) AddTopic(topic string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.topics[topic] = true
}

/* RemoveTopic removes a topic from the subscriber */
func (s *Subscriber) RemoveTopic(topic string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.topics, topic)
}

/* Topics returns all topics of the subscriber */
func (s *Subscriber) Topics() []string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	topics := []string{}
	for topic := range s.topics {
		topics = append(topics, topic)
	}
	return topics
}

/* Destruct marks the subscriber as inactive */
func (s *Subscriber) Destruct() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.active = false
	close(s.messages)
}

/* Signal sends a message to the subscriber */
func (s *Subscriber) Signal(msg *Message) {
	s.mutex.RLock()
	active := s.active
	s.mutex.RUnlock()
	if active {
		s.messages <- msg
	} else {
		close(s.messages)
	}
}

/* Listen listens to the message channel */
func (s *Subscriber) Listen() error {
	for {
		if msg, ok := <-s.messages; ok {
			fmt.Printf("Subscriber %s, received: %s from topic: %s\n", s.id, msg.Body(), msg.Topic())
		} else if !ok {
			s.mutex.RLock()
			active := s.active
			s.mutex.RUnlock()
			if active {
				return fmt.Errorf("message channel closed unexpectedly")
			}
			return nil
		}
	}
}

/* Messages returns the message channel */
func (s *Subscriber) Messages() chan *Message {
	return s.messages
}
