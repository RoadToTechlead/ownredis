// TODO 1: Implement TCP listener; Ausnamebehandlung integrieren
// TODO 2: Accept loop; wartet aktiv mit Accept() Operation

package connection

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
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
			continue // damit
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	/*
		Sicherstellen, dass die Verbindung wieder getrennt wird. Mit defer wird
		sichergestellt, dass conn.Close() aufgerufen wird, unabhängig vom Ergebnis
	*/
	defer conn.Close()
	// Befüllen des internen Byte Arrays (Puffers)
	read_ := bufio.NewReader(conn)
	//Anzahl der Elemente anzeigen
	b, err := read_.ReadString('\n')
	result, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(b, "*"))) // * und \r\n entfernen
	//zeigt ungefilterten/quoted String
	if err != nil {
		fmt.Println("String result konnte nicht geparst werden")
	}
	// loop
	for i := 0; i <= result; i += 1 {

	}
	//if err != nil {
	//	fmt.Println("Erstes Byte konnte nicht gelesen werden.")
	//}
	fmt.Println(result)
}
