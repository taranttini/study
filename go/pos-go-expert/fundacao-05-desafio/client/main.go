package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const TIMEOUT_PROCESS = 300 * time.Millisecond // 300ms

type UsdToBrl struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_PROCESS)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		print("ERR 1 \n")
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		print("ERR 2 \n")
		panic(err)
	}

	defer res.Body.Close()
	//io.Copy(os.Stdout, res.Body)
	//body, err := io.CopyBuffer(os.Stdout, res.Body, nil)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		print("ERR 3 \n")
		panic(err)
	}
	//data := string(body)
	contentType := res.Header.Get("Content-Type")

	if contentType == "application/json" {
		var usdToBrl UsdToBrl
		err = json.Unmarshal(body, &usdToBrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error to parser response: %v\n", err)
		}

		doSaveFile(usdToBrl)
		return
	}

	print("ERR 4 \n")
	fmt.Print(string(body))

}

func doSaveFile(valor UsdToBrl) {

	//f, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f, err := os.Create("cotacao.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	tmp := template.New("TemplateQuote")
	tmp, _ = tmp.Parse("Dólar: {{ .Bid }}")
	err = tmp.Execute(f, valor)

	if err != nil {
		log.Println(err)
	}
}
