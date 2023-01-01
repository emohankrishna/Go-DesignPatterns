package adapter

import "fmt"

type Mac struct{}

func (mac *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightening Connector is plugged into Mac machine")
}
