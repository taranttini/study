package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	timeInMilliseconds := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeInMilliseconds)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		print("\nxxx\n")
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		print("\n111\n")
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
