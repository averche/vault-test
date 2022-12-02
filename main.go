package main

import (
	"context"
	"log"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	ctx := context.Background()

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken("my-token"); err != nil {
		log.Fatal(err)
	}

	// write a secret
	_, err = client.Write(ctx, "/secret/data/my-secret", map[string]any{
		"data": map[string]any{
			"password1": "abc123",
			"password2": "correct horse battery staple",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret written succesffully")

	// read a secret
	r, err := client.Read(ctx, "/secret/data/my-secret")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret retrieved:", r.Data["data"])
}
