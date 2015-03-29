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

/*
        g := ui.NewGauge()
        g.Percent = 50
        g.Width = 50
        g.Height = 3
        g.Y = 11
        g.Border.Label = "Gauge"
        g.BarColor = ui.ColorRed
        g.Border.FgColor = ui.ColorWhite
        g.Border.LabelFgColor = ui.ColorCyan

        spark := ui.Sparkline{}
        spark.Height = 1
        spark.Title = "srv 0:"
        spdata := []int{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}
        spark.Data = spdata
        spark.LineColor = ui.ColorCyan
        spark.TitleColor = ui.ColorWhite

        spark1 := ui.Sparkline{}
        spark1.Height = 1
        spark1.Title = "srv 1:"
        spark1.Data = spdata
        spark1.TitleColor = ui.ColorWhite
        spark1.LineColor = ui.ColorRed

        sp := ui.NewSparklines(spark, spark1)
        sp.Width = 25
        sp.Height = 7
        sp.Border.Label = "Sparkline"
        sp.Y = 4
        sp.X = 25

        sinps := (func() []float64 {
                n := 220
                ps := make([]float64, n)
                for i := range ps {
                        ps[i] = 1 + math.Sin(float64(i)/5)
                }
                return ps
        })()

        lc := ui.NewLineChart()
        lc.Border.Label = "dot-mode Line Chart"
        lc.Data = sinps
        lc.Width = 50
        lc.Height = 11
        lc.X = 0
        lc.Y = 14
        lc.AxesColor = ui.ColorWhite
        lc.LineColor = ui.ColorRed | ui.AttrBold
        lc.Mode = "dot"

        bc := ui.NewBarChart()
        bcdata := []int{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
        bclabels := []string{"S0", "S1", "S2", "S3", "S4", "S5"}
        bc.Border.Label = "Bar Chart"
        bc.Width = 26
        bc.Height = 10
        bc.X = 51
        bc.Y = 0
        bc.DataLabels = bclabels
        bc.BarColor = ui.ColorGreen
        bc.NumColor = ui.ColorBlack

        lc1 := ui.NewLineChart()
        lc1.Border.Label = "braille-mode Line Chart"
        lc1.Data = sinps
        lc1.Width = 26
        lc1.Height = 11
        lc1.X = 51
        lc1.Y = 14
        lc1.AxesColor = ui.ColorWhite
        lc1.LineColor = ui.ColorYellow | ui.AttrBold
*/

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