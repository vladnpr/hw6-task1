package main

import "fmt"

type Sender interface {
	Send()
}

type Courier struct {
	parcel *Box
}

func (c *Courier) Send() {
	fmt.Printf("Відправлення посилки від %s до %s поштою\\n", c.parcel.GetSender(), c.parcel.GetRecipient())
}

type Postman struct {
	parcel *Envelope
}

func (p *Postman) Send() {
	fmt.Printf("Відправлення листа від %s до %s поштою\\n", p.parcel.GetSender(), p.parcel.GetRecipient())
}

type Package interface {
	GetSender() string
	GetRecipient() string
}

type Box struct {
	senderAddress string
	recipient     string
}

func (b *Box) GetSender() string {
	return b.senderAddress
}

func (b *Box) GetRecipient() string {
	return b.recipient
}

type Envelope struct {
	senderAddress string
	recipient     string
}

func (e *Envelope) GetSender() string {
	return e.senderAddress
}

func (e *Envelope) GetRecipient() string {
	return e.recipient
}

type SortingDepartment struct {
}

func (sd *SortingDepartment) SendPackage(p Package) {
	var sender Sender
	switch p.(type) {
	case *Box:
		fmt.Println("Сортувальний відділ: Виявлено коробку")
		sender = &Courier{parcel: p.(*Box)}
		sender.Send()
	case *Envelope:
		fmt.Println("Сортувальний відділ: Виявлено конверт")
		sender = &Postman{parcel: p.(*Envelope)}
		sender.Send()
	default:
		fmt.Println("Сортувальний відділ: Невідомий тип посилки")
	}
}

func main() {
	box := Box{
		senderAddress: "John",
		recipient:     "Alice",
	}

	envelope := Envelope{
		senderAddress: "Bob",
		recipient:     "Eve",
	}

	sortingDept := SortingDepartment{}
	sortingDept.SendPackage(&box)
	sortingDept.SendPackage(&envelope)
}
