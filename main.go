package main

import (
	"fmt"
	"time"
)

type Package interface {
	GetSender() string
	GetRecipient() string
	Send()
}

type Box struct {
	Sender    string
	Recipient string
}

func (b Box) GetSender() string {
	return b.Sender
}

func (b Box) GetRecipient() string {
	return b.Recipient
}

func (b Box) Send() {
	fmt.Printf("Відправлення коробки від %s до %s кур'єром\n", b.Sender, b.Recipient)
}

type Envelope struct {
	Sender    string
	Recipient string
}

func (e Envelope) GetSender() string {
	return e.Sender
}

func (e Envelope) GetRecipient() string {
	return e.Recipient
}

func (e Envelope) Send() {
	fmt.Printf("Відправлення конверта від %s до %s листоношою\n", e.Sender, e.Recipient)
}

type SortingDepartment struct{}

func (sd SortingDepartment) SendPackage(p Package) {
	fmt.Println("Сортування...")
	fmt.Printf("\nГотуємо відпавку для %s...\n", p.GetRecipient())
	time.Sleep(2 * time.Second)
	fmt.Printf("\nВизначаємо адресу %s...\n", p.GetRecipient())
	time.Sleep(time.Second)
	fmt.Printf("\nПередаємо кур'єру...\n")
	time.Sleep(time.Second)

	p.Send()
}

func main() {
	box := Box{
		Sender:    "John",
		Recipient: "Alice",
	}

	envelope := Envelope{
		Sender:    "Bob",
		Recipient: "Eve",
	}

	sortingDept := SortingDepartment{}
	sortingDept.SendPackage(box)
	sortingDept.SendPackage(envelope)
}
