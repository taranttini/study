package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisicao: %v\n", err)
		}
		defer response.Body.Close()

		dataRead, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}
		var dataViaCEP ViaCEP
		err = json.Unmarshal(dataRead, &dataViaCEP)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parser resposta: %v\n", err)
		}

		fmt.Println(dataViaCEP)

		file, err := os.Create(fmt.Sprintf("cep_%s.txt", cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao gerar arquivo: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(
			fmt.Sprintf("CEP: %s, Cidade: %s, UF: %s, Logradouro: %s / ",
				dataViaCEP.Cep,
				dataViaCEP.Localidade,
				dataViaCEP.Uf,
				dataViaCEP.Logradouro,
			))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao gravar no arquivo: %v\n", err)
		}

		fmt.Printf("Arquivo de cep gerado com sucesso\n")
	}
}
