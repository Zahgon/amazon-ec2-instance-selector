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

// Package cli provides functions to build the selector command line interface
package cli

import (
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/aws/amazon-ec2-instance-selector/v3/pkg/bytequantity"
	"github.com/aws/amazon-ec2-instance-selector/v3/pkg/selector"
)

const (
	// Usage Template to run on --help.
	usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Filter Flags:
{{.LocalNonPersistentFlags.FlagUsages | trimTrailingWhitespaces}}
%s
Global Flags:
{{.PersistentFlags.FlagUsages | trimTrailingWhitespaces}}

{{end}}`
)

// validator defines the function for providing validation on a flag.
type validator = func(val interface{}) error

// processor defines the function for providing mutating processing on a flag.
type processor = func(val interface{}) error

// CommandLineInterface is a type to group CLI funcs and state.
type CommandLineInterface struct {
	Command     *cobra.Command
	Flags       map[string]interface{}
	nilDefaults map[string]bool
	rangeFlags  map[string]bool
	validators  map[string]validator
	processors  map[string]processor
	suiteFlags  *pflag.FlagSet
}

// Float64Me takes an interface and returns a pointer to a float64 value
// If the underlying interface kind is not float64 or *float64 then nil is returned.
func (*CommandLineInterface) Float64Me(i interface{}) *float64 {
	_ = "STUB: not implemented"
	return nil
}

// IntMe takes an interface and returns a pointer to an int value
// If the underlying interface kind is not int or *int then nil is returned.
func (*CommandLineInterface) IntMe(i interface{}) *int { _ = "STUB: not implemented"; return nil }

// Int32Me takes an interface and returns a pointer to an int value
// If the underlying interface kind is not int or *int then nil is returned.
func (*CommandLineInterface) Int32Me(i interface{}) *int32 { _ = "STUB: not implemented"; return nil }

// IntRangeMe takes an interface and returns a pointer to an IntRangeFilter value
// If the underlying interface kind is not IntRangeFilter or *IntRangeFilter then nil is returned.
func (*CommandLineInterface) IntRangeMe(i interface{}) *selector.IntRangeFilter {
	_ = "STUB: not implemented"
	return nil
}

// Int32RangeMe takes an interface and returns a pointer to an Int32RangeFilter value
// If the underlying interface kind is not Int32RangeFilter or *Int32RangeFilter then nil is returned.
func (*CommandLineInterface) Int32RangeMe(i interface{}) *selector.Int32RangeFilter {
	_ = "STUB: not implemented"
	return nil
}

// ByteQuantityRangeMe takes an interface and returns a pointer to a ByteQuantityRangeFilter value
// If the underlying interface kind is not ByteQuantityRangeFilter or *ByteQuantityRangeFilter then nil is returned.
func (*CommandLineInterface) ByteQuantityRangeMe(i interface{}) *selector.ByteQuantityRangeFilter {
	_ = "STUB: not implemented"
	return nil
}

// Float64RangeMe takes an interface and returns a pointer to a Float64RangeFilter value
// If the underlying interface kind is not Float64RangeFilter or *Float64RangeFilter then nil is returned.
func (*CommandLineInterface) Float64RangeMe(i interface{}) *selector.Float64RangeFilter {
	_ = "STUB: not implemented"
	return nil
}

// StringMe takes an interface and returns a pointer to a string value
// If the underlying interface kind is not string or *string then nil is returned.
func (*CommandLineInterface) StringMe(i interface{}) *string { _ = "STUB: not implemented"; return nil }

// BoolMe takes an interface and returns a pointer to a bool value
// If the underlying interface kind is not bool or *bool then nil is returned.
func (*CommandLineInterface) BoolMe(i interface{}) *bool { _ = "STUB: not implemented"; return nil }

// StringSliceMe takes an interface and returns a pointer to a string slice
// If the underlying interface kind is not []string or *[]string then nil is returned.
func (*CommandLineInterface) StringSliceMe(i interface{}) *[]string {
	_ = "STUB: not implemented"
	return nil
}

// RegexMe takes an interface and returns a pointer to a regex
// If the underlying interface kind is not regexp.Regexp or *regexp.Regexp then nil is returned.
func (*CommandLineInterface) RegexMe(i interface{}) *regexp.Regexp {
	_ = "STUB: not implemented"
	return nil
}

// ByteQuantityMe takes an interface and returns a pointer to a ByteQuantity
// If the underlying interface kind is not bytequantity.ByteQuantity or *bytequantity.ByteQuantity then nil is returned.
func (*CommandLineInterface) ByteQuantityMe(i interface{}) *bytequantity.ByteQuantity {
	_ = "STUB: not implemented"
	return nil
}
