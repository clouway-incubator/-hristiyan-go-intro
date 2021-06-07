package sms_test

import (
	"fmt"
	"hristiyan-go-intro/sms"
	"testing"
)

func TestValidateShouldReturnTrueWhenUnder160Characters(t *testing.T) {
	msg := sms.SmsMessage{}
	msg.Message = "adkdpqwqdqowpdqd"
	validator := sms.SmsValidator{}
	actual := validator.Validate(msg)
	if !actual {
		t.Errorf("expected 'true', got %t", actual)
	}
}

func TestValidateShouldReturnFalseWhenAbove160Characters(t *testing.T) {
	msg := sms.SmsMessage{}
	msg.Message = "adkdpqwqdqowpdqdadkdpqwqdqowpdqdadkdpqwqdqowpdqdadkdpqwqdqowpdqdadkdpqwqdqowpdqdadkdpqwqdqowpdqdadkdpqwqdqowpdqdadkdpqwqdqowpdqdadkdpqwqdqowpdqdadkdpqwqdqowpdqd"
	validator := sms.SmsValidator{}
	actual := validator.Validate(msg)
	fmt.Print(actual)
	if actual {
		t.Errorf("expected 'false', got %t", actual)
	}
}

func TestReceive(t *testing.T) {
	msg := sms.SmsMessage{"msg"}
	gateway := sms.SmsGateway{}
	gateway.Receive(msg)
	actual := gateway.ReceivedMessage
	if actual != msg.Message {
		t.Errorf("expected %s, got %s", msg.Message, actual)
	}
}

func TestSend(t *testing.T) {
	msg := sms.SmsMessage{"msg"}
	sender := sms.SmsSender{}
	gateway := sms.SmsGateway{}
	sender.Send(msg, &gateway)
	validator := sms.SmsValidator{}
	isValid := validator.Validate(msg)
	if !isValid {
		t.Errorf("expected 'true', got %t", isValid)
	}
	if gateway.ReceivedMessage != msg.Message {
		t.Errorf("expected %s, got %s", msg.Message, gateway.ReceivedMessage)
	}
}
