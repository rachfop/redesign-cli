package project

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//go:embed starter_templates/* starter_templates/**/* starter_templates/**/.*
var starterTemplates embed.FS

const basePath string = "starter_templates"

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#673ab7"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#673ab7"))
	// cursorStyle  = focusedStyle.Copy()
	noStyle = lipgloss.NewStyle()

	blurredButton = lipgloss.NewStyle().Foreground(lipgloss.Color("#673ab7")).Render("[ Submit ]")
	focusedButton = lipgloss.NewStyle().Foreground(lipgloss.Color("#673ab7")).Render("[ Submit ]")
)

type model struct {
	inputs     []textinput.Model
	focusIndex int
	err        error
}

func initialModel() model {
	var inputs []textinput.Model
	inputs = make([]textinput.Model, 5)

	var t textinput.Model
	for i := range inputs {
		t = textinput.New()
		// t.CursorStyle = cursorStyle
		// t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Project Name"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Use case: Text | Embed | Image | Audio"
			t.CharLimit = 64
		case 2:
			t.Placeholder = "Hugging Face Model Name: openchat/openchat-3.5-0106"
		case 3:
			t.Placeholder = "CUDA Version: 12.5, 12.4, 12.3, 12.2, etc."
		case 4:
			t.Placeholder = "Python Version: 3.11, 3.10, 3.9, 3.8"
		}

		inputs[i] = t
	}

	return model{
		inputs:     inputs,
		focusIndex: 0,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs)-1 {
				return m, tea.Quit
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs)-1 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) - 1
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n", *button)

	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok {
		createNewProject(
			m.inputs[0].Value(),
			m.inputs[1].Value(),
			m.inputs[2].Value(),
			m.inputs[3].Value(),
			m.inputs[4].Value(),
		)
	}
}

func createNewProject(projectName, modelType, modelName, cudaVersion, pythonVersion string) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	projectDir := filepath.Join(currentDir, projectName)
	err = os.Mkdir(projectDir, 0755)
	if err != nil {
		fmt.Println("Error creating project directory:", err)
		return
	}

	fmt.Println("Creating project in directory:", projectDir)

	createProjectStructure(projectDir, modelType, modelName, cudaVersion, pythonVersion)

	fmt.Println("Project created successfully in:", projectDir)
}

func copyFiles(files fs.FS, source string, dest string) error {
	return fs.WalkDir(files, source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == source {
			return nil
		}

		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}

		newPath := filepath.Join(dest, relPath)
		if d.IsDir() {
			if err := os.MkdirAll(newPath, os.ModePerm); err != nil {
				return err
			}
		} else {
			content, err := fs.ReadFile(files, path)
			if err != nil {
				return err
			}
			if err := os.WriteFile(newPath, content, 0o644); err != nil {
				return err
			}
		}
		return nil
	})
}

func createProjectStructure(projectDir, modelType, modelName, cudaVersion, pythonVersion string) {
	// Create README file
	readmePath := filepath.Join(projectDir, "README.md")
	readmeContent := fmt.Sprintf(`# %s

Model Type: %s
Model Name: %s
CUDA Version: %s
Python Version: %s
`, projectName, modelType, modelName, cudaVersion, pythonVersion)
	err := os.WriteFile(readmePath, []byte(readmeContent), 0644)
	if err != nil {
		fmt.Println("Error writing README file:", err)
	}

	// Create src directory
	srcDir := filepath.Join(projectDir, "src")
	err = os.Mkdir(srcDir, 0755)
	if err != nil {
		fmt.Println("Error creating src directory:", err)
	}

	// Create a simple main.py file in src directory
	mainPyPath := filepath.Join(srcDir, "main.py")
	mainPyContent := fmt.Sprintf(`# main.py

print("Hello from %s!")
`, projectName)
	err = os.WriteFile(mainPyPath, []byte(mainPyContent), 0644)
	if err != nil {
		fmt.Println("Error writing main.py file:", err)
	}

	// Create Dockerfile
	dockerfilePath := filepath.Join(projectDir, "Dockerfile")
	dockerfileContent := fmt.Sprintf(`FROM nvidia/cuda:%s-cudnn8-devel-ubuntu20.04

RUN apt-get update && apt-get install -y \
    python%s \
    python3-pip \
    && rm -rf /var/lib/apt/lists/*

COPY src/ /app/
WORKDIR /app

RUN pip install --no-cache-dir -r requirements.txt

CMD ["python3", "main.py"]
`, cudaVersion, pythonVersion)
	err = os.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644)
	if err != nil {
		fmt.Println("Error writing Dockerfile:", err)
	}

	// Create requirements.txt
	requirementsPath := filepath.Join(projectDir, "requirements.txt")
	requirementsContent := fmt.Sprintf(`# requirements.txt

# Add your project dependencies here
# e.g., transformers==4.10.3
`)

	err = os.WriteFile(requirementsPath, []byte(requirementsContent), 0644)
	if err != nil {
		fmt.Println("Error writing requirements.txt file:", err)
	}
}
