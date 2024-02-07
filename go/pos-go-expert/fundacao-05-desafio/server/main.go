package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

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

type DataUsdToBrl struct {
	UsdToBrl UsdToBrl `json:"USDBRL"`
}

func main() {
	println("Server start")
	http.HandleFunc("/", handler)
	http.HandleFunc("/cotacao", handlerQuote)

	http.ListenAndServe(":8080", nil)
}

func checkIfDone(ctx context.Context, w http.ResponseWriter) bool {
	select {

	case <-ctx.Done():
		msg := "[ QUIT ] Request cancelada, ou limite atingido!\n"
		log.Println(msg)
		w.Write([]byte(msg))
		return true

	default:
		return false
	}
}

func handlerQuote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//for {
	//	time.Sleep(300 * time.Millisecond)

	//time.Sleep(4 * time.Second)
	dt, err := doGetUsdQuote(ctx, w)
	if err != nil {
		println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	time.Sleep(4)

	if checkIfDone(ctx, w) {
		return
	}

	fmt.Sprint(dt)

	cc, err := doSaveData(ctx, w, dt)
	if err != nil {
		println(err.Error())
		//ctx.Done()
		w.Write(cc) //[]byte(cc))
		return
	}
	println(string(cc))

	//log.Println(i)

	if checkIfDone(ctx, w) {
		return
	}

	select {

	case <-ctx.Done():
		// no console print
		log.Println("Request cancelada, ou limite atingido!")
		w.Write([]byte("Request cancelada, ou limite atingido!"))
		return

	default:

		log.Println("44")
		w.Write([]byte("OK"))
	}
}

func doGetUsdQuote(ctx context.Context, w http.ResponseWriter) (DataUsdToBrl, error) {

	//ctxInterno, cancelInterno := context.WithTimeout(context.Background(), 200 * time.Millisecond)
	//defer cancelInterno()

	//-----------------------

	log.Println("Request [start] /cotacao")
	defer log.Println("Request [end] /cotacao")

	client := http.Client{Timeout: 200 * time.Millisecond}
	response, err := client.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")

	if err != nil {
		println("timeout ao acessar a api de cotacoes")
		w.Write([]byte("408, timeout!\n"))
		return DataUsdToBrl{}, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		println("erro ao ler dados da api de cotacoes")
		w.Write([]byte("500, internal error!\n"))
		return DataUsdToBrl{}, err
	}

	var usd DataUsdToBrl
	err = json.Unmarshal(body, &usd)
	if err != nil {
		println("erro ao fazer parser do json da api de cotacoes")
		w.Write([]byte("500, internal error!\n"))
		return DataUsdToBrl{}, err
	}
	//println("\n", usd.UsdToBrl.Bid, "\n")

	//-----------------------
	msg := "Erro ao buscar os dados"
	if checkIfDone(ctx, w) {
		return DataUsdToBrl{}, errors.New(msg)
	}
	//if checkIfDone(ctxInterno, w) {
	//	return nil, errors.New(msg)
	//}

	return usd, nil
	//return []byte(body), nil

}

func doSaveData(ctx context.Context, w http.ResponseWriter, data DataUsdToBrl) ([]byte, error) {
	log.Println("salvar dados...")

	ctxInterno, cancelInterno := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelInterno()

	//time.Sleep(1 * time.Second)
	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
		println("erro ao conectar base de dados")
		w.Write([]byte("500, internal error!\n"))
		return nil, err
	}
	defer db.Close()

	print(db)

	sqlStmt := `
	DROP TABLE IF EXISTS QUOTES;

	CREATE TABLE QUOTES (
		ID 			INTEGER NOT NULL PRIMARY KEY,
		CODE 		TEXT,
		CODEIN  	TEXT,
		NAME   		TEXT,
		HIGH   		REAL,
		LOW        	REAL,
		VARBID     	INTEGER,
		PCTCHANGE  	INTEGER,
		BID        	INTEGER,
		ASK        	REAL,
		TIMESTAMP  	INTEGER,
		CREATEDATE 	TEXT
	);
	`
	_, err = db.ExecContext(ctxInterno, sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, errors.New("Erro ao criar tabela de dados")
	}

	stmt, err := db.Prepare(`
		INSERT INTO QUOTES (CODE, CODEIN, NAME, HIGH, LOW, VARBID, PCTCHANGE, BID, ASK, TIMESTAMP, CREATEDATE)  
		            VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// atencao com a ordem dos campos
	_, err = stmt.ExecContext(ctx, data.UsdToBrl.Code, data.UsdToBrl.CodeIn, data.UsdToBrl.Name, data.UsdToBrl.High, data.UsdToBrl.Low, data.UsdToBrl.VarBid, data.UsdToBrl.PctChange, data.UsdToBrl.Bid, data.UsdToBrl.Ask, data.UsdToBrl.Timestamp, data.UsdToBrl.CreateDate)
	if err != nil {
		return nil, err
	}
	print("Dados salvos com sucesso")

	msg := "Erro ao salvar os dados"
	if checkIfDone(ctx, w) {
		return nil, errors.New(msg)
	}

	if checkIfDone(ctxInterno, w) {
		return nil, errors.New(msg)
	}

	return []byte("d"), nil

}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request [start] /")
	defer log.Println("Request [end] /")

	// print no navegador
	w.Write([]byte("blank page"))
}
