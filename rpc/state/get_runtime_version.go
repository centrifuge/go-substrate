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

func (s *State) GetRuntimeVersion(blockHash types.Hash) (*types.RuntimeVersion, error) {
	return s.getRuntimeVersion(&blockHash)
}

func (s *State) GetRuntimeVersionLatest() (*types.RuntimeVersion, error) {
	return s.getRuntimeVersion(nil)
}

func (s *State) getRuntimeVersion(blockHash *types.Hash) (*types.RuntimeVersion, error) {
	var runtimeVersion types.RuntimeVersion
	var err error
	if blockHash == nil {
		err = (*s.client).Call(&runtimeVersion, "state_getRuntimeVersion")
	} else {
		hexHash, err := types.Hex(*blockHash)
		if err != nil {
			return &runtimeVersion, err
		}
		err = (*s.client).Call(&runtimeVersion, "state_getRuntimeVersion", hexHash)
		if err != nil {
			return &runtimeVersion, err
		}
	}
	if err != nil {
		return &runtimeVersion, err
	}
	return &runtimeVersion, err
}

//func (s *State) getRuntimeVersion(blockHash *types.Hash) (*types.RuntimeVersion, error) {
//	runtimeVersion := types.NewRuntimeVersion()
//
//	var res string
//	var err error
//	if blockHash == nil {
//		err = (*s.client).Call(&res, "state_getRuntimeVersion")
//	} else {
//		err = (*s.client).Call(&res, "state_getRuntimeVersion", *blockHash)
//	}
//	if err != nil {
//		return runtimeVersion, err
//	}
//
//	err = types.DecodeFromHexString(res, runtimeVersion)
//	return runtimeVersion, err
//}
