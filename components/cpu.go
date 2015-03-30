package components

import (
        "github.com/shirou/gopsutil/cpu"
        "time"
        "strconv"
        "math"
)

func GetCPUStats(t time.Time) (cpuLabels []string, cpuPercents []int) {

        cpuP, _ := cpu.CPUPercent(time.Since(t), true)

        for i := 0; i < len(cpuP); i++ {
                perc := int(math.Ceil(cpuP[i]))
                cpuPercents = append(cpuPercents, perc)
                cpuLabels = append(cpuLabels, " # " + strconv.Itoa(i+1))
        }

        if len(cpuP) == 0 {
                cpuPercents = append(cpuPercents, 0)
                cpuLabels = append(cpuLabels, "")
        }

        return
}