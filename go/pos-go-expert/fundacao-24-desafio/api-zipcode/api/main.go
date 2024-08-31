package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/taranttini/study/go/post/fundacao-24-desafio/api-zipcode/usecase"
)

func main() {

	http.HandleFunc("/", ZipCodeHandler)
	fmt.Println("Listen on :8080")
	http.ListenAndServe(":8080", nil)
}

func ZipCodeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	queryZipcode := r.URL.Query().Get("zipcode")
	queryOnly := r.URL.Query().Get("only")

	if !ValidateZipCode(queryZipcode) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}

	zipcode, err := usecase.NewUseCaseZipcode(queryZipcode)

	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(zipcode.Cep) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))
		return
	}

	if strings.ToUpper(queryOnly) == "ZIPCODE" {
		result, err := json.Marshal(zipcode)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	}

	temperature, err := ProcessWeather(zipcode.Localidade, zipcode.Uf, w, r)

	if err != nil {
		fmt.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	temperatureResult, err := json.Marshal(temperature)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(temperatureResult)
}

func ValidateZipCode(zipcode string) bool {
	if zipcode == "" {
		return false
	}
	if len(zipcode) != 8 {
		return false
	}
	_, err := strconv.Atoi(zipcode)

	return err == nil
}

type Temperature struct {
	Celcius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func ProcessWeather(city string, uf string, w http.ResponseWriter, r *http.Request) (Temperature, error) {

	urlEndpoint := fmt.Sprintf("http://app-weather:8090/?city=%s&uf=%s", url.QueryEscape(city), uf)

	response, err := http.Get(urlEndpoint)
	if err != nil {
		return Temperature{}, fmt.Errorf(fmt.Sprintf("Zipcode Weather: Erro ao fazer requisicao: %s\n", err.Error()))
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		w.WriteHeader(response.StatusCode)
		dataReadIn, err := io.ReadAll(response.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return Temperature{}, nil
		}
		w.Write(dataReadIn)
		return Temperature{}, nil
	}

	dataRead, err := io.ReadAll(response.Body)
	if err != nil {
		return Temperature{}, fmt.Errorf(fmt.Sprintf("Zipcode Weather: Erro ao ler resposta: %s\n", err.Error()))
	}

	var dataTemperature Temperature
	err = json.Unmarshal([]byte(dataRead), &dataTemperature)
	if err != nil {
		return Temperature{}, fmt.Errorf(fmt.Sprintf("Zipcode Weather: Erro ao fazer o parser resposta: %s\n", err.Error()))
	}
	return dataTemperature, nil
}
