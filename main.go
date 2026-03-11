package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var Version = "dev"

type model struct {
	volume int
}

func getInitialVolume() int {
	out, _ := exec.Command("mpc", "volume").Output()

	// volume:100% という形式で出力されるので、そこから数値だけを抜き取る
	result := strings.TrimSpace(string(out))
	value := strings.TrimSuffix(strings.TrimPrefix(string(result), "volume:"), "%")
	if v, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
		return v
	} else {
		// エラー終了
		println("Failed to get initial volume: ", result)
		os.Exit(1)
	}

	return 0
}

func setVolume(m model, diff int) int {
	// 音量を0-100の範囲に収める
	m.volume = min(max(m.volume+diff, 0), 100)
	exec.Command("mpc", "volume", fmt.Sprintf("%d", m.volume)).Run()

	return m.volume
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "enter", "esc":
			return m, tea.Quit
		case "up":
			m.volume = setVolume(m, 5)
		case "down":
			m.volume = setVolume(m, -5)
		case "right":
			m.volume = setVolume(m, 1)
		case "left":
			m.volume = setVolume(m, -1)
		}
	}

	return m, nil
}

func (m model) View() string {
	width := 30
	filled := min(max(m.volume*width/100, 0), width)

	bar := strings.Repeat("■", filled) + strings.Repeat(" ", width-filled)

	// return fmt.Sprintf("mpc volume [%-30s] %3d%%", bar, m.volume)
	return fmt.Sprintf("mpc-vol %s\n[%-30s] %3d%%", Version, bar, m.volume)
}

func main() {
	// 1. プログラムの起動
	p := tea.NewProgram(model{volume: getInitialVolume()})

	// 2. 実行（ここでBubble Teaが画面を占有・更新する）
	m, err := p.Run()
	if err != nil {
		os.Exit(1)
	}

	// 3. 終了後、最後に確定した状態を1回だけ表示する
	// 型アサーションで最終的なModelを取り出し、Viewを呼ぶだけ
	if finalModel, ok := m.(model); ok {
		fmt.Printf("\033[2A\r%s\n", finalModel.View())
	}
}
