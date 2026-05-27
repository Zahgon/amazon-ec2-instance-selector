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
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/patrickmn/go-cache"
)

const (
	SpotCacheFileName = "spot-pricing-cache.gob"
)

type SpotPricing struct {
	Region         string
	FullRefreshTTL time.Duration
	DirectoryPath  string
	cache          *cache.Cache
	ec2Client      ec2.DescribeSpotPriceHistoryAPIClient
	logger         *log.Logger
	sync.RWMutex
}

type spotPricingEntry struct {
	Timestamp time.Time
	SpotPrice float64
	Zone      string
}

func LoadSpotCacheOrNew(ctx context.Context, ec2Client ec2.DescribeSpotPriceHistoryAPIClient, region string, fullRefreshTTL time.Duration, directoryPath string, days int) (*SpotPricing, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start the cache refresh job

func loadSpotCacheFrom(itemTTL time.Duration, region string, expandedDirPath string) (*cache.Cache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSpotCacheFilePath(region string, directoryPath string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *SpotPricing) spotCacheRefreshJob(ctx context.Context, days int) {
	_ = "STUB: not implemented"
	return
}

func (c *SpotPricing) SetLogger(logger *log.Logger) { _ = "STUB: not implemented"; return }

func (c *SpotPricing) Refresh(ctx context.Context, days int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SpotPricing) Get(ctx context.Context, instanceType ec2types.InstanceType, zone string, days int) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *SpotPricing) contains(zone string, entries []*spotPricingEntry) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *SpotPricing) calculateSpotAggregate(spotPriceEntries []*spotPricingEntry) float64 {
	_ = "STUB: not implemented"
	return 0
}

// Sort slice by timestamp in descending order from the end time (most likely, now)

func (c *SpotPricing) filterOn(zone string, pricingEntries []*spotPricingEntry) []*spotPricingEntry {
	_ = "STUB: not implemented"
	return nil
}

// this takes the first zone, might be better to do all zones instead...

// Count of items in the cache.
func (c *SpotPricing) Count() int { _ = "STUB: not implemented"; return 0 }

func (c *SpotPricing) Save() error { _ = "STUB: not implemented"; return nil }

func (c *SpotPricing) Clear() error { _ = "STUB: not implemented"; return nil }

// fetchSpotPricingTimeSeries makes a bulk request to the ec2 api to retrieve all spot instance type pricing for the past n days
// If instanceType is empty, it will fetch for all instance types.
func (c *SpotPricing) fetchSpotPricingTimeSeries(ctx context.Context, instanceType ec2types.InstanceType, days int) (map[string][]*spotPricingEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Iterate through the Amazon S3 object pages.
