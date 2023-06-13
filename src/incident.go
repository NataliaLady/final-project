package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"

	"example.com/types/data"
)

func GetIncidentData() []data.IncidentData {
	client := http.Client{}

	resp, err := client.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 500 {
		log.Fatalln("Response failed with status code: ", resp.StatusCode)
		return []data.IncidentData{}
	}

	body, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var incidentDataSlice []data.IncidentData

	err = json.Unmarshal(body, &incidentDataSlice)
	if err != nil {
		return []data.IncidentData{}
	}

	fmt.Println("======================")
	fmt.Println("Информация о системе Incident:")
	for _, incidentData := range incidentDataSlice {
		fmt.Println(incidentData)
	}

	return incidentDataSlice
}

func prepareIncidentData(incidents []data.IncidentData) []data.IncidentData {
	sort.Slice(incidents, func(i, j int) bool {
		return incidents[i].Status < incidents[j].Status
	})

	return incidents
}
