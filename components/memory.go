package components

import (
        "github.com/shirou/gopsutil/mem"
        "reflect"
        "fmt"
)

func GetMemStats() (stats []string) {

        vm, _ := mem.VirtualMemory()
        sm, _ := mem.SwapMemory()

        vmInfo := reflect.ValueOf(*vm)
        smInfo := reflect.ValueOf(*sm)
        
        stats = append(stats, "")
        stats = append(stats, " Virtual")
        for i := 0; i < vmInfo.NumField(); i++ {
                label := vmInfo.Type().Field(i).Name
                item := "   " + label + ":  "
                for k := 14 - len(label); k > 0; k-- {
                        item += " "
                }
                if label == "UsedPercent" {
                        item += fmt.Sprintf("%.1f", vmInfo.Field(i).Float()) + "%"
                } else if label == "Wired" || label == "Shared" {
                        continue
                } else {
                        item += fmt.Sprintf("%.3f GB", float64(vmInfo.Field(i).Uint()) / float64(1000000000))
                }
                stats = append(stats, item)
        }
        
        stats = append(stats, " Swap")
        for i := 0; i < smInfo.NumField(); i++ {
                label := smInfo.Type().Field(i).Name
                item := "   " + label + ":  "
                for k := 14 - len(label); k > 0; k-- {
                        item += " "
                }
                if label == "UsedPercent" {
                        item += fmt.Sprintf("%.1f", smInfo.Field(i).Float()) + "%"
                } else if label == "Sin" || label == "Sout" {
                        continue
                } else {
                        item += fmt.Sprintf("%.3f GB", float64(smInfo.Field(i).Uint()) / float64(1000000000))
                }
                stats = append(stats, item)
        }

        return
}

func GetNextMemData() (vmd, smd float64) {

        vm, _ := mem.VirtualMemory()
        sm, _ := mem.SwapMemory()
        vmd = vm.UsedPercent
        smd = sm.UsedPercent

        return
}