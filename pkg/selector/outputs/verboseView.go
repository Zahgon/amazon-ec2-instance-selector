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

package outputs

import (
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	// verbose view formatting.
	outlinePadding = 8

	// controls.
	verboseControls = "Controls: ↑/↓ - up/down • esc - return to table • q - quit"
)

// verboseModel represents the current state of the verbose view.
type verboseModel struct {
	// model for verbose output viewport
	viewport viewport.Model

	// the instance which the verbose output is focused on
	focusedInstanceName ec2types.InstanceType
}

// styling for viewport.
var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
)

// initVerboseModel initializes and returns a new verboseModel based on the given
// instance type details.
func initVerboseModel() *verboseModel { _ = "STUB: not implemented"; return nil }

// resizeView will change the dimensions of the verbose viewport in order to accommodate
// the new window dimensions represented by the given tea.WindowSizeMsg.
func (m verboseModel) resizeView(msg tea.WindowSizeMsg) verboseModel {
	_ = "STUB: not implemented"
	// handle width changes
	return *new(verboseModel)
}

// handle height changes

// height too short to fit viewport

// update updates the state of the verboseModel.
func (m verboseModel) update(msg tea.Msg) (verboseModel, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(verboseModel), *new(tea.Cmd)
}

func (m verboseModel) view() string { _ = "STUB: not implemented"; return "" }

// format header for viewport

// format footer for viewport

// controls
