package chainOfResponsibility

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(patient *Patient) {
	if patient.paymentDone {
		fmt.Println("Payment done")
	} else {
		fmt.Println("Cashier getting money from patient ", patient.name)
		patient.paymentDone = true
	}
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
