package controller

type SMSController interface {
	SendSMS(mobile, message string) string
	CheckSMSStatus(id string) string
}

type AbstractSMSController struct{}
