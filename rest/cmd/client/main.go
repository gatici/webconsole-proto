package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	openapi "webconsole-proto/rest/api"
)

func main() {

	// Set up the HTTP client
	client := &http.Client{}

	// Send request every 5 seconds
	for {
		resp, err := client.Get("http://localhost:8080/config/access-and-mobility")
		if err != nil {
			log.Println("Error sending request:", err)
			continue
		}

		// Handle the response
		if resp.StatusCode == http.StatusOK {
			fmt.Println("Request successful:", resp.Status)
			var response openapi.AccessAndMobilityConfigResponse
			err := json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				log.Println("Error decoding response:", err)
				resp.Body.Close()
				continue
			}

			for _, cfg := range response.Configs {
				log.Printf("PLMN: MCC %s MNC %s, SNSSAI: SST %s SD %s, TACs: %v",
					cfg.PlmnId.Mcc, cfg.PlmnId.Mnc,
					cfg.Snssai.Sst, cfg.Snssai.Sd,
					cfg.Tacs,
				)
			}
		} else {
			fmt.Println("Request failed with status:", resp.Status)
		}

		resp.Body.Close()

		// Wait for 5 seconds before sending the next request
		time.Sleep(5 * time.Second)
	}
}
