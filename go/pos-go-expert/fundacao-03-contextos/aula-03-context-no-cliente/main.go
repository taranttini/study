package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

// essa tarefa depende da tarefa aula-02
// se eu passar 3 segundos gero erro,
// se eu passar 6 ou mais segundos, eu obtenho sucesso

func main() {

	//timeInSecond := 3 * time.Second
	timeInSecond := 6 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeInSecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
