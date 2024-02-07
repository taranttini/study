package main

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

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//http.NotFound(w, r)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, err := json.Marshal(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(result)

	// ou
	// json.NewEncoder(w).Encode(cep)

}

func BuscaCep(cep string) (*ViaCEP, error) {
	response, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%v/json/", cep))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var viaCepData ViaCEP
	err = json.Unmarshal(body, &viaCepData)
	if err != nil {
		return nil, err
	}

	return &viaCepData, nil
}
