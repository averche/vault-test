package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hashicorp/vault/vault"
	"github.com/hashicorp/vault/api"
)

func main() {
	// reference vault
	cluster := vault.NewTestCluster(nil, nil, &vault.TestClusterOptions{
		HandlerFunc: func(props *vault.HandlerProperties) http.Handler {
			fmt.Println("test")
			return nil
		},
	})

	cluster.Start()
	defer cluster.Cleanup()

	// reference api
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf(err)
	}

	raw, err := client.Logical().ReadRawWithContext(context.Background(), "secret/data/foo")
	if err != nil {
		log.Fatalf(err)
	}
	defer raw.Body.Close()

	b, err := io.ReadAll(raw.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("raw", string(b))
}
