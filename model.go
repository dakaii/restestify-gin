package restestify

import (
	"time"
	"log"
	"bytes"
	"net/http"
	"encoding/json"
)

type Request struct {
	Path     string    `json:"path"`
	Query    string    `json:"query"`
	Method   string    `json:"method"`
	Status   int       `json:"status"`
	Latency  int64     `json:"latency"`
	ClientIp string    `json:"client_ip"`
	Time     time.Time `json:"time"`
}

type RequestCollection struct {
	Requests []*Request `json:"requests"`
}

func (r *Request) Upload() error {
	// TODO bundle requests
	reqs := RequestCollection{}
	reqs.Requests = append(reqs.Requests, r)

	body, err := json.Marshal(&reqs)
	if err != nil {
		log.Printf("ERROR: while marshaling json request for restestify: %v", err.Error())
		return err
	}

	req, err := http.NewRequest("POST", "https://api.restestify.com/v1/requests", bytes.NewBuffer(body))
	req.Header.Set("X-Api-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{ Timeout: time.Duration(time.Second * 10) }
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: while uploading request logs to restestify: %v", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("error: restestify api responded with non 200 code: %d", resp.StatusCode)
	}

	return nil
}
