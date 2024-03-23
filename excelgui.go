package main

import (
	"fmt"
  "log"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello, World!")
	questionButton := widget.NewButton("Ask a Question", func() {
		f, err := excelize.OpenFile("D:/ymk.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// 读取指定工作表的数据
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	var brr, grr, hrr, irr []string

	for _, row := range rows {
		brr = append(brr, row[1]) // B列数据
		grr = append(grr, row[6]) // G列数据
		hrr = append(hrr, row[7]) // H列数据
		irr = append(irr, row[8]) // I列数据
	}

	// 输出读取的数据
	fmt.Println("B 列数据:", brr)
	fmt.Println("G 列数据:", grr)
	fmt.Println("H 列数据:", hrr)
	fmt.Println("I 列数据:", irr)
	})
  content := container.NewVBox(
		questionButton,
	)
	w.SetContent(content)
	w.ShowAndRun()
}

