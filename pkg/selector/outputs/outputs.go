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

// Package outputs provides types for implementing instance type output functions as well as prebuilt output functions.
package outputs

import (
	"reflect"

	"github.com/aws/amazon-ec2-instance-selector/v3/pkg/instancetypes"
)

const columnTag = "column"

// wideColumnsData stores the data that should be displayed on each column
// of a wide output row.
type wideColumnsData struct {
	instanceName       string `column:"Instance Type"`
	vcpu               int32  `column:"VCPUs"`
	memory             string `column:"Mem (GiB)"`
	hypervisor         string `column:"Hypervisor"`
	currentGen         bool   `column:"Current Gen"`
	hibernationSupport bool   `column:"Hibernation Support"`
	cpuArch            string `column:"CPU Arch"`
	networkPerformance string `column:"Network Performance"`
	eni                int32  `column:"ENIs"`
	gpu                int32  `column:"GPUs"`
	gpuMemory          string `column:"GPU Mem (GiB)"`
	gpuInfo            string `column:"GPU Info"`
	odPrice            string `column:"On-Demand Price/Hr"`
	spotPrice          string `column:"Spot Price/Hr"`
}

// SimpleInstanceTypeOutput is an OutputFn which outputs a slice of instance type names.
func SimpleInstanceTypeOutput(instanceTypeInfoSlice []*instancetypes.Details) []string {
	_ = "STUB: not implemented"
	return nil
}

// VerboseInstanceTypeOutput is an OutputFn which outputs a slice of instance type names.
func VerboseInstanceTypeOutput(instanceTypeInfoSlice []*instancetypes.Details) []string {
	_ = "STUB: not implemented"
	return nil
}

// TableOutputShort is an OutputFn which returns a CLI table for easy reading.
func TableOutputShort(instanceTypeInfoSlice []*instancetypes.Details) []string {
	_ = "STUB: not implemented"
	return nil
}

// TableOutputWide is an OutputFn which returns a detailed CLI table for easy reading.
func TableOutputWide(instanceTypeInfoSlice []*instancetypes.Details) []string {
	_ = "STUB: not implemented"
	return nil
}

// OneLineOutput is an output function which prints the instance type names on a single line separated by commas.
func OneLineOutput(instanceTypeInfoSlice []*instancetypes.Details) []string {
	_ = "STUB: not implemented"
	return nil
}

func formatFloat(f float64) string { _ = "STUB: not implemented"; return "" }

func reverse(s string) string { _ = "STUB: not implemented"; return "" }

// getWideColumnsData returns the column data necessary for a wide output for each of
// the given instance types.
func getWideColumnsData(instanceTypes []*instancetypes.Details) []*wideColumnsData {
	_ = "STUB: not implemented"
	return nil
}

// getUnderlyingValue returns the underlying value of the given
// reflect.Value type.
func getUnderlyingValue(value reflect.Value) interface{} { _ = "STUB: not implemented"; return nil }
