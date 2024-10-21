package main

import (
    "fmt"
    "strings"

    mqtt "github.com/eclipse/paho.mqtt.golang"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/bubbles/textinput"
    "github.com/charmbracelet/lipgloss"

)

type scrollBoxModel struct {
    lines       []string
    client      mqtt.Client
    ready       bool
    mqttChan    string
}

type addWordMsg string

type model struct {
    input       textinput.Model
    scrollBox   scrollBoxModel
    width       int
}

var (
    scrollBoxStyle = lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("63")).
        Padding(1,2).
        Margin(1).
        Width(80).
        Height(15).
        Align(lipgloss.Left)
    
    inputStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("205")).
        Background(lipgloss.Color("57")).
        Padding(0,1).
        Margin(0,1).
        Align(lipgloss.Center)
)

func newScrollBoxModel(mqttChan string) scrollBoxModel {
    opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
    client := mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        fmt.Println("MQTT connection failed:", token.Error())
    }

    return scrollBoxModel{
        lines:      []string{""},
        client:     client,
        ready:      true,
        mqttChan:   mqttChan,
    } 
}

func (s *scrollBoxModel) addWord(word string) tea.Cmd {
    if len(word) > 0 {
        if len(s.lines) == 0 || s.lines[len(s.lines)-1] == "" {
            s.lines[len(s.lines)-1] = word
        } else {
            s.lines = append(s.lines, word)
        }

        if s.ready {
            s.client.Publish(s.mqttChan, 0, false, word)
        }
    }

    return func() tea.Msg {
        return addWordMsg(word)
    }
}


func initialModel() model {
    ti := textinput.New()
    ti.Placeholder = "Type Here..."
    ti.Focus()
    ti.CharLimit = 156
    ti.Width = 20

    scrollBox := newScrollBoxModel("/4711/hive/speak")

    return model {
        input:      ti,
        scrollBox:  scrollBox,
        width: 80,
    }
}

func (m model) Init() tea.Cmd {
    return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        // currentLine := &m.scrollBox[len(m.scrollBox)-1]
        switch msg.String() {
        case "enter": 
            word := strings.TrimSpace(m.input.Value())
            if len(word) > 0 {
                return m, m.scrollBox.addWord(word)
                // *currentLine += " " + word
            }
            // m.scrollBox = append(m.scrollBox, "")
            m.input.SetValue("")
        case " ": 
            word := strings.TrimSpace(m.input.Value())
            if len(word) > 0 {
                return m, m.scrollBox.addWord(word)
                // if *currentLine == "" { *currentLine = word 
                // } else { *currentLine += " " + word }
            }
            m.input.SetValue("")
        case "ctrl+d", "ctrl+c":
            return m, tea.Quit
        }
    case addWordMsg:
        m.input.SetValue("")
    case tea.WindowSizeMsg:
        m.width = msg.Width
    }

    var cmd tea.Cmd
    m.input, cmd = m.input.Update(msg)
    return m, cmd
}

func (m model) View() string {
    // scrollView := strings.Join(m.scrollBox, "\n")
    scrollBoxWidth := m.width - 4
    scrollView := scrollBoxStyle.
        Width(scrollBoxWidth).
        Render(strings.Join(m.scrollBox.lines, "\n"))

    inputView := inputStyle.Render(m.input.View())

    return fmt.Sprintf( "%s\n\n%s", scrollView, inputView )
    // m.input.View(), )
}

func main() {
    p := tea.NewProgram(initialModel(), tea.WithAltScreen())
    if err := p.Start(); err != nil {
        fmt.Printf("Error: %v", err)
    }
}
