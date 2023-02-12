package main

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	ctx := context.Background()

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken("my-token"); err != nil {
		log.Fatal(err)
	}

	// write a secret
	_, err = client.Secrets.KVv2Write(ctx, "my-secret", schema.KVv2WriteRequest{
		Data: map[string]any{
			"password1": "abc123",
			"password2": "correct horse battery staple",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret written successfully")

	// read a secret
	s, err := client.Secrets.KVv2Read(ctx, "my-secret")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret retrieved:", s.Data)
}
