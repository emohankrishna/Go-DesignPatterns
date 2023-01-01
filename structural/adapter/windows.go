package adapter

import "fmt"

type Windows struct{}

func (windows *Windows) InsertIntoUSBPort() {
	fmt.Println("USB Connector plugged in to the Windows Machine")
}
