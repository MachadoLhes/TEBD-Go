package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	var opcao int
	var archiveName string

	fmt.Println("Você deseja submeter ou consultar?\n (1) submeter / (2) consultar")

	fmt.Scanf("%d", &opcao)

	if opcao == 1 {
		archiveName = "submeter.xml"
	} else if opcao == 2 {
		archiveName = "consultaStatus.xml"
	} else {
		fmt.Errorf("selecione uma opção valida")
	}

	xmlFile, err := os.Open(archiveName)

	if err != nil {
		fmt.Errorf("ERROR")
	}

	conn, _ := net.Dial("tcp", "127.0.0.1:5050")

	scanner := bufio.NewScanner(xmlFile) // Loop over all lines in the file and print them.

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(conn, line+"\n")
	}

	message := bufio.NewScanner(conn)
	fmt.Print("Message from server: \n")
	for message.Scan() {
		line := message.Text()
		fmt.Fprintf(os.Stdout, line+"\n")
	}

}
