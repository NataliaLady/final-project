package src

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"

	"example.com/types"
	"example.com/types/data"
	"golang.org/x/exp/slices"
)

var (
	smsProviders = []string{"Topolo", "Rond", "Kildy"}
	dataSMSArray [][]data.SMSData
)

func ReadSMSDataFromFile(fileName string, countries []types.Country) []data.SMSData {
	var smsDataSlice []data.SMSData

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil
	}

	fmt.Println("======================")
	fmt.Println("Информация о системе SMS")
	for _, line := range lines {
		if len(line) != 4 {
			continue
		}

		countryIndex := slices.IndexFunc(countries, func(country types.Country) bool {
			return country.Code == line[0]
		})

		if countryIndex == -1 {
			continue
		}

		if !slices.Contains(smsProviders, line[3]) {
			continue
		}

		smsData := data.SMSData{
			Country:      line[0],
			Bandwidth:    line[1],
			ResponseTime: line[2],
			Provider:     line[3],
		}

		fmt.Println(smsData)
		smsDataSlice = append(smsDataSlice, smsData)
	}

	return smsDataSlice
}

func prepareSMSData(sms []data.SMSData, countries []types.Country) [][]data.SMSData {
	copySMS := make([]data.SMSData, len(sms))

	for i := 0; i < len(sms); i++ {
		countryIndex := slices.IndexFunc(countries, func(country types.Country) bool {
			return country.Code == sms[i].Country
		})

		sms[i].Country = countries[countryIndex].Name
	}

	sort.Slice(sms, func(i, j int) bool {
		return sms[i].Provider < sms[j].Provider
	})

	dataSMSArray = append(dataSMSArray, sms)
	copy(copySMS, sms)

	sort.Slice(copySMS, func(i, j int) bool {
		return copySMS[i].Country < copySMS[j].Country
	})

	dataSMSArray = append(dataSMSArray, copySMS)
	return dataSMSArray
}
