package domain

import (
	"testing"
)

func TestAddSubscriber(t *testing.T) {
	broker := InitBroker()
	id1, sub1, err1 := NewSubscriber()
	if err1 != nil {
		t.Fatalf("Failed to create subscriber: %v", err1)
	}
	id2, sub2, err2 := NewSubscriber()
	if err2 != nil {
		t.Fatalf("Failed to create subscriber: %v", err2)
	}
	sub1.AddTopic("topic1")
	sub2.AddTopic("topic2")
	broker.AddSubscriber(id1, sub1)
	broker.AddSubscriber(id2, sub2)

	if _, ok := broker.GetSubscriber(id1); !ok {
		t.Errorf("Expected to find subscriber with id '%s'", id1)
	}

	if _, ok := broker.GetSubscriber(id2); !ok {
		t.Errorf("Expected to find subscriber with id '%s'", id2)
	}
}

func TestGetSubscribers(t *testing.T) {
	broker := InitBroker()
	id1, sub1, err1 := NewSubscriber()
	if err1 != nil {
		t.Fatalf("Failed to create subscriber: %v", err1)
	}
	id2, sub2, err2 := NewSubscriber()
	if err2 != nil {
		t.Fatalf("Failed to create subscriber: %v", err2)
	}
	sub1.AddTopic("topic1")
	sub2.AddTopic("topic2")
	broker.AddSubscriber(id1, sub1)
	broker.AddSubscriber(id2, sub2)

	if len(broker.GetSubscribers("topic1")) != 1 {
		t.Errorf("Expected to find 1 subscriber for 'topic1'")
	}

	if len(broker.GetSubscribers("topic2")) != 1 {
		t.Errorf("Expected to find 1 subscriber for 'topic2'")
	}
}

func TestGetTopics(t *testing.T) {
	broker := InitBroker()
	id1, sub1, err1 := NewSubscriber()
	if err1 != nil {
		t.Fatalf("Failed to create subscriber: %v", err1)
	}
	id2, sub2, err2 := NewSubscriber()
	if err2 != nil {
		t.Fatalf("Failed to create subscriber: %v", err2)
	}
	sub1.AddTopic("topic1")
	sub2.AddTopic("topic2")
	broker.AddSubscriber(id1, sub1)
	broker.AddSubscriber(id2, sub2)

	topics := broker.GetTopics()
	if len(topics) != 2 {
		t.Errorf("Expected to find 2 topics, got %d", len(topics))
	}
}

func TestBroadcastWithOneSubscriber(t *testing.T) {
	broker := InitBroker()
	id1, sub1, err1 := NewSubscriber()
	if err1 != nil {
		t.Fatalf("Failed to create subscriber: %v", err1)
	}
	sub1.AddTopic("topic1")
	broker.AddSubscriber(id1, sub1)

	msg := NewMessage("testBody", "topic1")
	topics := []string{"topic1"}
	broker.Broadcast(topics, msg.Body())

	receivedMsg, ok := <-sub1.Messages()
	if !ok || receivedMsg.Body() != "testBody" || receivedMsg.Topic() != "topic1" {
		t.Errorf("Failed to receive correct message from subscriber")
	}
}

func TestBroadcastWithMultipleSubscribers(t *testing.T) {
	broker := InitBroker()
	id1, sub1, err1 := NewSubscriber()
	if err1 != nil {
		t.Fatalf("Failed to create subscriber: %v", err1)
	}
	id2, sub2, err2 := NewSubscriber()
	if err2 != nil {
		t.Fatalf("Failed to create subscriber: %v", err2)
	}
	sub1.AddTopic("topic1")
	sub2.AddTopic("topic1")
	broker.AddSubscriber(id1, sub1)
	broker.AddSubscriber(id2, sub2)

	msg := NewMessage("testBody", "topic1")
	topics := []string{"topic1"}
	broker.Broadcast(topics, msg.Body())

	receivedMsg1, ok1 := <-sub1.Messages()
	if !ok1 || receivedMsg1.Body() != "testBody" || receivedMsg1.Topic() != "topic1" {
		t.Errorf("Failed to receive correct message from subscriber")
	}

	receivedMsg2, ok2 := <-sub2.Messages()
	if !ok2 || receivedMsg2.Body() != "testBody" || receivedMsg2.Topic() != "topic1" {
		t.Errorf("Failed to receive correct message from subscriber")
	}
}

func TestBroadcastWithMultipleTopics(t *testing.T) {
	broker := InitBroker()
	id1, sub1, err1 := NewSubscriber()
	if err1 != nil {
		t.Fatalf("Failed to create subscriber: %v", err1)
	}

	sub1.AddTopic("topic1")
	sub1.AddTopic("topic2")

	broker.AddSubscriber(id1, sub1)

	msg1 := NewMessage("testBody1", "topic1")
	msg2 := NewMessage("testBody2", "topic2")

	broker.Broadcast([]string{"topic1"}, msg1.Body())
	broker.Broadcast([]string{"topic2"}, msg2.Body())

	receivedMsg1, ok1 := <-sub1.Messages()
	if !ok1 || receivedMsg1.Body() != "testBody2" || receivedMsg1.Topic() != "topic2" {
		t.Log(receivedMsg1)
		t.Errorf("Failed to receive correct message from subscriber")
	}

	receivedMsg2, ok2 := <-sub1.Messages()
	if !ok2 || receivedMsg2.Body() != "testBody1" || receivedMsg2.Topic() != "topic1" {
		t.Log(receivedMsg2)
		t.Errorf("Failed to receive correct message from subscriber")
	}
}

func TestRemoveSubscriber(t *testing.T) {
	broker := InitBroker()
	id1, sub1, err1 := NewSubscriber()
	if err1 != nil {
		t.Fatalf("Failed to create subscriber: %v", err1)
	}
	id2, sub2, err2 := NewSubscriber()
	if err2 != nil {
		t.Fatalf("Failed to create subscriber: %v", err2)
	}
	sub1.AddTopic("topic1")
	sub2.AddTopic("topic2")
	broker.AddSubscriber(id1, sub1)
	broker.AddSubscriber(id2, sub2)

	broker.RemoveSubscriber(id1, sub1)
	if _, ok := broker.GetSubscriber(id1); ok {
		t.Errorf("Expected not to find subscriber with id '%s'", id1)
	}

	if NumberOfSubscribers(broker) != 1 {
		t.Errorf("Expected 1 subscriber, got %d", NumberOfSubscribers(broker))
	}
}
