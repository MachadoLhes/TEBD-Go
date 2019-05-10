package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func req(archiveName string) {
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
	for message.Scan() {
		line := message.Text()
		fmt.Fprintf(os.Stdout, line+"\n")
	}
}

func main() {
	opcao := 0

	for opcao != 3 {

		fmt.Println("Você deseja submeter ou consultar?\n (1) submeter / (2) consultar / (3) sair")
		fmt.Scanf("%d", &opcao)

		if opcao == 1 {
			req("submeter.xml")
		} else if opcao == 2 {
			req("consultaStatus.xml")
		} else if opcao == 3 {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("selecione uma opção valida")
		}
	}

}
