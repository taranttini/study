package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/taranttini/study/go/pos-go-expert/fundacao-23-desafio/internal/entity"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-23-desafio/internal/usecase"
)

func main() {
	//RequestWeatherByZipCode("02765070")
	http.HandleFunc("/", WeatherHandler)
	fmt.Println("Listen on :8080")
	http.ListenAndServe(":8080", nil)
}

func RequestWeatherByZipCode(zipcode string) (*entity.Temperature, *entity.CustomError) {

	if !usecase.NewValidateCep(zipcode) {
		return nil, entity.ErrorZipcodeInvalid()
	}

	cep, err := usecase.NewUseCaseCep(zipcode)

	if err != nil {
		println(err.Error())
		return nil, entity.ErrorInternal()
	}

	if len(cep.Cep) == 0 {
		return nil, entity.ErrorZipcodeNotFound()
	}

	// temporary key
	wheather, err := usecase.NewUseCaseWeather(cep.Localidade, cep.Uf, "2d4fc12c862a4c3bbd5234402242607")

	if err != nil {
		println(err.Error())
		return nil, entity.ErrorInternal()
	}

	temp := entity.TemperatureResponse(wheather.Current.TempCelcius)

	return temp, nil
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	zipcode := r.URL.Query().Get("zipcode")
	// if zipcode == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	cep, requestError := RequestWeatherByZipCode(zipcode)

	//err != nil
	if requestError != nil {
		w.WriteHeader(requestError.Code)
		w.Write([]byte(requestError.Message))
		return
	}

	result, err := json.Marshal(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

	// ou
	// json.NewEncoder(w).Encode(cep)

}
