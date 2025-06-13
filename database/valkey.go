package database

import "github.com/valkey-io/valkey-go"

var VClient valkey.Client

//valkey:6379
func init() {
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
		Password:    "861214959",
	})
	if err != nil {
		panic(err)
	}
	VClient = client
}
