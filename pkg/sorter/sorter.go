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

package sorter

import (
	"reflect"

	"github.com/aws/amazon-ec2-instance-selector/v3/pkg/instancetypes"
)

const (
	// Sort direction.

	SortAscending  = "ascending"
	SortAsc        = "asc"
	SortDescending = "descending"
	SortDesc       = "desc"

	// Not all fields can be reached through a json path (Ex: gpu count)
	// so we have special flags for such cases.

	GPUCountField              = "gpus"
	InferenceAcceleratorsField = "inference-accelerators"

	// shorthand flags.

	VCPUs                          = "vcpus"
	Memory                         = "memory"
	GPUMemoryTotal                 = "gpu-memory-total"
	NetworkInterfaces              = "network-interfaces"
	SpotPrice                      = "spot-price"
	ODPrice                        = "on-demand-price"
	InstanceStorage                = "instance-storage"
	EBSOptimizedBaselineBandwidth  = "ebs-optimized-baseline-bandwidth"
	EBSOptimizedBaselineThroughput = "ebs-optimized-baseline-throughput"
	EBSOptimizedBaselineIOPS       = "ebs-optimized-baseline-iops"

	// JSON field paths for shorthand flags.

	instanceNamePath                   = ".InstanceType"
	vcpuPath                           = ".VCpuInfo.DefaultVCpus"
	memoryPath                         = ".MemoryInfo.SizeInMiB"
	gpuMemoryTotalPath                 = ".GpuInfo.TotalGpuMemoryInMiB"
	networkInterfacesPath              = ".NetworkInfo.MaximumNetworkInterfaces"
	spotPricePath                      = ".SpotPrice"
	odPricePath                        = ".OndemandPricePerHour"
	instanceStoragePath                = ".InstanceStorageInfo.TotalSizeInGB"
	ebsOptimizedBaselineBandwidthPath  = ".EbsInfo.EbsOptimizedInfo.BaselineBandwidthInMbps"
	ebsOptimizedBaselineThroughputPath = ".EbsInfo.EbsOptimizedInfo.BaselineThroughputInMBps"
	ebsOptimizedBaselineIOPSPath       = ".EbsInfo.EbsOptimizedInfo.BaselineIops"
)

// sorterNode represents a sortable instance type which holds the value
// to sort by instance sort.
type sorterNode struct {
	instanceType *instancetypes.Details
	fieldValue   reflect.Value
}

// sorter is used to sort instance types based on a sorting field
// and direction.
type sorter struct {
	sorters      []*sorterNode
	sortField    string
	isDescending bool
}

// Sort sorts the given instance types by the given field in the given direction
//
// sortField is a json path to a field in the instancetypes.Details struct which represents
// the field to sort instance types by (Ex: ".MemoryInfo.SizeInMiB"). Quantity flags present
// in the CLI (memory, gpus, etc.) are also accepted.
//
// sortDirection represents the direction to sort in. Valid options: "ascending", "asc", "descending", "desc".
func Sort(instanceTypes []*instancetypes.Details, sortField string, sortDirection string) ([]*instancetypes.Details, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// determine if user used a shorthand for sorting flag

// newSorter creates a new Sorter object to be used to sort the given instance types
// based on the sorting field and direction
//
// sortField is a json path to a field in the instancetypes.Details struct which represents
// the field to sort instance types by (Ex: ".MemoryInfo.SizeInMiB").
//
// sortDirection represents the direction to sort in. Valid options: "ascending", "asc", "descending", "desc".
func newSorter(instanceTypes []*instancetypes.Details, sortField string, sortDirection string) (*sorter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create sorterNode objects for each instance type

// formatSortField reformats sortField to match the expected json path format
// of the json lookup library. Format is unchanged if the sorting field
// matches one of the special flags.
func formatSortField(sortField string) string {
	_ = "STUB: not implemented"
	// check to see if the sorting field matched one of the special exceptions
	return ""
}

// newSorterNode creates a new sorterNode object which represents the given instance type
// and can be used in sorting of instance types based on the given sortField.
func newSorterNode(instanceType *instancetypes.Details, sortField string) (*sorterNode, error) {
	_ = "STUB: not implemented"
	// some important fields (such as gpu count) can not be accessed directly in the instancetypes.Details
	// struct, so we have special hard-coded flags to handle such cases
	return nil, nil
}

// convert instance type into json

// unmarshal json instance types in order to get proper format
// for json path parsing

// get the desired field from the json data based on the passed in
// json path

// handle case where parent objects in path are null
// by setting result to nil

// sort the instance types in the Sorter based on the Sorter's sort field and
// direction.
func (s *sorter) sort() error { _ = "STUB: not implemented"; return nil }

// isLess determines whether the first value (valI) is less than the
// second value (valJ) or not.
func isLess(valI, valJ reflect.Value, isDescending bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// if valJ is not an int (can occur if the other value is nil)
// then valI is less. This will bubble invalid values to the end

// if valJ is not a uint (can occur if the other value is nil)
// then valI is less. This will bubble invalid values to the end

// if valJ is not a float (can occur if the other value is nil)
// then valI is less. This will bubble invalid values to the end

// if valJ is not a string (can occur if the other value is nil)
// then valI is less. This will bubble invalid values to the end

// Handle nil values by making non nil values always less than the nil values. That way the
// nil values can be bubbled up to the end of the list.

// if valJ is not a bool (can occur if the other value is nil)
// then valI is less. This will bubble invalid values to the end

// handle invalid values (like nil values) by making valid values
// always less than the invalid values. That way the invalid values
// always bubble up to the end of the list

// unsortable value

// instanceTypes returns the list of instance types held in the Sorter.
func (s *sorter) instanceTypes() []*instancetypes.Details { _ = "STUB: not implemented"; return nil }

// helper functions for special sorting fields

// getTotalGpusCount calculates the number of gpus in the given instance type.
func getTotalGpusCount(instanceType *instancetypes.Details) *int32 {
	_ = "STUB: not implemented"
	return nil
}

// getTotalAcceleratorsCount calculates the total number of inference accelerators
// in the given instance type.
func getTotalAcceleratorsCount(instanceType *instancetypes.Details) *int32 {
	_ = "STUB: not implemented"
	return nil
}
