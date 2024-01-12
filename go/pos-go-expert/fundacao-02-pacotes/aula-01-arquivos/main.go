package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// criar/sobrescrever um arquivo
	fileCreated, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso!\n")

	// escrevendo no string arquivo
	fileLength, err := fileCreated.WriteString("Ola mundo 1!\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Escrita com sucesso! Tamanho %d bytes\n", fileLength)

	// escrevendo no bytes no arquivo
	fileLength, err = fileCreated.Write([]byte("Ola mundo 2!\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Escrita com sucesso! Tamanho %d bytes\n", fileLength)

	// fechar
	err = fileCreated.Close()
	if err != nil {
		return
	}

	//
	// leitura
	//

	readFile, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\nArquivo lido com sucesso!\n\n")
	fmt.Println(string(readFile))

	//
	// leitura parcial do arquivo
	//

	// abrir arquivo
	openFile, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// criar leitor dos dados
	readBufferFile := bufio.NewReader(openFile)

	// data a ser lida
	bufferData := make([]byte, 6) // ler de 6 em 6 bytes

	for {
		dataRead, err := readBufferFile.Read(bufferData)
		if err != nil {
			// caso pare de ler, parar o processo
			break
		}
		fmt.Printf("\nleu os dados:\n  >>> %v <<<<", string(bufferData[:dataRead]))
	}

	fmt.Println("\n\nleitura finalizada!")

	//
	// exclusao do arquivo
	//
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("\n\narquivo removido com sucesso!")
}
