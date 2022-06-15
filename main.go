package main

import (
	"fmt"
	"github.com/hashicorp/vault/vault"
)

func main() {
	cluster := vault.NewTestCluster(nil, nil, &vault.TestClusterOptions{
		HandlerFunc: func(props *vault.HandlerProperties) http.Handler {
			fmt.Println("hi")
		},
	})

	cluster.Start()
	defer cluster.Cleanup()
}
