package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"

	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/taranttini/study/go/pos-go-expert/fundacao-10-desafio/configs"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-10-desafio/internal/entity"
)

const ERROR_TIMEOUT_API_OR_CANCEL = "Timeout get API or request canceled"
const ERROR_TIMEOUT_OR_CANCEL = "Request canceled or reached limit"
const ERROR_BRASILAPI_JSON_PARSER = "Error to json parser from brasilapi"
const ERROR_VIACEP_JSON_PARSER = "Error to json parser from viacep"

func main() {
	fmt.Print("Consultar CEP - ViaCEP X BrasilApi\n\n")
	config, _ := configs.LoadConfig(".")
	// println(config.HttpTimeout)
	// println(config.UrlBrasilApi)
	// println(config.UrlViaCep)
	// println(config.WebServerPort)

	var cep = "02765000"
	var arg = ""

	if len(os.Args) >= 2 {
		arg = os.Args[1]
	}

	if len(arg) != 8 {
		fmt.Printf("Não foi fornecido um cep como parametro, usando cep padrão [ %s ]\n", cep)
		arg = cep
	}

	_, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Não foi fornecido um cep no formato 00000000, usando cep ppadrão [ %s ]\n", cep)
		arg = cep
	}
	cep = arg

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)

	var payload = Payload{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.HttpTimeout)*time.Millisecond)
	defer cancel()

	go callApi(ctx, config.UrlBrasilApi, cep, &waitGroup, &payload)
	go callApi(ctx, config.UrlViaCep, cep, &waitGroup, &payload)

	time.Sleep(time.Duration(config.HttpTimeout) * time.Millisecond * 2)
}

type Payload struct {
	BrasilApi entity.BrasilApi `json:"brasilApi"`
	ViaCep    entity.ViaCep    `json:"viaCep"`
}

func callApi(ctx context.Context, url string, cep string, wg *sync.WaitGroup, p *Payload) {
	err := doGetDataFromUrl(ctx, url, cep, wg, p)
	if err != nil {
		fmt.Println(err.Error())
	}
	wg.Done()
}

func payloadBrasilApi(body []byte) (entity.BrasilApi, error) {
	var payload entity.BrasilApi

	err := json.Unmarshal(body, &payload)
	if err != nil {
		//log.Println(ERROR_BRASILAPI_JSON_PARSER)
		return entity.BrasilApi{}, err
	}
	return payload, nil
}

func payloadViaCep(body []byte) (entity.ViaCep, error) {
	var payload entity.ViaCep

	err := json.Unmarshal(body, &payload)
	if err != nil {
		//log.Println(ERROR_VIACEP_JSON_PARSER)
		return entity.ViaCep{}, err
	}
	return payload, nil
}

func payloadAny(body []byte, url string, wg *sync.WaitGroup, p *Payload) error {
	config, _ := configs.LoadConfig(".")

	switch url {
	case config.UrlBrasilApi:
		//if url == config.UrlBrasilApi {

		data, err := payloadBrasilApi(body)
		if err != nil {
			return err
		}
		// time.Sleep(900 * time.Millisecond) // forcar atraso
		if len(p.ViaCep.Cep) == 0 {
			p.BrasilApi = data

			fmt.Printf("[ A api da BrasilApi foi mais rápida ]\n")
			fmt.Printf("%s, %s - %s CEP %s\n", data.Street, data.City, data.State, data.Cep)
		}

	case config.UrlViaCep:
		//if url == config.UrlViaCep {
		data, err := payloadViaCep(body)
		if err != nil {
			return err
		}
		// time.Sleep(900 * time.Millisecond) //forcar atraso
		if len(p.BrasilApi.Cep) == 0 {
			p.ViaCep = data

			fmt.Printf("[ A api da ViaCep foi mais rápida ]\n")
			fmt.Printf("%s, %s - %s CEP %s\n", data.Logradouro, data.Localidade, data.Uf, data.Cep)
		}

	default:
		panic("url inválida")
	}
	wg.Done()
	return nil
}

func makeUrl(url string, cep string) string {
	config, _ := configs.LoadConfig(".")

	switch url {
	case config.UrlBrasilApi:
		return fmt.Sprintf("%s/%s", url, cep)

	case config.UrlViaCep:
		return fmt.Sprintf("%s/%s/json", url, cep)

	default:
		panic("url inválida")
	}
}

func doGetDataFromUrl(ctx context.Context, url string, cep string, wg *sync.WaitGroup, p *Payload) error {

	config, _ := configs.LoadConfig(".")

	ctxInterno, cancelInterno := context.WithTimeout(context.Background(), time.Duration(config.HttpTimeout)*time.Millisecond)
	defer cancelInterno()

	//log.Printf("Request [start] %s\n", makeUrl(url, cep))
	//defer log.Printf("Request [end] %s\n", makeUrl(url, cep))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.HttpTimeout)*time.Millisecond)
	defer cancel()
	//time.Sleep(1000 * time.Millisecond) //forcar timeout
	// criar request contexto
	req, err := http.NewRequestWithContext(ctx, "GET", makeUrl(url, cep), nil)
	if err != nil {
		//print("ERR 1 \n")
		return err
	}
	// executa request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		//print("ERR 2 \n")
		return err
	}
	// ler os dados
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		//print("ERR 3 \n")
		return err
	}

	// processa o dado
	err = payloadAny(body, url, wg, p)
	if err != nil {
		return err
	}

	if checkIfDone(ctx) {
		return errors.New(ERROR_TIMEOUT_API_OR_CANCEL)
	}

	if checkIfDone(ctxInterno) {
		return errors.New(ERROR_TIMEOUT_API_OR_CANCEL)
	}

	return nil
}

func checkIfDone(ctx context.Context) bool {
	select {

	case <-ctx.Done():
		log.Println(ERROR_TIMEOUT_OR_CANCEL)
		return true

	default:
		return false
	}
}
