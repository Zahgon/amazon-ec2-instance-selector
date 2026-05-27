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
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"

	"github.com/aws/amazon-ec2-instance-selector/v3/pkg/instancetypes"
)

const (
	// table formatting.
	headerAndFooterPadding = 8
	headerPadding          = 2

	// controls.
	tableControls = "Controls: ↑/↓ - up/down • ←/→  - left/right • shift + ←/→ - pg up/down • e - expand • f - filter • t - trim toggle • space - select • s - sort • q - quit"
	ellipses      = "..."

	jsonPathError = "INVALID JSON PATH"
)

type tableModel struct {
	// the model for the table output
	table table.Model

	// width and rows per page are inaccessible through
	// bubble tea implementation, so expose them here
	tableWidth       int
	tableRowsPerPage int

	// the model for the filtering text input
	filterTextInput textinput.Model

	// shows whether the rows are currently trimmed or not
	isTrimmed bool

	// the rows that existed on the table's creation
	originalRows []table.Row

	canSelectRows bool
}

var customBorder = table.Border{
	Top:    "─",
	Left:   "│",
	Right:  "│",
	Bottom: "─",

	TopRight:    "╮",
	TopLeft:     "╭",
	BottomRight: "╯",
	BottomLeft:  "╰",

	TopJunction:    "┬",
	LeftJunction:   "├",
	RightJunction:  "┤",
	BottomJunction: "┴",
	InnerJunction:  "┼",

	InnerDivider: "│",
}

// initTableModel initializes and returns a new tableModel based on the given
// instance type details.
func initTableModel(instanceTypes []*instancetypes.Details) *tableModel {
	_ = "STUB: not implemented"
	return nil
}

// createFilterTextInput creates and styles a text input for filtering.
func createFilterTextInput() textinput.Model {
	_ = "STUB: not implemented"
	return *new(textinput.Model)
}

// createRows creates a row for each instance type in the passed in list.
func createRows(columnsData []*wideColumnsData, instanceTypes []*instancetypes.Details) *[]table.Row {
	_ = "STUB: not implemented"
	return nil

	// create a row for each instance type
}

// create a new row by iterating through the column data
// struct and using struct tags as column keys

// add instance type as metaData

// add selected flag as metadata

// maxColWidth finds the maximum width element in the given column.
func maxColWidth(columnsData []*wideColumnsData, columnHeader string) int {
	_ = "STUB: not implemented"
	// default max width is the width of the header itself with padding
	return 0
}

// get data at given column

// see if the width of the current column element exceeds
// the previous max width

// createColumns creates columns based on the tags in the wideColumnsData
// struct.
func createColumns(columnsData []*wideColumnsData) *[]table.Column {
	_ = "STUB: not implemented"
	return nil

	// iterate through wideColumnsData struct and create a new column for each field tag
}

// createTableKeyMap creates a KeyMap with the controls for the table.
func createTableKeyMap() *table.KeyMap { _ = "STUB: not implemented"; return nil }

// createTable creates an intractable table which contains information about all of
// the given instance types.
func createTable(instanceTypes []*instancetypes.Details) table.Model {
	_ = "STUB: not implemented"
	// calculate and fetch all column data from instance types
	return *new(table.Model)
}

// resizeView will change the dimensions of the table in order to accommodate
// the new window dimensions represented by the given tea.WindowSizeMsg.
func (m tableModel) resizeView(msg tea.WindowSizeMsg) tableModel {
	_ = "STUB: not implemented"
	// handle width changes
	return *new(tableModel)
}

// handle height changes

// height too short to fit footer and header
// so only display 1 row

// updateFooter updates the page and controls string in the table footer.
func (m tableModel) updateFooter() tableModel { _ = "STUB: not implemented"; return *new(tableModel) }

// prevent controls text from wrapping to avoid table misprints

// update updates the state of the tableModel.
func (m tableModel) update(msg tea.Msg) (tableModel, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tableModel), *new(tea.Cmd)
}

// update filtering input field

// exit filter input and update controls string

// listen for specific inputs

// focus filter input field

// handle trimming to selected rows

// undo trim

// trim

// custom toggling of selected rows because bubble tea implementation
// breaks trimming

// flip selected flag

// update selected row with new selected state. Must iterate through
// original rows since the cursor index in the bubble tea table
// takes the filter into account and therefore returns an incorrect index

// update footer

// view returns a string representing the table view.
func (m tableModel) view() string { _ = "STUB: not implemented"; return "" }

// sortTable sorts the table based on the sorting direction and sorting filter.
func (m tableModel) sortTable(sortFilter string, sortDirection string) (tableModel, error) {
	_ = "STUB: not implemented"
	return *new(tableModel), nil
}

// sort instance types

// get sorted rows from sorted instance types

// apply truncation if needed

// getInstanceTypeFromRows goes through the rows of the table model and returns both a list of instance
// types and a mapping of instances to rows.
func (m tableModel) getInstanceTypeFromRows() ([]*instancetypes.Details, map[string]table.Row) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get current rows

// if current table is trimmed, get the stored untrimmed rows

// since table isn't trimmed, we should get the unfiltered rows
// so that our rows have the most updated selected flags

// getUnfilteredRows gets the rows in the given table model without any filtering applied.
func (m tableModel) getUnfilteredRows() []table.Row { _ = "STUB: not implemented"; return nil }

// trim will trim the table to only the selected rows.
func (m tableModel) trim() tableModel {
	_ = "STUB: not implemented"
	// store current state of rows before trimming
	return *new(tableModel)
}

// prevent rows from being selected until trim is
// undone

// untrim will return the table to the original rows.
func (m tableModel) untrim() tableModel { _ = "STUB: not implemented"; return *new(tableModel) }

// allow rows to be selected again
