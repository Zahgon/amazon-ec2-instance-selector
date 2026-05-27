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

// Package selector provides filtering logic for Amazon EC2 Instance Types based on declarative resource specfications.
package selector

import (
	"context"
	"log"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/aws/amazon-ec2-instance-selector/v3/pkg/instancetypes"
)

// Version is overridden at compilation with the version based on the git tag
var versionID = "dev"

const (
	locationFilterKey      = "location"
	zoneIDLocationType     = ec2types.LocationTypeAvailabilityZoneId
	zoneNameLocationType   = ec2types.LocationTypeAvailabilityZone
	regionNameLocationType = ec2types.LocationTypeRegion
	sdkName                = "instance-selector"

	// Filter Keys.

	cpuArchitecture                  = "cpuArchitecture"
	cpuManufacturer                  = "cpuManufacturer"
	usageClass                       = "usageClass"
	rootDeviceType                   = "rootDeviceType"
	hibernationSupported             = "hibernationSupported"
	vcpusRange                       = "vcpusRange"
	memoryRange                      = "memoryRange"
	gpuMemoryRange                   = "gpuMemoryRange"
	gpusRange                        = "gpusRange"
	gpuManufacturer                  = "gpuManufacturer"
	gpuModel                         = "gpuModel"
	inferenceAcceleratorsRange       = "inferenceAcceleratorsRange"
	inferenceAcceleratorManufacturer = "inferenceAcceleartorManufacturer"
	inferenceAcceleratorModel        = "inferenceAcceleratorModel"
	placementGroupStrategy           = "placementGroupStrategy"
	hypervisor                       = "hypervisor"
	baremetal                        = "baremetal"
	burstable                        = "burstable"
	fpga                             = "fpga"
	enaSupport                       = "enaSupport"
	efaSupport                       = "efaSupport"
	vcpusToMemoryRatio               = "vcpusToMemoryRatio"
	currentGeneration                = "currentGeneration"
	networkInterfaces                = "networkInterfaces"
	networkPerformance               = "networkPerformance"
	networkEncryption                = "networkEncryption"
	ipv6                             = "ipv6"
	allowList                        = "allowList"
	denyList                         = "denyList"
	instanceTypes                    = "instanceTypes"
	virtualizationType               = "virtualizationType"
	instanceStorageRange             = "instanceStorageRange"
	diskEncryption                   = "diskEncryption"
	diskType                         = "diskType"
	nvme                             = "nvme"
	ebsOptimized                     = "ebsOptimized"
	ebsOptimizedBaselineBandwidth    = "ebsOptimizedBaselineBandwidth"
	ebsOptimizedBaselineIOPS         = "ebsOptimizedBaselineIOPS"
	ebsOptimizedBaselineThroughput   = "ebsOptimizedBaselineThroughput"
	freeTier                         = "freeTier"
	autoRecovery                     = "autoRecovery"
	dedicatedHosts                   = "dedicatedHosts"
	generation                       = "generation"

	cpuArchitectureAMD64 = "amd64"

	virtualizationTypePV = "pv"

	pricePerHour = "pricePerHour"
)

