package adapter

func RunDemo() {
	client := &Client{}
	mac := &Mac{}
	windowsMachine := &Windows{}
	client.InsertLightningConnectorIntoComputer(mac)
	// The below don't work because it accepts Mac type object by we are passing Windows Machine
	// client.InsertLightningConnectorIntoComputer(windowsMachine)
	windowAdapter := WindowsAdaptar{
		windowsMachine: windowsMachine,
	}
	client.InsertLightningConnectorIntoComputer(&windowAdapter)
}
