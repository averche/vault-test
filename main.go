package main

import (
	"net/http"
	"fmt"

	"github.com/hashicorp/vault/vault"
)

func main() {
	cluster := vault.NewTestCluster(nil, nil, &vault.TestClusterOptions{
		HandlerFunc: func(props *vault.HandlerProperties) http.Handler {
			fmt.Println("hi")
			return nil
		},
	})

	cluster.Start()
	defer cluster.Cleanup()
}
