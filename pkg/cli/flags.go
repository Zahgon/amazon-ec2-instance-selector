// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cli

import (
	"math"

	"github.com/spf13/pflag"

	"github.com/aws/amazon-ec2-instance-selector/v3/pkg/bytequantity"
)

const (
	maxInt    = int(^uint(0) >> 1)
	max32Int  = int(^uint32(0) >> 1)
	maxUint64 = math.MaxUint64
)

// RatioFlag creates and registers a flag accepting a ratio.
func (cl *CommandLineInterface) RatioFlag(name string, shorthand *string, defaultValue *string, description string) {
	_ = "STUB: not implemented"
	return
}

// IntMinMaxRangeFlags creates and registers a min, max, and helper flag each accepting an int.
func (cl *CommandLineInterface) IntMinMaxRangeFlags(name string, shorthand *string, defaultValue *int, description string) {
	_ = "STUB: not implemented"
	return
}

// Int32MinMaxRangeFlags creates and registers a min, max, and helper flag each accepting an int.
func (cl *CommandLineInterface) Int32MinMaxRangeFlags(name string, shorthand *string, defaultValue *int32, description string) {
	_ = "STUB: not implemented"
	return
}

// ByteQuantityMinMaxRangeFlags creates and registers a min, max, and helper flag each accepting a byte quantity like 512mb.
func (cl *CommandLineInterface) ByteQuantityMinMaxRangeFlags(name string, shorthand *string, defaultValue *bytequantity.ByteQuantity, description string) {
	_ = "STUB: not implemented"
	return
}

// Float64MinMaxRangeFlags creates and registers a min, max, and helper flag each accepting a float64.
func (cl *CommandLineInterface) Float64MinMaxRangeFlags(name string, shorthand *string, defaultValue *float64, description string) {
	_ = "STUB: not implemented"
	return
}

// ByteQuantityFlag creates and registers a flag accepting a byte quantity like 512mb.
func (cl *CommandLineInterface) ByteQuantityFlag(name string, shorthand *string, defaultValue *bytequantity.ByteQuantity, description string) {
	_ = "STUB: not implemented"
	return
}

// IntFlag creates and registers a flag accepting an Integer.
func (cl *CommandLineInterface) IntFlag(name string, shorthand *string, defaultValue *int, description string) {
	_ = "STUB: not implemented"
	return
}

// StringFlag creates and registers a flag accepting a String and a validator function.
// The validator function is provided so that more complex flags can be created from a string input.
func (cl *CommandLineInterface) StringFlag(name string, shorthand *string, defaultValue *string, description string, validationFn validator) {
	_ = "STUB: not implemented"
	return
}

// StringSliceFlag creates and registers a flag accepting a list of strings.
func (cl *CommandLineInterface) StringSliceFlag(name string, shorthand *string, defaultValue []string, description string) {
	_ = "STUB: not implemented"
	return
}

// RegexFlag creates and registers a flag accepting a string and validates that it is a valid regex.
func (cl *CommandLineInterface) RegexFlag(name string, shorthand *string, defaultValue *string, description string) {
	_ = "STUB: not implemented"
	return
}

// PathFlag creates and registers a flag accepting a string representing a path and validates that it is a valid path.
func (cl *CommandLineInterface) PathFlag(name string, shorthand *string, defaultValue *string, description string) {
	_ = "STUB: not implemented"
	return
}

// StringOptionsFlag creates and registers a flag accepting a string and valid options for use in validation.
func (cl *CommandLineInterface) StringOptionsFlag(name string, shorthand *string, defaultValue *string, description string, validOpts []string) {
	_ = "STUB: not implemented"
	return
}

// BoolFlag creates and registers a flag accepting a boolean.
func (cl *CommandLineInterface) BoolFlag(name string, shorthand *string, defaultValue *bool, description string) {
	_ = "STUB: not implemented"
	return
}

// ConfigStringFlag creates and registers a flag accepting a String for configuration purposes.
// Config flags will be grouped at the bottom in the output of --help.
func (cl *CommandLineInterface) ConfigStringFlag(name string, shorthand *string, defaultValue *string, description string, validationFn validator) {
	_ = "STUB: not implemented"
	return
}

// ConfigStringSliceFlag creates and registers a flag accepting a list of strings.
// Config flags will be grouped at the bottom in the output of --help.
func (cl *CommandLineInterface) ConfigStringSliceFlag(name string, shorthand *string, defaultValue []string, description string) {
	_ = "STUB: not implemented"
	return
}

// ConfigPathFlag creates and registers a flag accepting a string representing a path and validates that it is a valid path.
// Config flags will be grouped at the bottom in the output of --help.
func (cl *CommandLineInterface) ConfigPathFlag(name string, shorthand *string, defaultValue *string, description string) {
	_ = "STUB: not implemented"
	return
}

// ConfigIntFlag creates and registers a flag accepting an Integer for configuration purposes.
// Config flags will be grouped at the bottom in the output of --help.
func (cl *CommandLineInterface) ConfigIntFlag(name string, shorthand *string, defaultValue *int, description string) {
	_ = "STUB: not implemented"
	return
}

