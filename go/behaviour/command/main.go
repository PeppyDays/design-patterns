package command

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}
	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	offButton := &Button{
		command: offCommand,
	}

	onButton.press()
	offButton.press()
}
