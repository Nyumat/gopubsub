package domain

// Message represents a message with a topic and body.
type Message struct {
	topic string
	body  string
}

// NewMessage creates a new Message and returns its pointer.
func NewMessage(body string, topic string) *Message {
	return &Message{
		topic: topic,
		body:  body,
	}
}

// Topic returns the topic of the message.
func (m *Message) Topic() string {
	return m.topic
}

// Body returns the body of the message.
func (m *Message) Body() string {
	return m.body
}
