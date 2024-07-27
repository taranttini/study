package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func NewUseCaseCep(cep string) (ViaCEP, error) {

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	response, err := http.Get(url)
	if err != nil {
		return ViaCEP{}, fmt.Errorf(fmt.Sprintf("Cep: Erro ao fazer requisicao: %s\n", err.Error()))
	}
	defer response.Body.Close()

	dataRead, err := io.ReadAll(response.Body)
	if err != nil {
		return ViaCEP{}, fmt.Errorf(fmt.Sprintf("Cep: Erro ao ler resposta: %s\n", err.Error()))
	}

	var dataViaCEP ViaCEP
	err = json.Unmarshal(dataRead, &dataViaCEP)
	if err != nil {
		return ViaCEP{}, fmt.Errorf(fmt.Sprintf("Cep: Erro ao fazer o parser resposta: %s\n", err.Error()))
	}

	return dataViaCEP, nil
}