// New creates an instance of Selector provided an aws session.
func New(ctx context.Context, cfg aws.Config) (*Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewWithCache creates an instance of Selector backed by an on-disk cache provided an aws session and cache configuration parameters.
func NewWithCache(ctx context.Context, cfg aws.Config, ttl time.Duration, cacheDir string) (*Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetLogger can be called to log more detailed logs about what selector is doing
// including things like API timings
// If SetLogger is not called, no logs will be displayed.
func (s *Selector) SetLogger(logger *log.Logger) { _ = "STUB: not implemented"; return }

// Save persists the selector cache data to disk if caching is configured.
func (s Selector) Save() error { _ = "STUB: not implemented"; return nil }

// Filter accepts a Filters struct which is used to select the available instance types
// matching the criteria within Filters and returns a simple list of instance type strings.
func (s Selector) Filter(ctx context.Context, filters Filters) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FilterVerbose accepts a Filters struct which is used to select the available instance types
// matching the criteria within Filters and returns a list instanceTypeInfo.
func (s Selector) FilterVerbose(ctx context.Context, filters Filters) ([]*instancetypes.Details, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FilterWithOutput accepts a Filters struct which is used to select the available instance types
// matching the criteria within Filters and returns a list of strings based on the custom outputFn.
func (s Selector) FilterWithOutput(ctx context.Context, filters Filters, outputFn InstanceTypesOutput) ([]string, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s Selector) truncateResults(maxResults *int, instanceTypeInfoSlice []*instancetypes.Details) ([]*instancetypes.Details, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// AggregateFilterTransform takes higher level filters which are used to affect multiple raw filters in an opinionated way.
func (s Selector) AggregateFilterTransform(ctx context.Context, filters Filters) (Filters, error) {
	_ = "STUB: not implemented"
	return *new(Filters), nil
}

// rawFilter accepts a Filters struct which is used to select the available instance types
// matching the criteria within Filters and returns the detailed specs of matching instance types.
func (s Selector) rawFilter(ctx context.Context, filters Filters) ([]*instancetypes.Details, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s Selector) prepareFilter(ctx context.Context, filters Filters, instanceTypeInfo instancetypes.Details, availabilityZones []string, locationInstanceOfferings map[ec2types.InstanceType]string) (*instancetypes.Details, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Price used to filter based on usage class

// If prices are fetched, populate the fields irrespective of the price filters

// If price filter is present, prices should be already fetched
// If prices are not fetched, filter should fail and the corresponding error is already printed

// If an empty slice is passed, treat the filter as nil

// filterToInstanceSpecMappingPairs is a map of filter name [key] to filter pair [value].
// A filter pair includes user input filter value and instance spec value retrieved from DescribeInstanceTypes

// sortInstanceTypeInfo will sort based on instance type info alpha-numerically.
func sortInstanceTypeInfo(instanceTypeInfoSlice []*instancetypes.Details) []*instancetypes.Details {
	_ = "STUB: not implemented"
	return nil
}

// executeFilters accepts a mapping of filter name to filter pairs which are iterated through
// to determine if the instance type matches the filter values.
func (s Selector) executeFilters(ctx context.Context, filterToInstanceSpecMapping map[string]filterPair, instanceType ec2types.InstanceType) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// exec executes a specific filterPair (user value & instance spec) with a specific instance type
// If the filterPair matches, true is returned.
func exec(instanceType ec2types.InstanceType, filterName string, filter filterPair) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// if filter is nil, user did not specify a filter, so skip evaluation

// Determine appropriate filter comparator by switching on filter type

// this allows us to copy a static pointer to f into filterOfPtrs
// since the pointer to f is updated on each loop iteration

// RetrieveInstanceTypesSupportedInLocations returns a map of instance type -> AZ or Region for all instance types supported in the intersected locations passed in
// The location can be a zone-id (ie. use1-az1), a zone-name (us-east-1a), or a region name (us-east-1).
// Note that zone names are not necessarily the same across accounts.
func (s Selector) RetrieveInstanceTypesSupportedInLocations(ctx context.Context, locations []string) (map[ec2types.InstanceType]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s Selector) getLocationType(ctx context.Context, location string) (ec2types.LocationType, error) {
	_ = "STUB: not implemented"
	return *new(ec2types.LocationType), nil
}

func isSupportedInLocation(instanceOfferings map[ec2types.InstanceType]string, instanceType ec2types.InstanceType) bool {
	_ = "STUB: not implemented"
	return false
}

func isInDenyList(denyRegex *regexp.Regexp, instanceTypeName ec2types.InstanceType) bool {
	_ = "STUB: not implemented"
	return false
}

func isInAllowList(allowRegex *regexp.Regexp, instanceTypeName ec2types.InstanceType) bool {
	_ = "STUB: not implemented"
	return false
}
