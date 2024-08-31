package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/taranttini/study/go/post/fundacao-24-desafio/api-weather/entity"
)

func NewUseCaseWeather(city string, uf string, apiKey string) (entity.Weather, error) {

	//fmt.Printf("%s - %s\n", url.QueryEscape(city), uf)

	urlEndpoint := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&aqi=no&q=%s,%s", apiKey, url.QueryEscape(city), uf)

	response, err := http.Get(urlEndpoint)
	if err != nil {
		return entity.Weather{}, fmt.Errorf(fmt.Sprintf("Weather: Erro ao fazer requisicao: %s\n", err.Error()))
	}
	defer response.Body.Close()

	dataRead, err := io.ReadAll(response.Body)
	if err != nil {
		return entity.Weather{}, fmt.Errorf(fmt.Sprintf("Weather: Erro ao ler resposta: %s\n", err.Error()))
	}

	var dataWeather entity.Weather
	err = json.Unmarshal(dataRead, &dataWeather)
	if err != nil {
		return entity.Weather{}, fmt.Errorf(fmt.Sprintf("Weather: Erro ao fazer o parser resposta: %s\n", err.Error()))
	}

	return dataWeather, nil
}
