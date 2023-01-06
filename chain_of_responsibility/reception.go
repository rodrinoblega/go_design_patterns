package main

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) execute(patient *Patient) {
	if patient.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(patient)
		return
	}
	fmt.Println("Reception registering patient")
	patient.registrationDone = true
	r.next.execute(patient)
}

func (r *Reception) setNext(department Department) {
	r.next = department
}
