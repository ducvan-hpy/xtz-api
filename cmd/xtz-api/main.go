package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ducvan-hpy/xtz-api/cmd/xtz-api/api"
	"github.com/ducvan-hpy/xtz-api/cmd/xtz-api/poller"
	"github.com/ducvan-hpy/xtz-api/internal/domain/repository"
	"github.com/ducvan-hpy/xtz-api/internal/infrastructure/externalapi"
	"github.com/ducvan-hpy/xtz-api/internal/infrastructure/persistence"
)

const (
	host         = "localhost"
	port         = 8090
	pollInterval = 10 * time.Second
)

func main() {
	repo := &repository.Repository{
		Delegation: persistence.NewInMemoryStorage(),
	}

	restAPI := api.New(repo)
	engine := api.NewGinRouter(restAPI)

	// Run poller in goroutine.
	tzktSDK := externalapi.NewTzktSDK("https://api.tzkt.io")
	delegationPoller := poller.New(tzktSDK.GetDelegations, pollInterval, repo)
	go delegationPoller.Start(context.Background())

	serverURL := fmt.Sprintf("%s:%d", host, port)
	err := http.ListenAndServe(serverURL, engine)
	if err != nil {
		log.Fatal(err)
	}
}
