package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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

func init() {
	viper.AutomaticEnv()
}

func NewUseCaseCep(cep string, ctx context.Context) (ViaCEP, error) {

	_, span := otel.Tracer(viper.GetString("SPAN_TRACE_NAME")).Start(ctx, "start-New-Use-Case-Cep")
	span.SetAttributes(attribute.String("cep", cep))
	defer span.End()

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

	_, span1 := otel.Tracer(viper.GetString("SPAN_TRACE_NAME")).Start(ctx, "end-New-Use-Case-Cep")
	defer span1.End()

	return dataViaCEP, nil
}
