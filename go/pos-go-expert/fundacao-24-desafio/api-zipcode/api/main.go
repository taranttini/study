package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"context"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
	"github.com/taranttini/study/go/post/fundacao-24-desafio/api-zipcode/usecase"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	viper.AutomaticEnv()
}

//func main() {
//	http.HandleFunc("/", ZipCodeHandler)
//	fmt.Printf("Listen on %s\n", viper.GetString("HTTP_PORT"))
//	http.ListenAndServe(viper.GetString("HTTP_PORT"), nil)
//}

func (h *Webserver) ZipCodeHandler(w http.ResponseWriter, r *http.Request) {

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

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	_, spanInicial := h.TemplateData.OTELTracer.Start(ctx, "SPAN_ZIPCODE")
	zipcode, err := usecase.NewUseCaseZipcode(queryZipcode)
	spanInicial.End()

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

	_, spanInicial1 := h.TemplateData.OTELTracer.Start(ctx, "SPAN_WEATHER")
	temperature, err := ProcessWeather(zipcode.Localidade, zipcode.Uf, w, r)
	spanInicial1.End()

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

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))

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

	urlEndpoint := fmt.Sprintf("%s/?city=%s&uf=%s",
		viper.GetString("WEATHER_ENDPOINT"), url.QueryEscape(city), uf)

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

// #############
// #############
//
//
//
// #############
// #############

func initProvider(serviceName, collectorURL string) (func(context.Context) error, error) {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}
	/*
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
			conn, err := grpc.DialContext(ctx, collectorURL,
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				//grpc.WithBlock(),
			)
			if err != nil {
				return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
			}

			traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
			if err != nil {
				return nil, fmt.Errorf("failed to create trace exporter: %w", err)
			}
	*/

	fmt.Print(collectorURL)

	traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpointURL(collectorURL))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider.Shutdown, nil
}

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := initProvider(viper.GetString("OTEL_SERVICE_NAME"), viper.GetString("OTEL_EXPORTER_OTLP_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	tracer := otel.Tracer("microservice-tracer")

	templateData := &TemplateData{
		Title:              viper.GetString("TITLE"),
		BackgroundColor:    viper.GetString("BACKGROUND_COLOR"),
		ResponseTime:       time.Duration(viper.GetInt("RESPONSE_TIME")),
		ExternalCallURL:    viper.GetString("EXTERNAL_CALL_URL"),
		ExternalCallMethod: viper.GetString("EXTERNAL_CALL_METHOD"),
		RequestNameOTEL:    viper.GetString("REQUEST_NAME_OTEL"),
		OTELTracer:         tracer,
	}
	server := NewServer(templateData)
	router := server.CreateServer()
	router.HandleFunc("/", server.ZipCodeHandler)

	go func() {
		log.Println("Starting server on port", viper.GetString("HTTP_PORT"))
		if err := http.ListenAndServe(viper.GetString("HTTP_PORT"), router); err != nil {
			log.Fatal(err)
		}
	}()

	select {
	case <-sigCh:
		log.Println("Shutting down gracefully, CTRL+C pressed...")
	case <-ctx.Done():
		log.Println("Shutting down due to other reason...")
	}

	// Create a timeout context for the graceful shutdown
	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
}

type Webserver struct {
	TemplateData *TemplateData
}

// NewServer creates a new server instance
func NewServer(templateData *TemplateData) *Webserver {
	return &Webserver{
		TemplateData: templateData,
	}
}

// createServer creates a new server instance with go chi router
func (we *Webserver) CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	// promhttp
	router.Handle("/metrics", promhttp.Handler())
	//router.Get("/", we.HandleRequest)
	return router
}

type TemplateData struct {
	Title              string
	BackgroundColor    string
	ResponseTime       time.Duration
	ExternalCallMethod string
	ExternalCallURL    string
	Content            string
	RequestNameOTEL    string
	OTELTracer         trace.Tracer
}

func (h *Webserver) HandleRequest(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanInicial := h.TemplateData.OTELTracer.Start(ctx, "SPAN_INICIAL"+h.TemplateData.RequestNameOTEL)
	time.Sleep(time.Second)
	spanInicial.End()

	println("xxxx")
	ctx, span := h.TemplateData.OTELTracer.Start(ctx, "Chama externa"+h.TemplateData.RequestNameOTEL)
	defer span.End()

	time.Sleep(time.Millisecond * h.TemplateData.ResponseTime)

	if h.TemplateData.ExternalCallURL != "" {
		var req *http.Request
		var err error
		if h.TemplateData.ExternalCallMethod == "GET" {
			req, err = http.NewRequestWithContext(ctx, "GET", h.TemplateData.ExternalCallURL, nil)
		} else if h.TemplateData.ExternalCallMethod == "POST" {
			req, err = http.NewRequestWithContext(ctx, "POST", h.TemplateData.ExternalCallURL, nil)
		} else {
			http.Error(w, "Invalid ExternalCallMethod", http.StatusInternalServerError)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.TemplateData.Content = string(bodyBytes)
	}

}
