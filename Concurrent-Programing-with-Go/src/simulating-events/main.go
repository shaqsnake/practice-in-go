package main

import (
	"fmt"
)

func main() {
	btn := MakeButton()
	handlerOne := make(chan string)
	handlerTwo := make(chan string)

	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)

	go func() {
		for {
			msg := <-handlerOne
			fmt.Println("Handler One: " + msg)
		}
	}()

	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println("Handler Two: " + msg)
		}
	}()

	btn.TriggerEvent("click", "Button clicked!")
	btn.RemoveEventLsterner("click", handlerTwo)
	btn.TriggerEvent("click", "Button clicked again!")

	fmt.Scanln()

}

type Button struct {
	eventListeners map[string][]chan string
}

func MakeButton() *Button {
	result := new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result
}

func (this *Button) AddEventListener(event string, responseChannel chan string) {
	if _, exists := this.eventListeners[event]; exists {
		this.eventListeners[event] = append(this.eventListeners[event], responseChannel)
	} else {
		this.eventListeners[event] = []chan string{responseChannel}
	}
}

func (this *Button) RemoveEventLsterner(event string, listenChannel chan string) {
	if _, exists := this.eventListeners[event]; exists {
		for idx := range this.eventListeners[event] {
			if this.eventListeners[event][idx] == listenChannel {
				this.eventListeners[event] = append(this.eventListeners[event][:idx], this.eventListeners[event][idx+1:]...)
				break
			}
		}
	}
}

func (this *Button) TriggerEvent(event string, response string) {
	if _, exists := this.eventListeners[event]; exists {
		for _, handler := range this.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}