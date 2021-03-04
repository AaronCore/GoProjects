package main

import "fmt"

type people interface {
	speak(string) string
}
type student1 struct{}

func (stu *student1) speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo people = &student1{}
	think := "bitch"
	fmt.Println(peo.speak(think))
}
