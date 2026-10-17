package main

import (
	"fmt"
	"time" 
)

type Actor struct {
	inbox chan string
}

func NewActor() *Actor {
	return &Actor{
		inbox: make(chan string),
	}
}

func (a *Actor) Receive() {
	for msg := range a.inbox {
		fmt.Println(msg)
	}
}

func (a *Actor) Send(msg string) {
	a.inbox <- msg
}



func main() {
	actor := NewActor()
	go actor.Receive()

	actor.Send("La suma del cuadrado de los catedos es igual") 
	actor.Send("Al cuadrado de la hipotenusa")  

	time.Sleep(time.Second) 

}
