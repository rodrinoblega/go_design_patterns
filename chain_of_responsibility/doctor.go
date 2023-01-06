package main

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) execute(patient *Patient) {
	if patient.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(patient)
		return
	}
	fmt.Println("Doctor checking patient")
	patient.doctorCheckUpDone = true
	d.next.execute(patient)
}

func (d *Doctor) setNext(department Department) {
	d.next = department
}
