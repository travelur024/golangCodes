//Abstract factory -Patron de diseno para crear familias de
//objetos
//Patron de diseno - plano para resolver problemas de software
package main

import "fmt"

type INotificationFactory interface {
	SendNotification()  //Send nofitication 1
	GetSender() ISender //Return metodo y canal 2
}

type ISender interface {
	GetSenderMethod() string  //Return method2.1
	GetSenderChannel() string //Return channel2.2
}

//sms
type SMSNotification struct {
}

//1
func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

//2
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

//2.1
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

//2.2
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

//email
type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

//2.1
func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

//2.2
func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

//Crear funcnion que retorne la clase concreta que vamos
//a utilizar para enviar notificaciones

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No Notification Type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")
	sendNotification(smsFactory)
	sendNotification(emailFactory)
	getMethod(smsFactory)
	getMethod(emailFactory)
}
