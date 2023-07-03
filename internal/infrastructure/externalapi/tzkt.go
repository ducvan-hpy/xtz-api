package externalapi

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	model "github.com/ducvan-hpy/xtz-api/internal/domain/model"
)

const getDelegationsPath = "/v1/operations/delegations"

type TzktSDK struct {
	server string
}

// Delegation defines model for Delegation.
type Delegation struct {
	ID        int       `json:"id,omitempty"`
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

func (t *TzktSDK) GetDelegations(ctx context.Context, limit, lastID int) ([]model.Delegation, error) {
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("offset.cr", strconv.Itoa(lastID))

	url := t.server + getDelegationsPath + "?" + params.Encode()

	log.Printf("Calling %s", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("fail to get delegations: %v", err)
		return []model.Delegation{}, err
	}

	var delegations []Delegation
	if err := json.NewDecoder(resp.Body).Decode(&delegations); err != nil {
		log.Printf("fail to decode delegations: %v", err)
		return []model.Delegation{}, err
	}

	domainDelegations := make([]model.Delegation, 0, len(delegations))
	for _, d := range delegations {
		domainDelegations = append(domainDelegations, d.ToDomain())
	}

	return domainDelegations, nil
}

func (d Delegation) ToDomain() model.Delegation {
	return model.Delegation{
		ID:        d.ID,
		Amount:    d.Amount,
		Block:     d.Block,
		Delegator: d.Sender.Address,
		Timestamp: d.Timestamp,
	}
}
