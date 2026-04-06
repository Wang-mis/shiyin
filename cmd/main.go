package main

import (
	"fmt"
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"shiyin/internal/data"
	"shiyin/internal/ui"
)

const usage = `shiyin — 终端诗词阅读器

用法:
  shiyin [集合] [-s]

集合 (可选):
  tang      唐诗三百首
  ci        宋词三百首
  all       全部
  fav       收藏夹

选项:
  -s        乱序 (随机打乱诗词顺序)

示例:
  shiyin              启动集合选择器
  shiyin tang         直接打开唐诗三百首
  shiyin fav          直接打开收藏夹
  shiyin all -s       全部诗词，乱序
`

func main() {
	collection := ""
	shuffle := false

	for _, arg := range os.Args[1:] {
		switch arg {
		case "-s", "--shuffle":
			shuffle = true
		case "-h", "--help":
			fmt.Print(usage)
			os.Exit(0)
		case "tang", "tang300":
			collection = "tang300"
		case "ci", "ci300":
			collection = "ci300"
		case "all":
			collection = "all"
		case "fav":
			collection = "fav"
		default:
			fmt.Fprintf(os.Stderr, "未知参数: %q\n\n%s", arg, usage)
			os.Exit(1)
		}
	}

	var model tea.Model
	if collection != "" {
		name, poems, err := loadCollection(collection)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if shuffle {
			rand.Shuffle(len(poems), func(i, j int) { poems[i], poems[j] = poems[j], poems[i] })
		}
		model = ui.NewAppModelWithCollection(collection, name, poems)
	} else {
		model = ui.NewAppModel()
	}

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func loadCollection(key string) (name string, poems []data.Poem, err error) {
	switch key {
	case "tang300":
		name = "唐诗三百首"
	case "ci300":
		name = "宋词三百首"
	case "all":
		name = "全部"
	case "fav":
		name = "收藏夹"
		poems, err = data.LoadFavorites()
		return
	}
	poems, err = data.Load(key)
	return
}
