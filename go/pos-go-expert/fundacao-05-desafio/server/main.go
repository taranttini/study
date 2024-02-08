package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const TIMEOUT_HTTP_CLIENT = 200 * time.Millisecond // 200ms
const TIMEOUT_DB_WRITES = 10 * time.Millisecond    // 10ms
const URL_API = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
const ERROR_TIMEOUT_DATABASE_OR_CANCEL = "Timeout to save on database or request canceled!"
const ERROR_TIMEOUT_API_OR_CANCEL = "Timeout get Quote API or request canceled"
const ERROR_TIMEOUT_OR_CANCEL = "Request canceled or reached limit"
const ERROR_500 = "500, internal error!\n"
const ERROR_408 = "408, timeout!\n"
const ERROR_TIMEOUT_API = "Timeout to catch data from quote api"
const ERROR_READ_DATA = "Error to read data from quote api"
const ERROR_JSON_PARSER = "Error to json parser from quote api"

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

type Payload struct {
	UsdToBrl UsdToBrl `json:"USDBRL"`
}

func main() {
	createDatabase()
	println("Server start")
	http.HandleFunc("/", handler)
	http.HandleFunc("/cotacao", handlerQuote)

	http.ListenAndServe(":8080", nil)
}

func checkIfDone(ctx context.Context, w http.ResponseWriter) bool {
	select {

	case <-ctx.Done():
		log.Println(ERROR_TIMEOUT_OR_CANCEL)
		w.Write([]byte(ERROR_TIMEOUT_OR_CANCEL))
		return true

	default:
		return false
	}
}

func handlerQuote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//time.Sleep(1000 * time.Millisecond)
	result, err := doGetUsdQuote(ctx, w)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	if checkIfDone(ctx, w) {
		return
	}

	err = doSaveData(ctx, w, result)
	if err != nil {
		log.Println(err.Error())
		//ctx.Done()
		return
	}
	//log.Println(string(actualBid))

	if checkIfDone(ctx, w) {
		return
	}

	select {

	case <-ctx.Done():
		// no console print
		log.Println(ERROR_TIMEOUT_OR_CANCEL)
		w.Write([]byte(ERROR_TIMEOUT_OR_CANCEL))
		return

	default:
		// response data to client
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result.UsdToBrl)
	}
}

func doGetUsdQuote(ctx context.Context, w http.ResponseWriter) (Payload, error) {

	ctxInterno, cancelInterno := context.WithTimeout(context.Background(), TIMEOUT_HTTP_CLIENT)
	defer cancelInterno()

	//-----------------------

	log.Println("Request [start] /cotacao")
	defer log.Println("Request [end] /cotacao")

	client := http.Client{} //Timeout: TIMEOUT_HTTP_CLIENT}
	response, err := client.Get(URL_API)

	if err != nil {
		log.Println(ERROR_TIMEOUT_API)
		w.Write([]byte(ERROR_408))
		return Payload{}, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Println(ERROR_READ_DATA)
		w.Write([]byte(ERROR_500))
		return Payload{}, err
	}

	var payload Payload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Println(ERROR_JSON_PARSER)
		w.Write([]byte(ERROR_500))
		return Payload{}, err
	}

	//-----------------------
	if checkIfDone(ctx, w) {
		return Payload{}, errors.New(ERROR_TIMEOUT_API_OR_CANCEL)
	}
	if checkIfDone(ctxInterno, w) {
		return Payload{}, errors.New(ERROR_TIMEOUT_API_OR_CANCEL)
	}

	//return []byte(body), nil
	return payload, nil
}

func createDatabase() {
	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStmt := `
	--DROP TABLE IF EXISTS QUOTES;

	CREATE TABLE IF NOT EXISTS QUOTES (
		ID 			INTEGER NOT NULL PRIMARY KEY,
		CODE 		TEXT,
		CODEIN  	TEXT,
		NAME   		TEXT,
		HIGH   		TEXT,
		LOW        	TEXT,
		VARBID     	TEXT,
		PCTCHANGE  	TEXT,
		BID        	TEXT,
		ASK        	TEXT,
		TIMESTAMP  	TEXT,
		CREATEDATE 	TEXT,
		REQUEST_AT  TEXT
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		panic(err)
	}
}

func doSaveData(ctx context.Context, w http.ResponseWriter, payload Payload) error {

	ctxInterno, cancelInterno := context.WithTimeout(context.Background(), TIMEOUT_DB_WRITES)
	defer cancelInterno()

	//time.Sleep(1 * time.Second)
	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		//panic(err)
		log.Println("Error to connect database")
		w.Write([]byte(ERROR_500))
		return err
	}
	defer db.Close()

	//print(db)

	stmt, err := db.Prepare(`
		INSERT INTO QUOTES (CODE, CODEIN, NAME, HIGH, LOW, VARBID, PCTCHANGE, BID, ASK, TIMESTAMP, CREATEDATE, REQUEST_AT)  
		            VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, datetime('now','localtime'))
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// atencao com a ordem dos campos
	_, err = stmt.ExecContext(ctx, payload.UsdToBrl.Code, payload.UsdToBrl.CodeIn, payload.UsdToBrl.Name, payload.UsdToBrl.High, payload.UsdToBrl.Low, payload.UsdToBrl.VarBid, payload.UsdToBrl.PctChange, payload.UsdToBrl.Bid, payload.UsdToBrl.Ask, payload.UsdToBrl.Timestamp, payload.UsdToBrl.CreateDate)
	if err != nil {
		return err
	}
	log.Print("Successful data saved")

	if checkIfDone(ctx, w) {
		return errors.New(ERROR_TIMEOUT_DATABASE_OR_CANCEL)
	}

	if checkIfDone(ctxInterno, w) {
		return errors.New(ERROR_TIMEOUT_DATABASE_OR_CANCEL)
	}

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request [start] /")
	defer log.Println("Request [end] /")

	// print no navegador
	w.Write([]byte("blank page"))
}
