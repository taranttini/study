package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/taranttini/study/go/post/fundacao-24-desafio/api-weather/entity"
	"github.com/taranttini/study/go/post/fundacao-24-desafio/api-weather/usecase"
)

func main() {
	http.HandleFunc("/", WeatherHandler)
	fmt.Println("Listen on :8090")
	http.ListenAndServe(":8090", nil)
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	queryCity := r.URL.Query().Get("city")
	queryUf := r.URL.Query().Get("uf")

	if !ValidateLocation(queryCity, queryUf) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid location"))
		return
	}

	state := map[string]string{
		"AC": "acre",
		"AL": "alagoas",
		"AP": "amapa",
		"AM": "amazonas",
		"BA": "bahia",
		"CE": "ceara",
		"DF": "distrito-federal",
		"ES": "espirito-Santo",
		"GO": "goias",
		"MA": "maranhao",
		"MT": "mato-grosso",
		"MS": "mato-grosso-do-sul",
		"MG": "minas-gerais",
		"PA": "para",
		"PB": "paraiba",
		"PR": "parana",
		"PE": "pernambuco",
		"PI": "piaui",
		"RJ": "rio-de-janeiro",
		"RN": "rio-grande-do-norte",
		"RS": "rio-grande-do-sul",
		"RO": "rondonia",
		"RR": "roraima",
		"SC": "santa-catarina",
		"SP": "sao-paulo",
		"SE": "sergipe",
		"TO": "tocantins",
	}

	uf := state[strings.ToUpper(queryUf)]

	apiKey := "2d4fc12c862a4c3bbd5234402242607"

	weather, err := usecase.NewUseCaseWeather(queryCity, uf, apiKey)

	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(weather.Location.Name) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))
		return
	}

	temp := entity.TemperatureResponse(weather.Current.TempCelcius)

	result, err := json.Marshal(temp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func ValidateLocation(city string, uf string) bool {
	if city == "" || uf == "" {
		return false
	}
	if len(city) < 3 || len(uf) != 2 {
		return false
	}
	return true
}