// ConfigBoolFlag creates and registers a flag accepting a boolean for configuration purposes.
// Config flags will be grouped at the bottom in the output of --help.
func (cl *CommandLineInterface) ConfigBoolFlag(name string, shorthand *string, defaultValue *bool, description string) {
	_ = "STUB: not implemented"
	return
}

// ConfigStringOptionsFlag creates and registers a flag accepting a string and valid options for use in validation.
// Config flags will be grouped at the bottom in the output of --help.
func (cl *CommandLineInterface) ConfigStringOptionsFlag(name string, shorthand *string, defaultValue *string, description string, validOpts []string) {
	_ = "STUB: not implemented"
	return
}

// SuiteBoolFlag creates and registers a flag accepting a boolean for aggregate filters.
// Suite flags will be grouped in the middle of the output --help.
func (cl *CommandLineInterface) SuiteBoolFlag(name string, shorthand *string, defaultValue *bool, description string) {
	_ = "STUB: not implemented"
	return
}

// SuiteStringFlag creates and registers a flag accepting a string for aggreagate filters.
// Suite flags will be grouped in the middle of the output --help.
func (cl *CommandLineInterface) SuiteStringFlag(name string, shorthand *string, defaultValue *string, description string, validationFn validator) {
	_ = "STUB: not implemented"
	return
}

// SuiteStringOptionsFlag creates and registers a flag accepting a string and valid options for use in validation.
// Suite flags will be grouped in the middle of the output --help.
func (cl *CommandLineInterface) SuiteStringOptionsFlag(name string, shorthand *string, defaultValue *string, description string, validOpts []string) {
	_ = "STUB: not implemented"
	return
}

// SuiteStringSliceFlag creates and registers a flag accepting a list of strings.
// Suite flags will be grouped in the middle of the output --help.
func (cl *CommandLineInterface) SuiteStringSliceFlag(name string, shorthand *string, defaultValue []string, description string) {
	_ = "STUB: not implemented"
	return
}

// BoolFlagOnFlagSet creates and registers a flag accepting a boolean for configuration purposes.
func (cl *CommandLineInterface) BoolFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *bool, description string) {
	_ = "STUB: not implemented"
	return
}

// IntMinMaxRangeFlagOnFlagSet creates and registers a min, max, and helper flag each accepting an int.
func (cl *CommandLineInterface) IntMinMaxRangeFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *int, description string) {
	_ = "STUB: not implemented"
	return
}

// Int32MinMaxRangeFlagOnFlagSet creates and registers a min, max, and helper flag each accepting an int.
func (cl *CommandLineInterface) Int32MinMaxRangeFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *int32, description string) {
	_ = "STUB: not implemented"
	return
}

// Float64MinMaxRangeFlagOnFlagSet creates and registers a min, max, and helper flag each accepting a float64.
func (cl *CommandLineInterface) Float64MinMaxRangeFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *float64, description string) {
	_ = "STUB: not implemented"
	return
}

// ByteQuantityMinMaxRangeFlagOnFlagSet creates and registers a min, max, and helper flag each accepting a ByteQuantity like 5mb or 12gb.
func (cl *CommandLineInterface) ByteQuantityMinMaxRangeFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *bytequantity.ByteQuantity, description string) {
	_ = "STUB: not implemented"
	return
}

// ByteQuantityFlagOnFlagSet creates and registers a flag accepting a ByteQuantity.
func (cl *CommandLineInterface) ByteQuantityFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *bytequantity.ByteQuantity, description string) {
	_ = "STUB: not implemented"
	return
}

// IntFlagOnFlagSet creates and registers a flag accepting an int.
func (cl *CommandLineInterface) IntFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *int, description string) {
	_ = "STUB: not implemented"
	return
}

// Int32FlagOnFlagSet creates and registers a flag accepting an int.
func (cl *CommandLineInterface) Int32FlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *int32, description string) {
	_ = "STUB: not implemented"
	return
}

// Float64FlagOnFlagSet creates and registers a flag accepting a float64.
func (cl *CommandLineInterface) Float64FlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *float64, description string) {
	_ = "STUB: not implemented"
	return
}

// StringFlagOnFlagSet creates and registers a flag accepting a string and a validator function.
// The validator function is provided so that more complex flags can be created from a string input.
func (cl *CommandLineInterface) StringFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *string, description string, processorFn processor, validationFn validator) {
	_ = "STUB: not implemented"
	return
}

// StringOptionsFlagOnFlagSet creates and registers a flag accepting a string with valid options.
// The validOpts slice of strings will be used to perform validation.
func (cl *CommandLineInterface) StringOptionsFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *string, description string, validOpts []string) {
	_ = "STUB: not implemented"
	return
}

// StringSliceFlagOnFlagSet creates and registers a flag accepting a string slice.
func (cl *CommandLineInterface) StringSliceFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue []string, description string) {
	_ = "STUB: not implemented"
	return
}

// RegexFlagOnFlagSet creates and registers a flag accepting a string slice of regular expressions.
func (cl *CommandLineInterface) RegexFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *string, description string) {
	_ = "STUB: not implemented"
	return
}

// PathFlagOnFlagSet creates and registers a flag accepting a string as a path.
func (cl *CommandLineInterface) PathFlagOnFlagSet(flagSet *pflag.FlagSet, name string, shorthand *string, defaultValue *string, description string) {
	_ = "STUB: not implemented"
	return
}
