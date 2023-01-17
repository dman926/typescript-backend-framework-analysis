package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type API struct {
	name         string
	port         int32
	mu           sync.Mutex
	totalTime    float64
	requestCount int
}

func sendRequest(method string, url string, api *API, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	start := time.Now()
	client.Do(req)
	elapsedTime := time.Since(start)
	api.mu.Lock()
	api.totalTime += elapsedTime.Seconds()
	api.requestCount++
	api.mu.Unlock()
	client.CloseIdleConnections()
}

func TestEnpoint(api *API, baseUrl string, requestCount int) {
	fmt.Printf("Testing %s on port %d with %d requests\n", api.name, api.port, requestCount)
	url := fmt.Sprintf("%s%d/", baseUrl, api.port)
	wg := new(sync.WaitGroup)
	for i := 0; i < requestCount/2; i++ {
		wg.Add(2)
		go sendRequest("GET", url, api, wg)
		go sendRequest("POST", url, api, wg)
	}
	wg.Wait()
	averageTime := api.totalTime / float64(api.requestCount)
	fmt.Printf("%s average request time: %f seconds/request\n", api.name, averageTime)
}

func shutdownApis(apis []API, baseUrl string) {
	client := &http.Client{}
	for i := 0; i < len(apis); i++ {
		fmt.Printf("Sending shutdown signal to %s\n", apis[i].name)
		url := fmt.Sprintf("%s%d/shutdown", baseUrl, apis[i].port)
		req, _ := http.NewRequest("POST", url, nil)
		client.Do(req)
	}
	client.CloseIdleConnections()
}

func main() {
	requestCount := 25_000

	// Wait an additional 7.5 seconds to be sure APIs have all started
	<-time.After(7500 * time.Millisecond)
	baseUrl := "http://localhost:"
	apis := []API{
		{name: "Express", port: 3000},
		{name: "Fastify", port: 3001},
	}
	for i := 0; i < len(apis); i++ {
		TestEnpoint(&apis[i], baseUrl, requestCount)
	}
	fmt.Println("Analysis Complete. Shutting down APIs")
	shutdownApis(apis, baseUrl)
}
