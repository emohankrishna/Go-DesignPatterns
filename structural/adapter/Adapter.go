package adapter

import "fmt"

type WindowsAdaptar struct {
	windowsMachine *Windows
}

func (wa *WindowsAdaptar) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	wa.windowsMachine.InsertIntoUSBPort()
}
