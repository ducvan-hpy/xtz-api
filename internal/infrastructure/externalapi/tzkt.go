package externalapi

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	models "github.com/ducvan-hpy/xtz-api/internal/domain/models"
)

const getDelegationsPath = "/v1/operations/delegations"

type TzktSDK struct {
	server string
}

// Delegation defines model for Delegation.
type Delegation struct {
	Id        int       `json:"id,omitempty"`
	Amount    int       `json:"amount,omitempty"`
	Block     string    `json:"block,omitempty"`
	Sender    Sender    `json:"sender,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type Sender struct {
	Address string `json:"address,omitempty"`
}

func NewTzktSDK(server string) *TzktSDK {
	return &TzktSDK{
		server: server,
	}
}

func (t *TzktSDK) GetDelegations(ctx context.Context) ([]models.Delegation, error) {
	var delegations []Delegation

	resp, err := http.Get(t.server + getDelegationsPath)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("fail to get delegations: %v", err)
		return []models.Delegation{}, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&delegations); err != nil {
		log.Printf("fail to decode delegations: %v", err)
		return []models.Delegation{}, err
	}

	domainDelegations := make([]models.Delegation, 0, len(delegations))
	for _, d := range delegations {
		domainDelegations = append(domainDelegations, d.ToDomain())
	}

	return domainDelegations, nil
}

func (d Delegation) ToDomain() models.Delegation {
	return models.Delegation{
		Amount:    d.Amount,
		Block:     d.Block,
		Delegator: d.Sender.Address,
		Timestamp: d.Timestamp,
	}
}
