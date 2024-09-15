package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// load env vars cfg
func init() {
	viper.AutomaticEnv()
}

type ZipCodeRequest struct {
	ZiCode string `json:"cep"`
}

type TemperatureResponse struct {
	City       string  `json:"city"`
	Celcius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func isInvalideCep(cep string) bool {

	if len(cep) != 8 {
		return true
	}

	_, err := strconv.Atoi(cep)

	return err != nil
}

type CepRequest struct {
	Cep string `json:"cep"`
}

func getCep(w http.ResponseWriter, r *http.Request) string {
	defer r.Body.Close()

	dataRead, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Cep: Erro ao ler resposta: %s\n", err.Error())))
		return ""
	}

	var cep CepRequest
	err = json.Unmarshal(dataRead, &cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Cep: Erro ao fazer o parser resposta: %s\n", err.Error())))
		return ""
	}

	return cep.Cep
}

func ZipCodeHandle(w http.ResponseWriter, r *http.Request) {

	_, span := otel.Tracer(viper.GetString("SPAN_TRACE_NAME")).Start(r.Context(), "start-Zip-Code-Handle")
	defer span.End()

	cep := getCep(w, r)
	if len(cep) == 0 {
		w.WriteHeader(422)
		w.Write([]byte("invalid zipcode"))
		return
	}

	if isInvalideCep(cep) {
		w.WriteHeader(422)
		w.Write([]byte("invalid zipcode"))
		return
	}

	url := fmt.Sprintf("%s/zipcode?zipcode=%s", viper.GetString("WEATHER_ENDPOINT"), cep)

	response, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Cep: Erro ao fazer requisicao: %s\n", err.Error())))
		return
	}

	defer response.Body.Close()
	otel.GetTextMapPropagator().Inject(r.Context(), propagation.HeaderCarrier(r.Header))

	dataRead, err := io.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Cep: Erro ao ler resposta: %s\n", err.Error())))
		return
	}

	if response.StatusCode != 200 {
		print("xxxx")
		w.WriteHeader(response.StatusCode)
		w.Write(dataRead)
		return
	}

	var temperature TemperatureResponse
	err = json.Unmarshal(dataRead, &temperature)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Cep: Erro ao fazer o parser resposta: %s\n", err.Error())))
		return
	}

	jsonData, err := json.Marshal(temperature)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, span = otel.Tracer(viper.GetString("SPAN_TRACE_NAME")).Start(r.Context(), "end-Zip-Code-Handle")
	defer span.End()

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
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

	fmt.Println(viper.GetString("SPAN_TRACE_NAME"))
	initTracer(viper.GetString("SPAN_TRACE_NAME"))

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Post("/", ZipCodeHandle)

	port := viper.GetString("HTTP_PORT")
	fmt.Printf("Listen on %s\n", port)
	http.ListenAndServe(port, r)
}
