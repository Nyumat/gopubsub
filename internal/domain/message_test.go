package domain

import (
	"testing"
)

func TestNewMessage(t *testing.T) {
	msg := NewMessage("testBody", "testTopic")
	if msg.Topic() != "testTopic" {
		t.Errorf("Expected topic to be 'testTopic', got %s", msg.Topic())
	}
	if msg.Body() != "testBody" {
		t.Errorf("Expected body to be 'testBody', got %s", msg.Body())
	}
}

func TestNewMessageWithEmptyBody(t *testing.T) {
	msg := NewMessage("", "testTopic")
	if msg.Topic() != "testTopic" {
		t.Errorf("Expected topic to be 'testTopic', got %s", msg.Topic())
	}
	if msg.Body() != "" {
		t.Errorf("Expected body to be '', got %s", msg.Body())
	}
}

func TestNewMessageWithEmptyTopic(t *testing.T) {
	msg := NewMessage("testBody", "")
	if msg.Topic() != "" {
		t.Errorf("Expected topic to be '', got %s", msg.Topic())
	}
	if msg.Body() != "testBody" {
		t.Errorf("Expected body to be 'testBody', got %s", msg.Body())
	}
}

func TestNewMessageWithEmptyBodyAndTopic(t *testing.T) {
	msg := NewMessage("", "")
	if msg.Topic() != "" {
		t.Errorf("Expected topic to be '', got %s", msg.Topic())
	}
	if msg.Body() != "" {
		t.Errorf("Expected body to be '', got %s", msg.Body())
	}
}
