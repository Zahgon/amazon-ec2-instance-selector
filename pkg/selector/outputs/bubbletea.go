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
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/aws/amazon-ec2-instance-selector/v3/pkg/instancetypes"
)

const (
	// can't get terminal dimensions on startup, so use this.
	initialDimensionVal = 30

	instanceTypeKey = "instance type"
	selectedKey     = "selected"
)

const (
	// table states.
	stateTable   = "table"
	stateVerbose = "verbose"
	stateSorting = "sorting"
)

var controlsStyle = lipgloss.NewStyle().Faint(true)

// BubbleTeaModel is used to hold the state of the bubble tea TUI.
type BubbleTeaModel struct {
	// holds the output currentState of the model
	currentState string

	// the model for the table view
	tableModel tableModel

	// holds state for the verbose view
	verboseModel verboseModel

	// holds the state for the sorting view
	sortingModel sortingModel
}

// NewBubbleTeaModel initializes a new bubble tea Model which represents
// a stylized table to display instance types.
func NewBubbleTeaModel(instanceTypes []*instancetypes.Details) BubbleTeaModel {
	_ = "STUB: not implemented"
	return *new(BubbleTeaModel)
}

// Init is used by bubble tea to initialize a bubble tea table.
func (m BubbleTeaModel) Init() tea.Cmd {
	_ = "STUB: not implemented"

	// Update is used by bubble tea to update the state of the bubble
	// tea model based on user input.
	return *new(tea.Cmd)
}

func (m BubbleTeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// don't listen for input if currently typing into text field

// see if we should sort and switch states to table

// check for quit or change in state

// switch from table state to verbose state

// get focused instance type

// set content of view

// move viewport to top of printout

// switch from table state to verbose state

// switch from table view to sorting view

// sort and switch states to table

// switch from sorting state or verbose state to table state

// This is needed to handle a bug with bubble tea
// where resizing causes misprints (https://github.com/Evertras/bubble-table/issues/121)
//nolint:staticcheck

// handle screen resizing

// update currently active state

// View is used by bubble tea to render the bubble tea model.
func (m BubbleTeaModel) View() string { _ = "STUB: not implemented"; return "" }
