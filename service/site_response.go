package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"maxsol/model"
	"net/http"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var ResultResponse string

func ResponseSize(url string) {
	defer wg.Done()

	fmt.Println("Step1: ", url)
	ResultResponse += "Step1: " + url
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	ResultResponse += "Step2: " + url
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	ResultResponse += "Step3: " + url
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
	ResultResponse += "Step4: " + strconv.Itoa(len(body))
}

func RunResponse(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	wg.Add(3)
	//fmt.Println("Start Goroutines")

	go ResponseSize("https://www.golangprograms.com")
	go ResponseSize("https://stackoverflow.com")
	go ResponseSize("https://coderwall.com")

	// Wait for the goroutines to finish.
	wg.Wait()
	//fmt.Println("Terminating Program")

	response.Status = "Berhasil"
	response.Message = "Goroutines"
	response.Data = ResultResponse

	fmt.Println("Endpoint Hit: respond goroutines")
	json.NewEncoder(w).Encode(response)
}
