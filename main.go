package main

import (
	"context"
	"log"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	ctx := context.Background()

	// prepare a client with default configuration, except for the address
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// authenticate
	resp, err := client.Auth.PostAuthApproleLogin(ctx, vault.ApproleLoginRequest{
		RoleId:   "91504982-e048-3702-5dae-d3cde41a8b15",
		SecretId: "d5a5a11a-ed72-9dfb-d3e7-7be83454348b",
	})
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken(resp.Auth.ClientToken)

	// write a secret
	_, err = client.Write(ctx, "/secret/data/my-secret", map[string]interface{}{
		"data": map[string]interface{}{
			"password1": "abc123",
			"password2": "trustno1",
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
