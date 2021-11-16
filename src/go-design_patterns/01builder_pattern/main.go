package main

import "fmt"

//建造者模式
//建造者模式是一种创建型模式，主要用来创建比较复杂的对象。
//建造者模式的使用场景：
//建造者模式通常适用于有多个构造器参数或者需要较多构建步骤的场景。使用建造者模式可以精简构造器参数的数量，让构建过程更有条理。
//可以为同一个产品提供两个不同的实现。比如，在下面的代码中，为house类型创建了两个不同的实现：通过iglooBuilder构建的igloo（冰屋），以及通过cabinBuilder构建的cabin（木屋）
//可以应用于构建过程不允许被中断的场景。仍然参考刚才的代码，house类型的对象要么彻底完成，要么压根没有创建，不会存在中间状态，这是因为struct director封装了相应的过程，中间状态仅存在于ConcreteBuilder中。
func main() {
	cabinBuilder := getBuilder("cabin")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(cabinBuilder)
	cabinHouse := director.buildHouse()

	fmt.Printf("Cabin House Door Type: %s\n", cabinHouse.doorType)
	fmt.Printf("Cabin House Window Type: %s\n", cabinHouse.windowType)
	fmt.Printf("Cabin House Num Floor: %d\n", cabinHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
