package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"example.com/types/data"
)

func GetSupportData() []data.SupportData {
	client := http.Client{}

	resp, err := client.Get("http://127.0.0.1:8383/support")
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 500 {
		log.Fatalln("Response failed with status code: ", resp.StatusCode)
		return []data.SupportData{}
	}

	body, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var supportDataSlice []data.SupportData

	err = json.Unmarshal(body, &supportDataSlice)
	if err != nil {
		panic(err)
	}

	fmt.Println("======================")
	fmt.Println("Информация о системе Support:")
	for _, supportData := range supportDataSlice {
		fmt.Println(supportData)
	}

	return supportDataSlice
}

func prepareSupportData(support []data.SupportData) []int {
	var dataSupport []int

	var summaryActiveTickets int
	for _, supportData := range support {
		summaryActiveTickets += supportData.ActiveTickets
	}

	supportTime := 60 / 18 * summaryActiveTickets

	if supportTime < 9 {
		dataSupport = append(dataSupport, 1, supportTime)
	} else if supportTime >= 9 && supportTime < 16 {
		dataSupport = append(dataSupport, 2, supportTime)
	} else {
		dataSupport = append(dataSupport, 3, supportTime)
	}

	return dataSupport
}
