package adapter

type WindowsAdapter struct {
	windows Windows
}

func (wa *WindowsAdapter) InsertIntoLightningPort() {
	println("Adapter converts Lightling signal to USB.")
	wa.windows.InsertIntoUSBPort()
}
