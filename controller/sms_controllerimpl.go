package controller

import (
	"log"
	"math/rand"
)

func (a *AbstractSMSController) SendSMS(mobile, message string) {
	log.Printf("Sending SMS to %s: %s\n", mobile, message)
}

func (a *AbstractSMSController) CheckSMSStatus(id string) string {
	if rand.Intn(2) == 0 {
		return "success"
	}
	return "failure"
}
