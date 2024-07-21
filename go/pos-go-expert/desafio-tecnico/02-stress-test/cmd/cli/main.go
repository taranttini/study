package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/taranttini/study/go/pos-go-expert/desafio-tecnico/02-stress-test/internal/usecase"
)

func main() {

	//fmt.Println(len(os.Args), os.Args)

	if len(os.Args) < 7 {
		ErroArguments()
		return
	}

	urlArg, existsUrl := checkArg("-u")
	requestArg, existsRequest := checkArg("-r")
	concurrencyArg, existsConcurrency := checkArg("-c")

	if !existsConcurrency || !existsRequest || !existsUrl {
		ErroArguments()
		return
	}
	if !isValidUrl(urlArg) {
		println("Url need an url valid")
		ErroArguments()
		return
	}
	totalRequest, err := strconv.Atoi(requestArg)
	if err != nil {
		println("Request need an integer value")
		ErroArguments()
		return
	}
	concurrency, err := strconv.Atoi(concurrencyArg)
	if err != nil {
		println("Concurrency need an integer value")
		ErroArguments()
		return
	}

	// fmt.Printf("%s \n", urlArg)
	// fmt.Printf("%v \n", totalRequest)
	// fmt.Printf("%v \n", concurrency)
	usecase.NewStressTest(urlArg, totalRequest, concurrency)

}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
func checkArg(argParam string) (string, bool) {

	for idx, value := range os.Args {
		if value == argParam {
			//println(idx + 1)
			//println(len(os.Args))
			if idx+1 >= len(os.Args) {
				println("args without a value")
				return "", false
			}
			return os.Args[idx+1], true
		}
	}
	return "", false

}

func ErroArguments() {
	fmt.Println("Args necessary: -u  -r  -c")
	fmt.Println("  -u URL                  // url used in requests")
	fmt.Println("  -r Quantity of Request  // quantity of requests you need")
	fmt.Println("  -c Concurrency          // quantity request in same time")
	fmt.Println("example:")
	fmt.Println("main.go -u https://google.com -r 100 -c 10")
}
