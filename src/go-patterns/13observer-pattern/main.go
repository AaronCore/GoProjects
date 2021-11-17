package main

// 观察者模式
func main() {
	shirtItem := newItem("GoLang Design Patterns")
	observerFirst := &customer{id: "robin@zhyea.com"}
	observerSecond := &customer{id: "golang@zhyea.com"}
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.updateAvailability()
}
