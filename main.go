package main

import (
	"RoadToTechlead/ownredis/connection"
	"fmt"
)

func main() {
	fmt.Println("Starte Server")
	connection.StartListener()
}
