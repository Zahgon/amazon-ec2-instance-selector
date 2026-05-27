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

package instancetypes

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/patrickmn/go-cache"
)

var CacheFileName = "ec2-instance-types.json"

// Details hold all the information on an ec2 instance type.
type Details struct {
	ec2types.InstanceTypeInfo
	OndemandPricePerHour *float64
	SpotPrice            *float64
}

type Provider struct {
	Region          string
	DirectoryPath   string
	FullRefreshTTL  time.Duration
	lastFullRefresh *time.Time
	ec2Client       ec2.DescribeInstanceTypesAPIClient
	cache           *cache.Cache
	logger          *log.Logger
}

// NewProvider creates a new Instance Types provider used to fetch Instance Type information from EC2.
func NewProvider(region string, ec2Client ec2.DescribeInstanceTypesAPIClient) *Provider {
	_ = "STUB: not implemented"
	return nil
}

// NewProvider creates a new Instance Types provider used to fetch Instance Type information from EC2 and optionally cache.
func LoadFromOrNew(directoryPath string, region string, ttl time.Duration, ec2Client ec2.DescribeInstanceTypesAPIClient) (*Provider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadFrom(ttl time.Duration, region string, expandedDirPath string) (*cache.Cache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getCacheFilePath(region string, expandedDirPath string) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *Provider) SetLogger(logger *log.Logger) { _ = "STUB: not implemented"; return }

func (p *Provider) Get(ctx context.Context, instanceTypes []ec2types.InstanceType) ([]*Details, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// need to reassign, so we're not sharing the loop iterators memory space

// if we were able to retrieve all from cache, return here, else continue to do a remote lookup

func (p *Provider) isFullRefreshNeeded() bool { _ = "STUB: not implemented"; return false }

func (p *Provider) Save() error { _ = "STUB: not implemented"; return nil }

func (p *Provider) Clear() error { _ = "STUB: not implemented"; return nil }

func (p *Provider) CacheCount() int { _ = "STUB: not implemented"; return 0 }
