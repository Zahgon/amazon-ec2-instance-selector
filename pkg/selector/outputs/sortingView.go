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
	"io"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/aws/amazon-ec2-instance-selector/v3/pkg/instancetypes"
)

const (
	// formatting.
	sortDirectionPadding = 2
	sortingTitlePadding  = 3
	sortingFooterPadding = 2

	// controls.
	sortingListControls = "Controls: ↑/↓ - up/down • enter - select filter • tab - toggle direction • esc - return to table • q - quit"
	sortingTextControls = "Controls: ↑/↓ - up/down • tab - toggle direction • enter - enter json path"

	// sort direction text.
	ascendingText  = "ASCENDING"
	descendingText = "DESCENDING"
)

// sortingModel holds the state for the sorting view.
type sortingModel struct {
	// list which holds the available shorting shorthands
	shorthandList list.Model

	// text input for json paths
	sortTextInput textinput.Model

	instanceTypes []*instancetypes.Details

	isDescending bool
}

// format styles.
var (
	// list.
	listTitleStyle    = lipgloss.NewStyle().Bold(true).Underline(true)
	listItemStyle     = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))

	// text.
	descendingStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#0096FF"))
	ascendingStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#DAF7A6"))
	sortDirectionStyle = lipgloss.NewStyle().Bold(true).Underline(true).PaddingLeft(2)
)

// implement Item interface for list.
type item string

func (i item) FilterValue() string { _ = "STUB: not implemented"; return "" }
func (i item) Title() string       { _ = "STUB: not implemented"; return "" }
func (i item) Description() string {
	_ = "STUB: not implemented"

	// implement ItemDelegate for list.
	return ""
}

type itemDelegate struct{}

func (d itemDelegate) Height() int  { _ = "STUB: not implemented"; return 0 }
func (d itemDelegate) Spacing() int { _ = "STUB: not implemented"; return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	_ = "STUB: not implemented"
	return *new(tea.Cmd)
}
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	_ = "STUB: not implemented"
	return
}

// initSortingModel initializes and returns a new tableModel based on the given
// instance type details.
func initSortingModel(instanceTypes []*instancetypes.Details) *sortingModel {
	_ = "STUB: not implemented"
	return nil
}

// createListKeyMap creates a KeyMap with the controls for the shorthand list.
func createListKeyMap() list.KeyMap { _ = "STUB: not implemented"; return *new(list.KeyMap) }

// createListItems creates a list item for shorthand sorting flag.
func createListItems() *[]list.Item { _ = "STUB: not implemented"; return nil }

// resizeSortingView will change the dimensions of the sorting view
// in order to accommodate the new window dimensions represented by
// the given tea.WindowSizeMsg.
func (m sortingModel) resizeView(msg tea.WindowSizeMsg) sortingModel {
	_ = "STUB: not implemented"
	return *new(sortingModel)
}

// ensure that text input is right below last option

// ensure cursor of list is still hidden after resize

// update updates the state of the sortingModel.
func (m sortingModel) update(msg tea.Msg) (sortingModel, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(sortingModel), *new(tea.Cmd)
}

// focus text input and hide cursor in shorthand list

// go back to list from text input

// view returns a string representing the sorting view.
func (m sortingModel) view() string { _ = "STUB: not implemented"; return "" }

// draw sort direction

// draw list

// draw text input

// draw controls
