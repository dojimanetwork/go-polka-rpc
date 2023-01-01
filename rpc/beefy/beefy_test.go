package beefy

import (
	"os"
	"testing"

	"github.com/dojimanetwork/go-polka-rpc/v5/client"
	"github.com/dojimanetwork/go-polka-rpc/v5/config"
)

var testBeefy Beefy

func TestMain(m *testing.M) {
	cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	testBeefy = NewBeefy(cl)
	os.Exit(m.Run())
}
