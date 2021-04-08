package offchain

import (
	"os"
	"testing"

	"github.com/centrifuge/go-substrate-rpc-client/v2/client"
	"github.com/centrifuge/go-substrate-rpc-client/v2/config"
)

var offchain *Offchain

func TestMain(m *testing.M) {
	cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	offchain = NewOffchain(cl)
	os.Exit(m.Run())
}
