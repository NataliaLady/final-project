package src

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"example.com/internal"
	"example.com/types"
	"example.com/types/data"
	"golang.org/x/exp/slices"
)

var voiceProviders = []string{"TransparentCalls", "E-Voice", "JustPhone"}

func ReadVoiceDataFromFile(fileName string, countries []types.Country) []data.VoiceData {
	var voiceDataSlice []data.VoiceData

	file, err := os.Open(fileName)
	if err != nil {
		return []data.VoiceData{}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	result, err := io.ReadAll(file)
	if err != nil {
		return []data.VoiceData{}
	}

	lines := strings.Split(strings.Trim(string(result), " "), "\n")

	fmt.Println("======================")
	fmt.Println("Информация о системе Voice")
	for _, line := range lines {
		split := strings.Split(line, ";")

		if len(split) != 8 {
			continue
		}

		countryIndex := slices.IndexFunc(countries, func(country types.Country) bool {
			return country.Code == split[0]
		})

		if countryIndex == -1 {
			continue
		}

		if !slices.Contains(voiceProviders, split[3]) {
			continue
		}

		connectionStability, err := strconv.ParseFloat(split[4], 32)
		if err != nil {
			return []data.VoiceData{}
		}

		voiceData := data.VoiceData{
			Country:             split[0],
			Bandwidth:           internal.StrToInt(split[1]),
			AverageResponse:     internal.StrToInt(split[2]),
			Provider:            split[3],
			ConnectionStability: float32(connectionStability),
			TTFB:                internal.StrToInt(split[5]),
			VoicePurity:         internal.StrToInt(split[6]),
			MedianCallDuration:  internal.StrToInt(split[7]),
		}

		voiceDataSlice = append(voiceDataSlice, voiceData)
		fmt.Println(voiceData)
	}

	return voiceDataSlice
}
