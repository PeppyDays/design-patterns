package adapter

func main() {
	client := &Client{}

	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	windows := Windows{}
	// errors cause windows doesn't implement Computer interface
	// client.InsertLightningConnectorIntoComputer(&windows)
	windowsAdapter := &WindowsAdapter{windows}
	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}
