// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
// Copyright (C) 2019  Centrifuge GmbH
//
// This file is part of Go Substrate RPC Client (GSRPC).
//
// GSRPC is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// GSRPC is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package tests_e2e //nolint:stylecheck,golint

import (
	"fmt"
	"testing"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/config"
	"github.com/stretchr/testify/assert"
)

func TestGetBlockHashAndVersion(t *testing.T) {
	api, err := gsrpc.NewSubstrateAPI(config.NewDefaultConfig().RPCURL)
	assert.NoError(t, err)
	hash, err := api.RPC.Chain.GetBlockHashLatest()
	assert.NoError(t, err)
	runtimeVersion, err := api.RPC.State.GetRuntimeVersionLatest()
	assert.NoError(t, err)

	fmt.Printf("Connected to node %v | latest block hash: %v | authoringVersion: %v | specVersion: %v | "+
		"implVersion: %v\n", (*api.Client).GetURL(), hash.Hex(), runtimeVersion.AuthoringVersion,
		runtimeVersion.SpecVersion, runtimeVersion.ImplVersion)
}
