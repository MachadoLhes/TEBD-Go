package main

import (
	"encoding/xml"
	"fmt"
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

}
