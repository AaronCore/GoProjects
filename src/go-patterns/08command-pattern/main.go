package main

// 命令模式
func main() {
	tv := &tv{}
	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}

	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
