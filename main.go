package main

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	ctx := context.Background()

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken("my-token"); err != nil {
		log.Fatal(err)
	}

	// write a secret
	_, err = client.Secrets.KvV2Write(
		ctx,
		"my-secret",
		schema.KvV2WriteRequest{
			Data: map[string]any{
				"password1": "abc123",
				"password2": "correct horse battery staple",
			},
		},
		vault.WithMountPath("secret"),
		vault.WithRequestCallbacks(func(r *http.Request) {
			log.Println(*r)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret written successfully")

	// read a secret
	s, err := client.Secrets.KvV2Read(ctx, "my-secret", vault.WithMountPath("secret"))
	if err != nil {
		log.Fatal(err)
	}

	v := s.Data.Data["password1"]

	log.Println("secret retrieved:", v)
}
