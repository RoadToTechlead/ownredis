// TODO 1: Implement TCP listener; Ausnamebehandlung integrieren
// TODO 2: Accept loop; wartet aktiv mit Accept() Operation

package connection

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func StartListener() {
	ln, err := net.Listen("tcp", ":6379")

	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("Server is listening on port 6379")

	for {
		conn, err := ln.Accept() //Blockieren, bis Client anklopft
		if err != nil {
			fmt.Println("Verbindung konnte nicht akzeptiert werden.")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Befüllen des internen Byte Arrays (Puffers)
	read_ := bufio.NewReader(conn)
	
	b, err := read_.ReadByte()

	//Sicherstellen, dass die Verbindung wieder getrennt wird
	defer conn.Close()
}
