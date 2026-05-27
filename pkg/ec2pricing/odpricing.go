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

	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	pricingtypes "github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/patrickmn/go-cache"
)

const (
	ODCacheFileName = "on-demand-pricing-cache.json"
)

type OnDemandPricing struct {
	Region         string
	FullRefreshTTL time.Duration
	DirectoryPath  string
	cache          *cache.Cache
	pricingClient  pricing.GetProductsAPIClient
	logger         *log.Logger
	sync.RWMutex
}

type PricingList struct {
	Product         PricingListProduct `json:"product"`
	ServiceCode     string             `json:"serviceCode"`
	Terms           ProductTerms       `json:"terms"`
	Version         string             `json:"version"`
	PublicationDate string             `json:"publicationDate"`
}

type PricingListProduct struct {
	ProductFamily     string            `json:"productFamily"`
	ProductAttributes map[string]string `json:"attributes"`
	SKU               string            `json:"sku"`
}

type ProductTerms struct {
	OnDemand map[string]ProductPricingInfo `json:"OnDemand"`
	Reserved map[string]ProductPricingInfo `json:"Reserved"`
}

type ProductPricingInfo struct {
	PriceDimensions map[string]PriceDimensionInfo `json:"priceDimensions"`
	SKU             string                        `json:"sku"`
	EffectiveDate   string                        `json:"effectiveDate"`
	OfferTermCode   string                        `json:"offerTermCode"`
	TermAttributes  map[string]string             `json:"termAttributes"`
}

type PriceDimensionInfo struct {
	Unit         string            `json:"unit"`
	EndRange     string            `json:"endRange"`
	Description  string            `json:"description"`
	AppliesTo    []string          `json:"appliesTo"`
	RateCode     string            `json:"rateCode"`
	BeginRange   string            `json:"beginRange"`
	PricePerUnit map[string]string `json:"pricePerUnit"`
}

func LoadODCacheOrNew(ctx context.Context, pricingClient pricing.GetProductsAPIClient, region string, fullRefreshTTL time.Duration, directoryPath string) (*OnDemandPricing, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start the cache refresh job

func loadODCacheFrom(itemTTL time.Duration, region string, expandedDirPath string) (*cache.Cache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getODCacheFilePath(region string, directoryPath string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *OnDemandPricing) odCacheRefreshJob(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *OnDemandPricing) SetLogger(logger *log.Logger) { _ = "STUB: not implemented"; return }

func (c *OnDemandPricing) Refresh(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *OnDemandPricing) Get(ctx context.Context, instanceType ec2types.InstanceType) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Count of items in the cache.
func (c *OnDemandPricing) Count() int { _ = "STUB: not implemented"; return 0 }

func (c *OnDemandPricing) Save() error { _ = "STUB: not implemented"; return nil }

func (c *OnDemandPricing) Clear() error { _ = "STUB: not implemented"; return nil }

// fetchOnDemandPricing makes a bulk request to the pricing api to retrieve all instance type pricing if the instanceType is the empty string
//
//	or, if instanceType is specified, it can request a specific instance type pricing
func (c *OnDemandPricing) fetchOnDemandPricing(ctx context.Context, instanceType ec2types.InstanceType) (map[string]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StringMe takes an interface and returns a pointer to a string value
// If the underlying interface kind is not string or *string then nil is returned.
func (c *OnDemandPricing) StringMe(i interface{}) *string { _ = "STUB: not implemented"; return nil }

func (c *OnDemandPricing) getProductsInputFilters(instanceType ec2types.InstanceType) []pricingtypes.Filter {
	_ = "STUB: not implemented"
	return nil
}

// parseOndemandUnitPrice takes a priceList from the pricing API and parses its weirdness.
func (c *OnDemandPricing) parseOndemandUnitPrice(priceList string) (string, float64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}
