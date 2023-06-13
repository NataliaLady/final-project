package src

import (
	"io"
	"os"
	"sort"
	"strings"

	"example.com/internal"
	"example.com/types"
	"example.com/types/data"
	"golang.org/x/exp/slices"
)

var emailProviders = []string{
	"Gmail", "Yahoo", "Hotmail", "MSN",
	"Orange", "Comcast", "AOL", "Live",
	"RediffMail", "GMX", "Proton Mail", "Yandex",
	"Mail.ru"}

func ReadEmailDataFromFile(fileName string, countries []types.Country) []data.EmailData {
	var emailDataSlice []data.EmailData

	file, err := os.Open(fileName)
	if err != nil {
		return []data.EmailData{}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	result, err := io.ReadAll(file)
	if err != nil {
		return []data.EmailData{}
	}

	lines := strings.Split(strings.Trim(string(result), " "), "\n")

	for _, line := range lines {
		split := strings.Split(line, ";")

		if len(split) != 3 {
			continue
		}

		countryIndex := slices.IndexFunc(countries, func(country types.Country) bool {
			return country.Code == split[0]
		})

		if countryIndex == -1 {
			continue
		}

		if !slices.Contains(emailProviders, split[1]) {
			continue
		}

		emailData := data.EmailData{
			Country:      split[0],
			Provider:     split[1],
			DeliveryTime: internal.StrToInt(split[2]),
		}

		emailDataSlice = append(emailDataSlice, emailData)
	}

	return emailDataSlice
}

func prepareEmailData(email []data.EmailData, countries []types.Country) map[string][][]data.EmailData {
	result := make(map[string][][]data.EmailData)

	var emailCountries []types.Country
	for i := 0; i < len(email); i++ {
		countryIndex := slices.IndexFunc(countries, func(country types.Country) bool {
			return country.Code == email[i].Country
		})

		emailCountries = append(emailCountries, countries[countryIndex])
	}

	for _, country := range emailCountries {
		var emailDataSlices [][]data.EmailData

		emailDataSlices = append(emailDataSlices, getMin(email, country.Code, 3))
		emailDataSlices = append(emailDataSlices, getMax(email, country.Code, 3))

		result[country.Code] = emailDataSlices
	}

	return result
}

func getMax(email []data.EmailData, code string, count int) []data.EmailData {
	var result []data.EmailData

	for _, emailData := range email {
		if emailData.Country == code {
			result = append(result, emailData)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].DeliveryTime > result[j].DeliveryTime
	})

	if len(result) < count {
		return result
	}

	return result[:count]
}

func getMin(email []data.EmailData, code string, count int) []data.EmailData {
	var result []data.EmailData

	for _, emailData := range email {
		if emailData.Country == code {
			result = append(result, emailData)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].DeliveryTime < result[j].DeliveryTime
	})

	if len(result) < count {
		return result
	}

	return result[:count]
}
