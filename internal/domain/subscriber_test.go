package domain

import (
	"testing"
	"time"
)

func TestNewSubscriber(t *testing.T) {
	_, sub, err := NewSubscriber()
	if err != nil {
		t.Fatalf("Failed to create subscriber: %v", err)
	}
	if sub == nil {
		t.Errorf("Failed to create a new subscriber")
	}
}

func TestAddTopic(t *testing.T) {
	_, sub, _ := NewSubscriber()
	sub.AddTopic("testTopic")
	if !sub.topics["testTopic"] {
		t.Errorf("Failed to add topic to subscriber")
	}
}

func TestRemoveTopic(t *testing.T) {
	_, sub, _ := NewSubscriber()
	sub.AddTopic("testTopic")
	sub.RemoveTopic("testTopic")
	if sub.topics["testTopic"] {
		t.Errorf("Failed to remove topic from subscriber")
	}
}

func TestSignal(t *testing.T) {
	_, sub, _ := NewSubscriber()
	msg := NewMessage("testBody", "testTopic")
	go sub.Signal(msg)
	receivedMsg, ok := <-sub.messages
	if !ok || receivedMsg.Body() != "testBody" || receivedMsg.Topic() != "testTopic" {
		t.Errorf("Failed to receive correct message from subscriber")
	}
}

func TestDestruct(t *testing.T) {
	_, sub, _ := NewSubscriber()
	sub.Destruct()
	if sub.active {
		t.Errorf("Failed to destruct subscriber")
	}
}

func TestSignalAfterDestruct(t *testing.T) {
	_, sub, _ := NewSubscriber()
	sub.Destruct()
	go func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		msg := NewMessage("testBody", "testTopic")
		sub.Signal(msg)
	}()
	time.Sleep(1 * time.Second)
}
