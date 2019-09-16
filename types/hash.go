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

package types

import "fmt"

// H160 is a hash containing 160 bits (20 bytes), typically used in blocks, extrinsics and as a sane default
type H160 [20]byte

// NewH160 creates a new H160 type
func NewH160(b [20]byte) H160 {
	return H160(b)
}

// Hex returns a hex string representation of the value (not of the encoded value)
func (h H160) Hex() string {
	return fmt.Sprintf("%#x", h[:])
}

// H256 is a hash containing 256 bits (32 bytes), typically used in blocks, extrinsics and as a sane default
type H256 [32]byte

// NewH256 creates a new H256 type
func NewH256(b [32]byte) H256 {
	return H256(b)
}

// Hex returns a hex string representation of the value (not of the encoded value)
func (h H256) Hex() string {
	return fmt.Sprintf("%#x", h[:])
}

// H512 is a hash containing 512 bits (64 bytes), typically used for signature
type H512 [64]byte

// NewH512 creates a new H512 type
func NewH512(b [64]byte) H512 {
	return H512(b)
}

// Hex returns a hex string representation of the value (not of the encoded value)
func (h H512) Hex() string {
	return fmt.Sprintf("%#x", h[:])
}

// Hash is the default hash that is used across the system. It is just a thin wrapper around H256
type Hash H256

// NewHash creates a new Hash type
func NewHash(b [32]byte) Hash {
	return Hash(b)
}

// Hex returns a hex string representation of the value (not of the encoded value)
func (h Hash) Hex() string {
	return fmt.Sprintf("%#x", h[:])
}
