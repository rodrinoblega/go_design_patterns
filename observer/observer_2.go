package main

import "fmt"

type Subject interface {
	register(observer ObserverAux)
	deregister(observer ObserverAux)
	notifyAll()
}

type ItemAux struct {
	observers []ObserverAux
	name      string
	inStock   bool
}

func newItem(name string) *ItemAux {
	return &ItemAux{
		name: name,
	}
}

func (i *ItemAux) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *ItemAux) register(obsever ObserverAux) {
	i.observers = append(i.observers, obsever)
}

func (i *ItemAux) deregister(observer ObserverAux) {
	i.observers = removeFromSlice(i.observers, observer)
}

func (i *ItemAux) notifyAll() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func removeFromSlice(observers []ObserverAux, observerToRemove ObserverAux) []ObserverAux {
	observerListLength := len(observers)
	for i, observer := range observers {
		if observerToRemove.getId() == observer.getId() {
			observers[observerListLength-1], observers[i] = observers[i], observers[observerListLength-1]
			return observers[:observerListLength-1]
		}
	}
	return observers
}

type ObserverAux interface {
	updateValue(string)
	getId() string
}

type Customer struct {
	id string
}

func (c *Customer) updateValue(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getId() string {
	return c.id
}

func main() {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
