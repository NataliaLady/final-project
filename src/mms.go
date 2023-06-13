package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"

	"example.com/types"
	"example.com/types/data"
	"golang.org/x/exp/slices"
)

var (
	mmsProviders = []string{"Topolo", "Rond", "Kildy"}
	dataMMSArray [][]data.MMSData
)

func GetMMSData(countries []types.Country) []data.MMSData {
	client := http.Client{}

	resp, err := client.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return []data.MMSData{}
	}

	if resp.StatusCode == 500 {
		log.Fatalln("Response failed with status code: ", resp.StatusCode)
		return []data.MMSData{}
	}

	body, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var mmsDataSlice []data.MMSData

	err = json.Unmarshal(body, &mmsDataSlice)
	if err != nil {
		panic(err)
	}

	fmt.Println("======================")
	fmt.Println("Информация о системе MMS")
	for i, mmsData := range mmsDataSlice {
		countryIndex := slices.IndexFunc(countries, func(country types.Country) bool {
			return country.Code == mmsData.Country
		})

		if countryIndex == -1 {
			mmsDataSlice = append(mmsDataSlice[:i], mmsDataSlice[i:]...)
			continue
		}

		if !slices.Contains(mmsProviders, mmsData.Provider) {
			mmsDataSlice = append(mmsDataSlice[:i], mmsDataSlice[i:]...)
			continue
		}

		fmt.Println(mmsData)
	}

	return mmsDataSlice
}

func prepareMMSData(mms []data.MMSData, countries []types.Country) [][]data.MMSData {
	copyMMS := make([]data.MMSData, len(mms))

	for i := 0; i < len(mms); i++ {
		countryIndex := slices.IndexFunc(countries, func(country types.Country) bool {
			return country.Code == mms[i].Country
		})

		mms[i].Country = countries[countryIndex].Name
	}

	sort.Slice(mms, func(i, j int) bool {
		return mms[i].Provider < mms[j].Provider
	})

	dataMMSArray = append(dataMMSArray, mms)
	copy(copyMMS, mms)

	sort.Slice(copyMMS, func(i, j int) bool {
		return copyMMS[i].Country < copyMMS[j].Country
	})

	dataMMSArray = append(dataMMSArray, mms)
	return dataMMSArray
}
