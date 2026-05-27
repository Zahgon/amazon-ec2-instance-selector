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

package selector

import (
	"regexp"

	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const (
	supported = "supported"
	required  = "required"
)

var (
	networkPerfRE = regexp.MustCompile(`[0-9]+ Gigabit`)
	generationRE  = regexp.MustCompile(`[a-zA-Z]+([0-9]+)`)
)

func isSupportedFromString(instanceTypeValue *string, target *string) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedFromStrings(instanceTypeValues []*string, target *string) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedWithRangeInt(instanceTypeValue *int, target *IntRangeFilter) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedWithFloat64(instanceTypeValue *float64, target *float64) bool {
	_ = "STUB: not implemented"
	return false
}

// compare up to values' two decimal floor

func isSupportedUsageClassType(instanceTypeValue []ec2types.UsageClassType, target *ec2types.UsageClassType) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedArchitectureType(instanceTypeValue []ec2types.ArchitectureType, target *ec2types.ArchitectureType) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedVirtualizationType(instanceTypeValue []ec2types.VirtualizationType, target *ec2types.VirtualizationType) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedInstanceTypeHypervisorType(instanceTypeValue ec2types.InstanceTypeHypervisor, target *ec2types.InstanceTypeHypervisor) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedRootDeviceType(instanceTypeValue []ec2types.RootDeviceType, target *ec2types.RootDeviceType) bool {
	_ = "STUB: not implemented"
	return false
}

func isMatchingCpuArchitecture(instanceTypeValue CPUManufacturer, target *CPUManufacturer) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedWithRangeInt64(instanceTypeValue *int64, target *IntRangeFilter) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedWithRangeInt32(instanceTypeValue *int32, target *Int32RangeFilter) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedWithRangeUint64(instanceTypeValue *int64, target *Uint64RangeFilter) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedWithRangeFloat64(instanceTypeValue *float64, target *Float64RangeFilter) bool {
	_ = "STUB: not implemented"
	return false
}

func isSupportedWithBool(instanceTypeValue *bool, target *bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Helper functions for aggregating data parsed from AWS API calls

func getTotalAcceleratorsCount(acceleratorInfo *ec2types.InferenceAcceleratorInfo) *int32 {
	_ = "STUB: not implemented"
	return nil
}

func getTotalGpusCount(gpusInfo *ec2types.GpuInfo) *int32 { _ = "STUB: not implemented"; return nil }

func getTotalGpuMemory(gpusInfo *ec2types.GpuInfo) *int64 { _ = "STUB: not implemented"; return nil }

func getGPUManufacturers(gpusInfo *ec2types.GpuInfo) []*string {
	_ = "STUB: not implemented"
	return nil
}

func getGPUModels(gpusInfo *ec2types.GpuInfo) []*string { _ = "STUB: not implemented"; return nil }

func getInferenceAcceleratorManufacturers(acceleratorInfo *ec2types.InferenceAcceleratorInfo) []*string {
	_ = "STUB: not implemented"
	return nil
}

func getInferenceAcceleratorModels(acceleratorInfo *ec2types.InferenceAcceleratorInfo) []*string {
	_ = "STUB: not implemented"
	return nil
}

func getNetworkPerformance(networkPerformance *string) *int { _ = "STUB: not implemented"; return nil }

func getInstanceStorage(instanceStorageInfo *ec2types.InstanceStorageInfo) *int64 {
	_ = "STUB: not implemented"
	return nil
}

func getDiskType(instanceStorageInfo *ec2types.InstanceStorageInfo) *string {
	_ = "STUB: not implemented"
	return nil
}

func getNVMESupport(instanceStorageInfo *ec2types.InstanceStorageInfo, ebsInfo *ec2types.EbsInfo) *bool {
	_ = "STUB: not implemented"
	return nil
}

func getDiskEncryptionSupport(instanceStorageInfo *ec2types.InstanceStorageInfo, ebsInfo *ec2types.EbsInfo) *bool {
	_ = "STUB: not implemented"
	return nil
}

func getEBSOptimizedBaselineBandwidth(ebsInfo *ec2types.EbsInfo) *int32 {
	_ = "STUB: not implemented"
	return nil
}

func getEBSOptimizedBaselineThroughput(ebsInfo *ec2types.EbsInfo) *float64 {
	_ = "STUB: not implemented"
	return nil
}

func getEBSOptimizedBaselineIOPS(ebsInfo *ec2types.EbsInfo) *int32 {
	_ = "STUB: not implemented"
	return nil
}

// getInstanceTypeGeneration returns the generation from an instance type name
// i.e. c7i.xlarge -> 7
// if any error occurs, 0 will be returned.
func getInstanceTypeGeneration(instanceTypeName string) *int { _ = "STUB: not implemented"; return nil }

// supportSyntaxToBool takes an instance spec field that uses ["unsupported", "supported", "required", or "default"]
// and transforms it to a *bool to use in filter execution.
func supportSyntaxToBool(instanceTypeSupport *string) *bool { _ = "STUB: not implemented"; return nil }

func calculateVCpusToMemoryRatio(vcpusVal *int32, memoryVal *int64) *float64 {
	_ = "STUB: not implemented"
	return nil
}

// normalize vcpus to a mebivcpu value

// Slice helper function

func contains(slice []*string, target string) bool { _ = "STUB: not implemented"; return false }
