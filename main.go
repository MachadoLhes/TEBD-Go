package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	archiveName := "submeter.xml"

	xmlFile, err := os.Open(archiveName)

	if err != nil {
		fmt.Errorf("ERROR")
	}

	conn, _ := net.Dial("tcp", "127.0.0.1:5050")
	// read in input from file
	scanner := bufio.NewScanner(xmlFile) // Loop over all lines in the file and print them.
	// send to socket
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(conn, line+"\n")
	}

	// listen for reply
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)

}
