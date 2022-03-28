package sms

import "fmt"

type SmsMessage struct {
	Message string
}

type SmsSender struct {
}

type SmsValidator struct {
}

type SmsGateway struct {
	ReceivedMessage string
}

func (s SmsSender) Send(msg SmsMessage, gateway *SmsGateway) {
	validator := SmsValidator{}
	if validator.Validate(msg) {
		gateway.Receive(msg)
		fmt.Println("message sent")
	} else {
		fmt.Println("message is not valid")
	}
}

func (s SmsValidator) Validate(msg SmsMessage) bool {
	return len(msg.Message) < 160
}

func (s *SmsGateway) Receive(msg SmsMessage) {
	s.ReceivedMessage = msg.Message
	fmt.Println(s.ReceivedMessage)
}
