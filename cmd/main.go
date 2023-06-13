package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/api"
	"example.com/src"
)

func main() {
	var port = flag.Int("port", 9000, "Set port")

	var address string
	flag.StringVar(&address, "addr", "127.0.0.1", "Set address")

	log.Print("Подготовка к запуску сервера")
	server := api.NewServer(address, *port)

	smsData := src.ReadSMSDataFromFile("simulator/sms.data", server.Countries)
	mmsData := src.GetMMSData(server.Countries)
	voiceData := src.ReadVoiceDataFromFile("simulator/voice.data", server.Countries)
	emailData := src.ReadEmailDataFromFile("simulator/email.data", server.Countries)
	billingData := src.ReadBillingDataFromFile("simulator/billing.data")
	supportData := src.GetSupportData()
	incidentData := src.GetIncidentData()

	server.ResultT = src.GetResultData(smsData, mmsData, voiceData, emailData, billingData, supportData, incidentData, server.Countries)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("server is listen on port", server.Addr)

	<-done

	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
