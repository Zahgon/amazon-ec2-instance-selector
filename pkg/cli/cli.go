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
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type runFunc = func(cmd *cobra.Command, args []string)

// New creates an instance of CommandLineInterface.
func New(binaryName string, shortUsage string, longUsage, examples string, run runFunc) CommandLineInterface {
	_ = "STUB: not implemented"
	return *new(CommandLineInterface)
}

// ParseFlags will parse flags registered in this instance of CLI from os.Args.
func (cl *CommandLineInterface) ParseFlags() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// Remove Suite Flags so that args only include Config and Filter Flags
		nil
}

// This parses Config and Filter flags only

// Remove Config and Filter flags so that only suite flags are parsed

// Add suite flags to Command flagset so that other processing can occur
// This has to be done after usage is printed so that the flagsets can be grouped properly when printed

// ParseAndValidateFlags will parse flags registered in this instance of CLI from os.Args
// and then perform validation.
func (cl *CommandLineInterface) ParseAndValidateFlags() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProcessFlags iterates through any registered processors and executes them
// Processors are executed before validators.
func (cl *CommandLineInterface) ProcessFlags() error { _ = "STUB: not implemented"; return nil }

// ValidateFlags iterates through any registered validators and executes them.
func (cl *CommandLineInterface) ValidateFlags() error { _ = "STUB: not implemented"; return nil }

func removeIntersectingArgs(flagSet *pflag.FlagSet) []string { _ = "STUB: not implemented"; return nil }

func shorthandLookup(flagSet *pflag.FlagSet, arg string) *pflag.Flag {
	_ = "STUB: not implemented"
	return nil
}

func (cl *CommandLineInterface) setUsageTemplate() { _ = "STUB: not implemented"; return }

// SetUntouchedFlagValuesToNil iterates through all flags and sets their value to nil if they were not specifically set by the user
// This allows for a specified value, a negative value (like false or empty string), or an unspecified (nil) entry.
func (cl *CommandLineInterface) SetUntouchedFlagValuesToNil() error {
	_ = "STUB: not implemented"
	return nil
}

// If nilDefaults entry for flag is set to false, do not change default

// ProcessRangeFilterFlags sets min and max to the appropriate 0 or max bounds based on the 3-tuple that a user specifies for base flag, min, and/or max.
func (cl *CommandLineInterface) ProcessRangeFilterFlags() error {
	_ = "STUB: not implemented"
	return nil
}
