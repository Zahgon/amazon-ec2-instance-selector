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
package selector

import (
	"context"
	"regexp"
)

const (
	// AggregateLowPercentile is the default lower percentile for resource ranges on similar instance type comparisons.
	AggregateLowPercentile = 0.9
	// AggregateHighPercentile is the default upper percentile for resource ranges on similar instance type comparisons.
	AggregateHighPercentile = 1.2
)

var baseAllowedInstanceTypesRE = regexp.MustCompile(`^[cmr][3-9][agi]?\..*$|^t[2-9][gi]?\..*$`)

// FiltersTransform can be implemented to provide custom transforms.
type FiltersTransform interface {
	Transform(context.Context, Filters) (Filters, error)
}

// TransformFn is the func type definition for a FiltersTransform.
type TransformFn func(context.Context, Filters) (Filters, error)

// Transform implements FiltersTransform interface on TransformFn
// This allows any TransformFn to be passed into funcs accepting FiltersTransform interface.
func (fn TransformFn) Transform(ctx context.Context, filters Filters) (Filters, error) {
	_ = "STUB: not implemented"
	return *

	// TransformBaseInstanceType transforms lower level filters based on the instanceTypeBase specs.
	new(Filters), nil
}

func (itf Selector) TransformBaseInstanceType(ctx context.Context, filters Filters) (Filters, error) {
	_ = "STUB: not implemented"
	return *new(Filters), nil
}

// TransformFlexible transforms lower level filters based on a set of opinions.
func (itf Selector) TransformFlexible(ctx context.Context, filters Filters) (Filters, error) {
	_ = "STUB: not implemented"
	return *new(Filters), nil
}

// TransformForService transforms lower level filters based on the service.
func (itf Selector) TransformForService(ctx context.Context, filters Filters) (Filters, error) {
	_ = "STUB: not implemented"
	return *new(Filters), nil
}
