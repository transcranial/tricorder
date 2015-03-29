package main 

import (
        ui "github.com/gizak/termui"
        tm "github.com/nsf/termbox-go"

        "time"

        "github.com/transcranial/tricorder/components"
)

func main() {

        err := ui.Init()
        if err != nil {
                panic(err)
        }
        defer ui.Close()

        currentTime := time.Now()

        title := ui.NewPar("\nTRICORDER System Monitor\n")
        title.HasBorder = false
        title.Height = 3
        title.TextFgColor = ui.ColorWhite | ui.AttrBold | ui.AttrUnderline

        hostInfo := ui.NewList()
        hostInfo.Items = components.GetHostStats()
        hostInfo.ItemFgColor = ui.ColorWhite
        hostInfo.Border.Label = " Host Info "
        hostInfo.Border.LabelFgColor = ui.ColorWhite
        hostInfo.Border.LabelBgColor = ui.ColorMagenta
        hostInfo.Border.FgColor = ui.ColorMagenta
        hostInfo.Height = len(hostInfo.Items) + 2

        cpu := ui.NewBarChart()
        cpu.DataLabels, cpu.Data = components.GetCPUStats(currentTime)
        cpu.Border.Label = " CPU "
        cpu.Border.LabelFgColor = ui.ColorWhite
        cpu.Border.LabelBgColor = ui.ColorGreen
        cpu.Border.FgColor = ui.ColorGreen
        cpu.Height = len(hostInfo.Items) + 2
        cpu.BarColor = ui.ColorGreen
        cpu.NumColor = ui.ColorBlack

        commands := ui.NewPar("\n  [q] Quit")
        commands.Height = 4
        commands.TextFgColor = ui.ColorWhite
        commands.Border.Label = " Commands "
        commands.Border.LabelFgColor = ui.ColorWhite
        commands.Border.LabelBgColor = ui.ColorRed
        commands.Border.FgColor = ui.ColorRed

        colophon := ui.NewPar("\n  github.com/transcranial/tricorder")
        colophon.Height = 4
        colophon.TextFgColor = ui.ColorWhite
        colophon.Border.Label = " Colophon "
        colophon.Border.LabelFgColor = ui.ColorWhite
        colophon.Border.LabelBgColor = ui.ColorCyan
        colophon.Border.FgColor = ui.ColorCyan

        ui.Body.AddRows(
                ui.NewRow(
                        ui.NewCol(12, 0, title)),
                ui.NewRow(
                        ui.NewCol(6, 0, hostInfo),
                        ui.NewCol(6, 0, cpu)),
                ui.NewRow(
                        ui.NewCol(6, 0, commands),
                        ui.NewCol(6, 0, colophon)))

        ui.Body.Align()

        draw := func(t int) {
                cpu.DataLabels, cpu.Data = components.GetCPUStats(currentTime)
                ui.Render(ui.Body)
                currentTime = time.Now()
        }

        evt := make(chan tm.Event)
        go func() {
                for {
                        evt <- tm.PollEvent()
                }
        }()

        i := 0
        for {
                select {
                case e := <-evt:
                        if e.Type == tm.EventKey && e.Ch == 'q' {
                                return
                        }
                        if e.Type == tm.EventResize {
                                ui.Body.Width = ui.TermWidth()
                                ui.Body.Align()
                        }
                default:
                        draw(i)
                        i++
                        time.Sleep(time.Second / 2)
                }
        }
}