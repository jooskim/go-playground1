package sampleModule

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func Hello() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Hello")
}

func FetchWeb(url string, wg *sync.WaitGroup, ch chan []byte) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer func() {
		resp.Body.Close()
		close(ch)
		wg.Done()
	}()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	ch <- body
}

type Boohoo struct {
	Id int
	Name string
	Age int
}

func Greeting(name string) string {
	return "Hello " + name
}

func TestRun() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Hello()
	}()

	c2 := make(chan []byte)
	wg.Add(1)
	go FetchWeb("https://google.com", wg, c2)

	for results := range(c2) {
		fmt.Println(string(results))
	}

	wg.Wait()
}