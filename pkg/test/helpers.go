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

package test

import (
	"testing"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) { _ = "STUB: not implemented"; return }

// Nok fails the test if an err is nil.
func Nok(tb testing.TB, err error) { _ = "STUB: not implemented"; return }

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) { _ = "STUB: not implemented"; return }
