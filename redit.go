package main

import (
	"bufio"
	_ "embed"
	"os"

	g "github.com/AllenDang/giu"
)

//go:embed font/WenQuanDengKuanZhengHei-1.ttf
var fontContent []byte

//go:embed font/FiraCode-Regular.ttf
var defaultFontContent []byte

var inputText string
var bigFont *g.FontInfo
var wnd *g.MasterWindow
var fileName = "untitled.txt"

func onSave() {
	if fileName != "untitled.txt" {
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		write := bufio.NewWriter(file)
		if _, err := write.WriteString(inputText); err != nil {
			panic(err)
		}
		if err := write.Flush(); err != nil {
			panic(err)
		}
	}
	wnd.Close()
}

func loop() {
	w, h := wnd.GetSize()
	g.SingleWindow().Layout(
		g.Column(
			g.Align(g.AlignCenter).To(
				g.Style().SetFont(bigFont).To(g.InputTextMultiline(&inputText).Size(float32(w-12), float32(h-60))),
			),
			g.Align(g.AlignRight).To(
				g.Style().SetFont(bigFont).To(g.Button("Save").Size(50, 32).OnClick(onSave)),
			),
		),
	)
}

func main() {
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	g.SetDefaultFontFromBytes(defaultFontContent, 20)
	bigFont = g.AddFontFromBytes("WenQuanDengKuanZhengHei", fontContent, 20)
	wnd = g.NewMasterWindow(fileName, 480, 260, 0)
	wnd.Run(loop)
}
