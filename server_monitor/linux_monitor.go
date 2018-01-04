package server_monitor

import (
	"github.com/kevin-zx/go-util/linuxCommandUtil"
	"strings"
	"strconv"
	"fmt"
)

type LinuxMonitor struct {
	linuxCommandUtil.LinuxServer
}

type Disk struct {
	Filesystem string
	Size       string
	Used       string
	Avail      string
	Use        float64
	MountedOn  string
}

func (d *Disk) ToString() string  {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%f%%\t%s",d.Filesystem,d.Size,d.Used,d.Avail,d.Use,d.MountedOn)
}


type es_health struct {
	epoch,timestamp,cluster,status string
	total_node,data_node,shards,pri,relo,init,unassign,pending_tasks int
	max_task_wait_time string
	active_shards_percent string
}

func (lm *LinuxMonitor) MonitorDisk () ([]Disk,error)  {
	df_result,err := lm.ExecCommand("df -h")
	if err != nil{
		return nil,err
	}
	disks := []Disk{}
	df_result = strings.Replace(df_result,"d on","dOn",1)
	df_result = strings.Replace(df_result,"\n"," ",-1)
	data := []string{}
	parts := strings.Split(df_result ," ")
	for _,p := range parts {
		if len(p) > 0 {
			data = append(data, p)
		}
	}
	if len(data)!=0 && len(data)% 6 != 0{
		println(df_result)
		return nil,fmt.Errorf("%s \n can't decode", df_result)
	}
	all_group_count :=len(data)/6
	for i:=0;i<all_group_count ;i++  {
		if data[6*i] == "Filesystem" || data[6*i] =="文件系统"{
			continue
		}
		use,err := strconv.ParseFloat(strings.Replace(data[6*i+4],"%","",1),64)
		if err != nil{
			return nil,err
		}

		disks = append(disks, Disk{Filesystem:data[6*i],Size:data[6*i+1],Used:data[6*i+2],Avail:data[6*i+3],Use:use,MountedOn:data[6*i+5]})

	}

	return disks,nil
}
