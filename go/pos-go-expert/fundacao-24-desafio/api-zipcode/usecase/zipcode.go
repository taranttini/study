package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/taranttini/study/go/post/fundacao-24-desafio/api-zipcode/entity"
)

func NewUseCaseZipcode(zipcode string) (entity.Zipcode, error) {

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	response, err := http.Get(url)
	if err != nil {
		return entity.Zipcode{}, fmt.Errorf(fmt.Sprintf("Cep: Erro ao fazer requisicao: %s\n", err.Error()))
	}
	defer response.Body.Close()

	dataRead, err := io.ReadAll(response.Body)
	if err != nil {
		return entity.Zipcode{}, fmt.Errorf(fmt.Sprintf("Cep: Erro ao ler resposta: %s\n", err.Error()))
	}

	var dataViaCEP entity.Zipcode
	err = json.Unmarshal(dataRead, &dataViaCEP)
	if err != nil {
		return entity.Zipcode{}, fmt.Errorf(fmt.Sprintf("Cep: Erro ao fazer o parser resposta: %s\n", err.Error()))
	}

	return dataViaCEP, nil
}
