package ui

import (
	//lib "charmlab/lib"
	"fmt"
	"os"

	//kv store

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// func whenEnterFolder(){}
// func whenSearchFolder(){}
// func whenSearchFromBaseOfStack

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list           list.Model
	bucketSelected bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

// Update func global variables.
var selectionFromUpdateLoop string = ""

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "q":
			return m, tea.Quit
		case "y":
			selectionFromUpdateLoop = m.list.SelectedItem().FilterValue()
			return m, tea.Quit
		case "enter", " ":
			if m.list.FilterState().String() != "filtering" {
				selectionFromUpdateLoop = m.list.SelectedItem().FilterValue()
				return m, tea.Quit
			}
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	// TODO put this on a coroutine and have it update once every 10 seconds.
	//m.updateCacheLog("cache.log")

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) updateCacheLog(logFilename string) {
	os.Remove(logFilename)
	f, err := tea.LogToFile(logFilename, "")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	var cachedData string = ""

	for i := 1; i < len(m.list.Items()); i++ {
		cachedData += m.list.Items()[i].FilterValue() + "\n"
	}

	f.WriteString(cachedData)

}

func Home(title string, filteredList []string) string {

	items := []list.Item{}

	for i := 1; i < len(filteredList); i++ {
		items = append(items, item{title: filteredList[i], desc: ""})
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = title
	m.list.Styles.Title = lipgloss.NewStyle().
		Background(lipgloss.Color("#C678DD")).
		Foreground(lipgloss.Color("#282828")).
		Padding(0, 1)

	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return selectionFromUpdateLoop
}

/*
TODO:
- display filesize and date for files in desc.
- settings menu w/ togglable options.
- file preview support.
	- built-in support for .orc files.
- open in browser.
*/
