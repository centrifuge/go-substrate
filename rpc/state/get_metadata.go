// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package state

import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func (s *State) GetMetadata(blockHash types.Hash) (*types.Metadata, error) {
	return s.getMetadata(&blockHash)
}

func (s *State) GetMetadataLatest() (*types.Metadata, error) {
	return s.getMetadata(nil)
}

func (s *State) getMetadata(blockHash *types.Hash) (*types.Metadata, error) {
	metadata := types.NewMetadata()

	var res string
	var err error
	if blockHash == nil {
		err = (*s.client).Call(&res, "state_getMetadata")
	} else {
		hexHash, err := types.Hex(*blockHash)
		if err != nil {
			return metadata, err
		}
		err = (*s.client).Call(&res, "state_getMetadata", hexHash)
		if err != nil {
			return metadata, err
		}
	}
	if err != nil {
		return metadata, err
	}

	err = types.DecodeFromHexString(res, metadata)
	return metadata, err
}
