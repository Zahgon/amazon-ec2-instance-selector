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

package ec2pricing

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
)

const (
	productDescription = "Linux/UNIX (Amazon VPC)"
	serviceCode        = "AmazonEC2"
)

var DefaultSpotDaysBack = 30

// EC2Pricing is the public struct to interface with AWS pricing APIs.
type EC2Pricing struct {
	ODPricing   *OnDemandPricing
	SpotPricing *SpotPricing
	logger      *log.Logger
}

// EC2PricingIface is the EC2Pricing interface mainly used to mock out ec2pricing during testing.
type EC2PricingIface interface {
	GetOnDemandInstanceTypeCost(ctx context.Context, instanceType ec2types.InstanceType) (float64, error)
	GetSpotInstanceTypeNDayAvgCost(ctx context.Context, instanceType ec2types.InstanceType, availabilityZones []string, days int) (float64, error)
	RefreshOnDemandCache(ctx context.Context) error
	RefreshSpotCache(ctx context.Context, days int) error
	OnDemandCacheCount() int
	SpotCacheCount() int
	Save() error
	SetLogger(*log.Logger)
}

// use us-east-1 since pricing only has endpoints in us-east-1 and ap-south-1
// TODO: In the future we may want to allow the client to select which endpoint is used through some mechanism
//
//	but that would likely happen through overriding this entire function as its signature is fixed
func modifyPricingRegion(opt *pricing.Options) { _ = "STUB: not implemented"; return }

// New creates an instance of instance-selector EC2Pricing.
func New(ctx context.Context, cfg aws.Config) (*EC2Pricing, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewWithCache(ctx context.Context, cfg aws.Config, ttl time.Duration, cacheDir string) (*EC2Pricing, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *EC2Pricing) SetLogger(logger *log.Logger) { _ = "STUB: not implemented"; return }

// OnDemandCacheCount returns the number of items in the OD cache.
func (p *EC2Pricing) OnDemandCacheCount() int { _ = "STUB: not implemented"; return 0 }

// SpotCacheCount returns the number of items in the spot cache.
func (p *EC2Pricing) SpotCacheCount() int { _ = "STUB: not implemented"; return 0 }

// GetSpotInstanceTypeNDayAvgCost retrieves the spot price history for a given AZ from the past N days and averages the price
// Passing an empty list for availabilityZones will retrieve avg cost for all AZs in the current AWSSession's region.
func (p *EC2Pricing) GetSpotInstanceTypeNDayAvgCost(ctx context.Context, instanceType ec2types.InstanceType, availabilityZones []string, days int) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetOnDemandInstanceTypeCost retrieves the on-demand hourly cost for the specified instance type.
func (p *EC2Pricing) GetOnDemandInstanceTypeCost(ctx context.Context, instanceType ec2types.InstanceType) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// RefreshOnDemandCache makes a bulk request to the pricing api to retrieve all instance type pricing and stores them in a local cache.
func (p *EC2Pricing) RefreshOnDemandCache(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// RefreshSpotCache makes a bulk request to the ec2 api to retrieve all spot instance type pricing and stores them in a local cache.
func (p *EC2Pricing) RefreshSpotCache(ctx context.Context, days int) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *EC2Pricing) Save() error { _ = "STUB: not implemented"; return nil }
