package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"net"
	"os"
)

func main() {

	type Parametro struct {
		XMLName xml.Name `xml:"parametro"`
		Nome    string   `xml:"nome"`
		Valor   string   `xml:"valor"`
	}

	type Parametros struct {
		XMLName   xml.Name  `xml:"parametros"`
		Parametro Parametro `xml:"parametro"`
	}

	type Metodo struct {
		XMLName    xml.Name   `xml:"metodo"`
		Nome       string     `xml:"nome"`
		Parametros Parametros `xml:"parametros"`
	}

	type Requisicao struct {
		XMLName xml.Name `xml:"requisicao"`
		Metodo  Metodo   `xml:"metodo"`
	}

	v := &Requisicao{Metodo: Metodo{Nome: "submeter", Parametros: Parametros{Parametro: Parametro{Nome: "Boletim", Valor: "xml"}}}}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")

	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	conn, _ := net.Dial("tcp", "127.0.0.1:5050")
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdout)
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}

}
