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
	"github.com/blang/semver/v4"
)

const (
	fallbackVersion = "5.20.0"
)

// EMR is a Service type for a custom service filter transform.
type EMR struct{}

// Filters implements the Service interface contract for EMR.
func (e EMR) Filters(version string) (Filters, error) {
	_ = "STUB: not implemented"
	return *new(Filters), nil
}

// getEMRInstanceTypes returns a list of instance types that emr supports.
func (e EMR) getEMRInstanceTypes(version semver.Version) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (EMR) isEMR_5_13_0_plus(instanceType string) bool { _ = "STUB: not implemented"; return false }

func (EMR) isOnlyEMR_5_20_0_plus(instanceType string) bool { _ = "STUB: not implemented"; return false }

func (EMR) isOnlyEMR_5_25_0_plus(instanceType string) bool { _ = "STUB: not implemented"; return false }

func (EMR) isOnlyEMR_5_33_0_plus(instanceType string) bool { _ = "STUB: not implemented"; return false }

func (EMR) getAllEMRInstanceTypes() []string { _ = "STUB: not implemented"; return nil }
