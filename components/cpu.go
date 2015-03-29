package components

import (
        "github.com/shirou/gopsutil/cpu"
        "time"
        "strconv"
)

func GetCPUStats(t time.Time) (cpuLabels []string, cpuPercents []int) {

        cpuP, _ := cpu.CPUPercent(time.Since(t), true)

        for i := 0; i < len(cpuP); i++ {
                cpuPercents[i] = int(100 * cpuP[i])
                cpuLabels[i] = strconv.Itoa(i)
        }

        if len(cpuP) == 0 {
                cpuPercents = append(cpuPercents, 0)
                cpuLabels = append(cpuLabels, "")
        }

        return
}