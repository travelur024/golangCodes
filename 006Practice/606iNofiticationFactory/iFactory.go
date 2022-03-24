package main

import "fmt"

//Abstrac factory
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

//Abstract Product
type ISender interface {
	GetSenderMethod() string  //sms or email
	GetSenderChannel() string //service
}

//Concrete product
type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

//SMS  Isender
type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Service Twilio"
}

//email Concrete product
type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

//Part of concrete product
type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

//Create iNotificationFactory concreta
func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No Notification Type")
}

//Print method with interface
func getMethod(f INotificationFactory) {
	fmt.Println("The methdo you use is:", f.GetSender().GetSenderMethod())
}

//Send notification
func SendNotification(f INotificationFactory) {
	f.SendNotification()
}

func main() {
	smsFactory, err01 := getNotificationFactory("SMS")

	if err01 == nil {
		getMethod(smsFactory)
		SendNotification(smsFactory)
	} else {
		fmt.Println(err01.Error())
	}
	emailFactory, err02 := getNotificationFactory("Email")
	if err02 == nil {
		getMethod(emailFactory)
		SendNotification(emailFactory)
	} else {
		fmt.Println(err02.Error())
	}

}
