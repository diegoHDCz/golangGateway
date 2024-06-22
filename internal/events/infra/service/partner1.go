package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner1 struct {
	BaseURL string
}

type Partner1ReservationRequest struct {
	Spots      []string `json:"spots"`
	TicketKind string   `json:"ticket_kind"`
	Email      string   `json:email`
}

type Partner1ReservationResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketKind string `json:"ticket_kind"`
	Status     string `json:"status"`
	EventId    string `json:"event_id"`
}

func (p *Partner1) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	partnerRequest := Partner1ReservationRequest{
		Spots:      req.Spots,
		TicketKind: req.TicketKind,
		Email:      req.TicketKind,
	}
	body, err := json.Marshal(partnerRequest)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprint("%s/events/%s/reserve", p.BaseURL, req.EventID)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected statuscode: %d", httpResp.StatusCode)
	}

	var partnerResp []Partner1ReservationResponse
	if err := json.NewDecoder(httpReq.Body).Decode(&partnerResp); err != nil {
		return nil, err
	}

	responses := make([]ReservationResponse, len(partnerResp))
	for i, r := range partnerResp {
		responses[i] = ReservationResponse{
			ID:     r.ID,
			Spot:   r.Spot,
			Status: r.Status,
		}
	}
	return responses, nil
}
