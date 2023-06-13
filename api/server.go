package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"example.com/types"
	"example.com/types/data"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	*http.Server
	Countries []types.Country
	ResultT   data.ResultT
}

func NewServer(address string, port int) *Server {
	router := chi.NewRouter()

	addr := fmt.Sprintf("%s:%d", address, port)
	s := &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}

	s.Countries = getCountries()

	router.Get("/", s.Home())
	router.Get("/get-data", s.GetData())
	return s
}

func getCountries() []types.Country {
	client := http.Client{}

	resp, err := client.Get("https://datahub.io/core/country-list/r/data.json")
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var countries []types.Country
	err = json.Unmarshal(body, &countries)
	if err != nil {
		panic(err)
	}

	return countries
}

func (s *Server) Home() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		writer.WriteHeader(http.StatusOK)
		ok := "Ok"
		writer.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(writer).Encode(&ok); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (s *Server) GetData() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(writer).Encode(&s.ResultT); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
