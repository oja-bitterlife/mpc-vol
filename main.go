package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	volume int
}

func getVolume() int {
	out, _ := exec.Command("mpc", "volume").Output()
	s := strings.Fields(string(out))
	if len(s) > 1 {
		v, _ := strconv.Atoi(strings.TrimSuffix(s[1], "%"))
		return v
	}
	return 0
}

func setVolume(diff int) int {
	arg := fmt.Sprintf("%+d", diff)
	if diff > 0 {
		arg = "+" + strconv.Itoa(diff)
	}
	exec.Command("mpc", "volume", arg).Run()
	return getVolume()
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "enter", "esc":
			return m, tea.Quit
		case "up":
			m.volume = setVolume(5)
		case "down":
			m.volume = setVolume(-5)
		case "right":
			m.volume = setVolume(1)
		case "left":
			m.volume = setVolume(-1)
		}
	}
	return m, nil
}

func (m model) View() string {
	width := 30
	filled := min(max(m.volume*width/100, 0), width)

	bar := strings.Repeat("■", filled) + strings.Repeat(" ", width-filled)

	return fmt.Sprintf("mpc volume [%-30s] %3d%%", bar, m.volume)
}

func main() {
	// 1. プログラムの起動
	p := tea.NewProgram(model{volume: getVolume()})

	// 2. 実行（ここでBubble Teaが画面を占有・更新する）
	m, err := p.Run()
	if err != nil {
		os.Exit(1)
	}

	// 3. 終了後、最後に確定した状態を1回だけ表示する
	// 型アサーションで最終的なModelを取り出し、Viewを呼ぶだけ
	if finalModel, ok := m.(model); ok {
		fmt.Printf("\r%s\n", finalModel.View())
	}
}
