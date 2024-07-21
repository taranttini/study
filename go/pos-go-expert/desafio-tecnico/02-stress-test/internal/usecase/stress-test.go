package usecase

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var statusCodes = make(map[string]int) // criando lista vazia chave string, valor int
var urlCheck = ""
var lineSize = 60

var lock = sync.RWMutex{}

func read(key string) int {
	lock.RLock()
	defer lock.RUnlock()
	return statusCodes[key]
}

func increment(key string, value int) {
	total := read(key)

	lock.Lock()
	defer lock.Unlock()
	statusCodes[key] = total + value
}

func NewStressTest(url string, requestTotal int, concurrency int) {

	//statusCode := GetUrlStatusCode("http://google.com/not-found")
	//statusCodes[statusCode] += 1

	//statusCode = GetUrlStatusCode("http://google.com/not-found")
	//statusCodes[statusCode] += 1

	urlCheck = url

	start := time.Now()

	fmt.Printf("== %s ==\n", CompleteTextWith("=", "=", lineSize))
	text := fmt.Sprintf("Processing %v requests, please wait...", requestTotal)
	fmt.Printf("== %s ==\n", CompleteTextWith(text, " ", lineSize))

	text = fmt.Sprintf("Concurrency %v ", concurrency)
	fmt.Printf("== %s ==\n", CompleteTextWith(text, " ", lineSize))

	fmt.Printf("== %s ==\n", CompleteTextWith("=", "=", lineSize))

	RunConcurrency(concurrency, requestTotal)

	fmt.Printf("== %s ==\n", CompleteTextWith("Status Codes Report", " ", lineSize))
	fmt.Printf("== %s ==\n", CompleteTextWith(" ", " ", lineSize))

	textSpaceSize := 2
	textRightSize := len(fmt.Sprintf("%v", requestTotal))
	textLeftSize := lineSize - textSpaceSize - textRightSize

	total := 0
	for key, value := range statusCodes {

		vKey := CompleteTextWith(key, ".", textLeftSize)
		vValue := CompleteTextWith(fmt.Sprintf("%v", value), ".", textRightSize)
		fmt.Printf("== %s|.%s ==\n", vKey, vValue)
		total += value
	}
	if total < requestTotal {
		vKey := CompleteTextWith("Lost request", ".", textLeftSize)
		vValue := CompleteTextWith(fmt.Sprintf("%v", requestTotal-total), ".", textRightSize)
		fmt.Printf("== %s|.%s ==\n", vKey, vValue)
		total = requestTotal
	}

	fmt.Printf("== %s ==\n", CompleteTextWith(" ", " ", lineSize))
	text = fmt.Sprintf("Total of requests %v ", total)
	fmt.Printf("== %s ==\n", CompleteTextWith(text, " ", lineSize))

	fmt.Printf("== %s ==\n", CompleteTextWith(" ", " ", lineSize))
	fmt.Printf("== %s ==\n", CompleteTextWith("=", "=", lineSize))
	text = fmt.Sprintf("Process finished in %v ms", time.Since(start).Milliseconds())
	fmt.Printf("== %s ==\n", CompleteTextWith(text, " ", lineSize))
	fmt.Printf("== %s ==\n", CompleteTextWith("=", "=", lineSize))
}

func RunConcurrency(concurrency int, totalRequest int) {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(totalRequest)

	sequency := 1
	rest := 0
	if concurrency > 0 {
		sequency = totalRequest / concurrency
		rest = totalRequest % concurrency
	}

	for i := 0; i < sequency; i++ {
		go ConcurrencyTask(&waitGroup, concurrency)
	}
	if rest > 0 {
		go ConcurrencyTask(&waitGroup, rest)
	}
	waitGroup.Wait()
}

func CompleteTextWith(text string, with string, lineSize int) string {
	size := len(text)
	for size < lineSize {
		text += with
		size = len(text)
	}
	return text
}

func ConcurrencyTask(wg *sync.WaitGroup, qty int) {

	for i := 0; i < qty; i++ {

		statusCode := GetUrlStatusCode(urlCheck)
		if len(statusCodes) > 0 {
			increment(statusCode, 1)
		} else {
			increment("500 Internal Error", 1)
		}

		//time.Sleep(1 * time.Second)
		wg.Done()
	}
}

func GetUrlStatusCode(url string) string {

	timeInSecond := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeInSecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "500 Internal Error"
		//err.Error()
		//return "INTERNAL_ERROR"
		//panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "500 Internal Error"
		//err.Error()
		//return "INTERNAL_ERROR"
		//panic(err)
	}
	defer res.Body.Close()

	if checkIfDone(ctx) {
		return "408 Timeout"
	}
	//fmt.Printf("\n%s %v", url, res.StatusCode)
	//fmt.Printf("\n%s %v", url, res.Status)

	return res.Status
}

func checkIfDone(ctx context.Context) bool {
	select {

	case <-ctx.Done():
		return true

	default:
		return false
	}
}
