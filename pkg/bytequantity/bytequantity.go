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

package bytequantity

import (
	"math"
)

const (
	/// Examples:          1mb, 1 gb, 1.0tb, 1mib, 2g, 2.001 t.
	byteQuantityRegex = `^([0-9]+\.?[0-9]{0,3})[ ]?(mi?b?|gi?b?|ti?b?)?$`
	mib               = "MiB"
	gib               = "GiB"
	tib               = "TiB"
	gbConvert         = 1 << 10
	tbConvert         = gbConvert << 10
	maxGiB            = math.MaxUint64 / gbConvert
	maxTiB            = math.MaxUint64 / tbConvert
)

// ByteQuantity is a data type representing a byte quantity.
type ByteQuantity struct {
	Quantity uint64
}

// ParseToByteQuantity parses a string representation of a byte quantity to a ByteQuantity type.
// A unit can be appended such as 16 GiB. If no unit is appended, GiB is assumed.
func ParseToByteQuantity(byteQuantityStr string) (ByteQuantity, error) {
	_ = "STUB: not implemented"
	return *new(ByteQuantity), nil
}

// mib

// need error here so that this quantity doesn't bind in the local scope

// gib

// tib

// FromTiB returns a byte quantity of the passed in tebibytes quantity.
func FromTiB(tib uint64) ByteQuantity { _ = "STUB: not implemented"; return *new(ByteQuantity) }

// FromGiB returns a byte quantity of the passed in gibibytes quantity.
func FromGiB(gib uint64) ByteQuantity { _ = "STUB: not implemented"; return *new(ByteQuantity) }

// FromMiB returns a byte quantity of the passed in mebibytes quantity.
func FromMiB(mib uint64) ByteQuantity { _ = "STUB: not implemented"; return *new(ByteQuantity) }

// StringMiB returns a byte quantity in a mebibytes string representation.
func (bq ByteQuantity) StringMiB() string { _ = "STUB: not implemented"; return "" }

// StringGiB returns a byte quantity in a gibibytes string representation.
func (bq ByteQuantity) StringGiB() string { _ = "STUB: not implemented"; return "" }

// StringTiB returns a byte quantity in a tebibytes string representation.
func (bq ByteQuantity) StringTiB() string { _ = "STUB: not implemented"; return "" }

// MiB returns a byte quantity in mebibytes.
func (bq ByteQuantity) MiB() float64 { _ = "STUB: not implemented"; return 0 }

// GiB returns a byte quantity in gibibytes.
func (bq ByteQuantity) GiB() float64 { _ = "STUB: not implemented"; return 0 }

// TiB returns a byte quantity in tebibytes.
func (bq ByteQuantity) TiB() float64 { _ = "STUB: not implemented"; return 0 }
