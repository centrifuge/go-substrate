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

import (
	"encoding/json"
)

// I8 is a signed 8-bit integer
type I8 int8

// NewI8 creates a new I8 type
func NewI8(i int8) I8 {
	return I8(i)
}

// UnmarshalJSON fills i with the JSON encoded byte array given by b
func (i *I8) UnmarshalJSON(b []byte) error {
	var tmp int8
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*i = I8(tmp)
	return nil
}

// MarshalJSON returns a JSON encoded byte array of i
func (i *I8) MarshalJSON() ([]byte, error) {
	return json.Marshal(int8(*i))
}

// I16 is a signed 16-bit integer
type I16 int16

// NewI16 creates a new 16 type
func NewI16(i int16) I16 {
	return I16(i)
}

// UnmarshalJSON fills i with the JSON encoded byte array given by b
func (i *I16) UnmarshalJSON(b []byte) error {
	var tmp int16
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*i = I16(tmp)
	return nil
}

// MarshalJSON returns a JSON encoded byte array of i
func (i *I16) MarshalJSON() ([]byte, error) {
	return json.Marshal(int16(*i))
}

// I32 is a signed 32-bit integer
type I32 int32

// NewI32 creates a new 32 type
func NewI32(i int32) I32 {
	return I32(i)
}

// UnmarshalJSON fills i with the JSON encoded byte array given by b
func (i *I32) UnmarshalJSON(b []byte) error {
	var tmp int32
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*i = I32(tmp)
	return nil
}

// MarshalJSON returns a JSON encoded byte array of i
func (i *I32) MarshalJSON() ([]byte, error) {
	return json.Marshal(int32(*i))
}

// I64 is a signed 64-bit integer
type I64 int64

// NewI64 creates a new 64 type
func NewI64(i int64) I64 {
	return I64(i)
}

// UnmarshalJSON fills i with the JSON encoded byte array given by b
func (i *I64) UnmarshalJSON(b []byte) error {
	var tmp int64
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*i = I64(tmp)
	return nil
}

// MarshalJSON returns a JSON encoded byte array of i
func (i *I64) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(*i))
}
