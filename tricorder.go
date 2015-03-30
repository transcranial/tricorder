package main 

import (
        ui "github.com/gizak/termui"
        tm "github.com/nsf/termbox-go"

        "time"
        "math"

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
        cpu.NumColor = ui.ColorWhite
        cpu.TextColor = ui.ColorGreen

        memory := ui.NewList()
        memory.Items = components.GetMemStats()
        memory.ItemFgColor = ui.ColorWhite
        memory.Border.Label = " Memory "
        memory.Border.LabelFgColor = ui.ColorWhite
        memory.Border.LabelBgColor = ui.ColorYellow
        memory.Border.FgColor = ui.ColorYellow
        memory.Height = len(memory.Items) + 2

        memoryLC := make([]*ui.LineChart, 2)
        memoryLC[0] = ui.NewLineChart()
        memoryLC[0].Border.Label = " virtual "
        memoryLC[0].Height = memory.Height / 2
        memoryLC[0].Mode = "dot"
        memoryLC[0].AxesColor = ui.ColorWhite
        memoryLC[0].LineColor = ui.ColorYellow
        memoryLC[0].Border.LabelFgColor = ui.ColorYellow
        memoryLC[0].Border.FgColor = ui.ColorYellow
        memoryLC[0].AxesColor = ui.ColorYellow
        memoryLC[1] = ui.NewLineChart()
        memoryLC[1].Border.Label = " swap "
        memoryLC[1].Height = memory.Height / 2
        memoryLC[1].Mode = "dot"
        memoryLC[1].AxesColor = ui.ColorWhite
        memoryLC[1].LineColor = ui.ColorYellow
        memoryLC[1].Border.LabelFgColor = ui.ColorYellow
        memoryLC[1].Border.FgColor = ui.ColorYellow
        memoryLC[1].AxesColor = ui.ColorYellow

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
                        ui.NewCol(3, 0, memory),
                        ui.NewCol(3, 0, memoryLC[0], memoryLC[1])),
                ui.NewRow(
                        ui.NewCol(6, 0, commands),
                        ui.NewCol(6, 0, colophon)))

        ui.Body.Align()

        cpu.BarWidth = int(math.Floor((float64(cpu.Width) / float64(len(cpu.Data))) * 0.5))
        cpu.BarGap = int(math.Floor((float64(cpu.Width) / float64(len(cpu.Data))) * 0.4))
        memVirtData := make([]float64, memoryLC[0].Width - 10)
        memSwapData := make([]float64, memoryLC[1].Width - 10)

        updateMemData := func() {
                memVirtData = memVirtData[1:len(memVirtData)]
                memSwapData = memSwapData[1:len(memSwapData)]
                memory.Items = components.GetMemStats()
                vmd, smd := components.GetNextMemData()
                memVirtData = append(memVirtData, vmd)
                memSwapData = append(memSwapData, smd)
                memoryLC[0].Data = memVirtData
                memoryLC[1].Data = memSwapData
        }

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
                        if i % 5 == 0 {
                                // update memory data every 5 seconds
                                updateMemData()
                        }
                        // redraw every second
                        draw(i)
                        i++
                        time.Sleep(time.Second)
                }
        }
}