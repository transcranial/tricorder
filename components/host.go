package components

import (
        "github.com/shirou/gopsutil/host"
        "reflect"
        "strconv"
        "time"
        "fmt"
        "math"
)

func GetHostStats() (stats []string) {

        h, _ := host.HostInfo()

        hostinfo := reflect.ValueOf(*h)
        
        stats = append(stats, "")
        for i := 0; i < hostinfo.NumField(); i++ {
                label := hostinfo.Type().Field(i).Name
                item := " " + label + ":  "
                for k := 24 - len(label); k > 0; k-- {
                        item += " "
                }
                if label == "Procs" {
                        item += strconv.FormatUint(hostinfo.Field(i).Uint(), 10)
                } else if label == "Uptime" {
                        uptime := time.Since(time.Unix(int64(hostinfo.Field(i).Uint()), 0))
                        item += fmt.Sprintf("%.0f", math.Trunc(uptime.Hours() / 24)) + " days, "
                        item += fmt.Sprintf("%.0f", uptime.Hours() - 24*math.Trunc(uptime.Hours() / 24)) + " hours"
                } else {
                        item += hostinfo.Field(i).String()
                }
                stats = append(stats, item)
        }

        return
}