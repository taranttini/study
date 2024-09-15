package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
	"github.com/taranttini/study/go/post/fundacao-24-desafio/api-temperature/entity"
	"github.com/taranttini/study/go/post/fundacao-24-desafio/api-temperature/usecase"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"

	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func init() {
	viper.AutomaticEnv()
}

// #############
// #############
//
//
//
// #############
// #############

func initTracer(serviceName string) {

	exporter, err := zipkin.New(viper.GetString("ZIPKIN_ENDPOINT"))
	if err != nil {
		log.Fatalf("Fail to create Zipkin exporter: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
}

func main() {
	initTracer(viper.GetString("SPAN_TRACE_NAME"))

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/zipcode", ZipCodeHandler)
	r.Get("/weather", WeatherHandler)

	port := viper.GetString("HTTP_PORT")
	fmt.Printf("Listen on %s\n", port)
	http.ListenAndServe(port, r)
}

func RequestWeatherByZipCode(zipcode string, ctx context.Context) (*entity.Temperature, *entity.CustomError) {

	if !usecase.NewValidateCep(zipcode) {
		return nil, entity.ErrorZipcodeInvalid()
	}

	cep, err := usecase.NewUseCaseCep(zipcode, ctx)

	if err != nil {
		println(err.Error())
		return nil, entity.ErrorInternal()
	}

	if len(cep.Cep) == 0 {
		return nil, entity.ErrorZipcodeNotFound()
	}

	// temporary key
	wheather, err := usecase.NewUseCaseWeather(cep.Localidade, cep.Uf, viper.GetString("TOKEN"), ctx)

	if err != nil {
		println(err.Error())
		return nil, entity.ErrorInternal()
	}

	temp := entity.TemperatureResponse(cep.Localidade, wheather.Current.TempCelcius)

	return temp, nil
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("city")
	uf := r.URL.Query().Get("uf")

	wheather, err := usecase.NewUseCaseWeather(city, uf, viper.GetString("TOKEN"), r.Context())
	otel.GetTextMapPropagator().Inject(r.Context(), propagation.HeaderCarrier(r.Header))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	temp := entity.TemperatureResponse(city, wheather.Current.TempCelcius)

	result, err := json.Marshal(temp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func ZipCodeHandler(w http.ResponseWriter, r *http.Request) {

	zipcode := r.URL.Query().Get("zipcode")

	cep, requestError := RequestWeatherByZipCode(zipcode, r.Context())
	otel.GetTextMapPropagator().Inject(r.Context(), propagation.HeaderCarrier(r.Header))

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

}
