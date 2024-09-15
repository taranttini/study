package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type Weather struct {
	Location WeatherLocation `json:"location"`
	Current  WeatherCurrent  `json:"current"`
}

type WeatherLocation struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzId           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type WeatherCurrent struct {
	TempCelcius    float64 `json:"temp_c"`
	TempFahrenheit float64 `json:"temp_f"`
}

func init() {
	viper.AutomaticEnv()
}

func NewUseCaseWeather(city string, uf string, apiKey string, ctx context.Context) (*Weather, error) {

	_, span := otel.Tracer(viper.GetString("SPAN_TRACE_NAME")).Start(ctx, "start-New-Use-Case-Weather")
	span.SetAttributes(attribute.String("city", city))
	span.SetAttributes(attribute.String("uf", uf))
	defer span.End()

	urlEndpoint := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&aqi=no&q=%s,%s", apiKey, url.QueryEscape(city), uf)

	response, err := http.Get(urlEndpoint)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Weather: Erro ao fazer requisicao: %s\n", err.Error()))
	}
	defer response.Body.Close()

	dataRead, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Weather: Erro ao ler resposta: %s\n", err.Error()))
	}

	var dataWeather Weather
	err = json.Unmarshal(dataRead, &dataWeather)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Weather: Erro ao fazer o parser resposta: %s\n", err.Error()))
	}

	_, span1 := otel.Tracer(viper.GetString("SPAN_TRACE_NAME")).Start(ctx, "end-New-Use-Case-Weather")
	defer span1.End()

	return &dataWeather, nil
}
