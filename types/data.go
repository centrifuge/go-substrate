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
	"fmt"
	"io"

	"github.com/centrifuge/go-substrate-rpc-client/scale"
)

// Data is a raw data structure, containing raw bytes that are not decoded/encoded (without any length encoding).
// Be careful using this in your own structs – it only works as the last value in a struct since it will consume the
// remainder of the encoded data. The reason for this is that it does not contain any length encoding, so it would
// not know where to stop.
type Data []byte

// NewData creates a new Data type
func NewData(b []byte) Data {
	return Data(b)
}

// Encode implements encoding for Data, which just unwraps the bytes of Data
func (d Data) Encode(encoder scale.Encoder) error {
	return encoder.Write(d)
}

// Decode implements decoding for Data, which just reads all the remaining bytes into Data
func (d *Data) Decode(decoder scale.Decoder) error {
	for i := 0; true; i++ {
		b, err := decoder.ReadOneByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		*d = append((*d)[:i], b)
	}
	return nil
}

// Hex returns a hex string representation of the value
func (d Data) Hex() string {
	return fmt.Sprintf("%#x", d)
}
