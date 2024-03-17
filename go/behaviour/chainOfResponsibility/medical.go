package chainOfResponsibility

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
	} else {
		fmt.Println("Medical giving medicine to patient")
		p.medicineDone = true
	}
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}
