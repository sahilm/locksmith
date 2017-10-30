package main

import (
	"log"

	"time"

	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/sahilm/locksmith/git"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = false

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Fatal(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}

func keybindings(g *gocui.Gui) interface{} {
	return g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
}

func quit(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("main", 30, -1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		files, err := git.ListFiles("./.git", 1*time.Second)
		if err != nil {
			return err
		}
		for _, f := range files {
			fmt.Fprintln(v, f)
		}
	}
	return nil
}
