package domain

import (
	"sync"
)

/*
 * Broker is a message broker that manages subscribers and broadcasts messages
 */
type Broker struct {
	subscribers map[string]*Subscriber            // map of subscribers id:Subscriber
	topics      map[string]map[string]*Subscriber // map of topic to subscribers
	mut         sync.RWMutex                      // mutex lock
}

/*
 * InitBroker initializes a new broker
 */

func InitBroker() *Broker {
	return &Broker{
		subscribers: make(map[string]*Subscriber),
		topics:      make(map[string]map[string]*Subscriber),
	}
}

/*
 * AddSubscriber adds a subscriber to the broker
 */
func (b *Broker) AddSubscriber(id string, sub *Subscriber) {
	b.mut.Lock()
	defer b.mut.Unlock()
	b.subscribers[id] = sub
	for _, topic := range sub.Topics() {
		if _, ok := b.topics[topic]; !ok {
			b.topics[topic] = make(map[string]*Subscriber)
		}
		b.topics[topic][id] = sub
	}
}

/*
 * RemoveSubscriber removes a subscriber from the broker
 */
func (b *Broker) RemoveSubscriber(id string, sub *Subscriber) {
	b.mut.Lock()
	defer b.mut.Unlock()
	delete(b.subscribers, id)
	for _, topic := range sub.Topics() {
		delete(b.topics[topic], id)
	}
}

/*
 * Broadcast sends a message to a list of given topics
 */
func (b *Broker) Broadcast(topics []string, msg string) {
	for _, topic := range topics {
		b.Publish(topic, msg)
	}
}

/*
 * GetSubscribers returns a list of subscribers for a topic
 */
func (b *Broker) GetSubscribers(topic string) []*Subscriber {
	b.mut.RLock()
	defer b.mut.RUnlock()
	subs := make([]*Subscriber, 0, len(b.topics[topic]))
	for _, sub := range b.topics[topic] {
		subs = append(subs, sub)
	}
	return subs
}

/*
 * NumberOfSubscribers returns the number of subscribers
 */

func NumberOfSubscribers(b *Broker) int {
	b.mut.RLock()
	defer b.mut.RUnlock()
	return len(b.subscribers)
}

/*
 * GetSubscriber returns a subscriber by id
 */
func (b *Broker) GetSubscriber(id string) (*Subscriber, bool) {
	b.mut.RLock()
	defer b.mut.RUnlock()
	sub, ok := b.subscribers[id]
	return sub, ok
}

/*
 * GetTopics returns a list of topics
 */
func (b *Broker) GetTopics() []string {
	b.mut.RLock()
	defer b.mut.RUnlock()
	topics := make([]string, 0, len(b.topics))
	for topic := range b.topics {
		topics = append(topics, topic)
	}
	return topics
}

/*
 * Publish sends a message to all subscribers of a topic
 */
func (b *Broker) Publish(topic string, msg string) {
	b.mut.RLock()
	bTopics := b.topics[topic]
	b.mut.RUnlock()
	for _, s := range bTopics {
		m := NewMessage(msg, topic)
		if !s.active {
			return
		}
		go (func(s *Subscriber) {
			s.Signal(m)
		})(s)
	}
}
