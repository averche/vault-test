package main

import (
	"context"
	"io"
	"log"

	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/vault"
)

func main() {
	// reference vault
	cluster := vault.NewTestCluster(nil, nil, &vault.TestClusterOptions{})
	cluster.Start()
	defer cluster.Cleanup()

	// reference api
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	raw, err := client.Logical().ReadRawWithContext(context.Background(), "secret/data/foo")
	if err != nil {
		log.Fatal(err)
	}
	defer raw.Body.Close()

	b, err := io.ReadAll(raw.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("raw", string(b))
}
