package server_monitor

import (
	"github.com/kevin-zx/go-util/commandUtil"
	"fmt"
)

func monitor_disk_free_space(caption string) {
	//wmic LogicalDisk where "Caption='C:'" get FreeSpace,Size /value 查看c盘使用率
	caption_arg := fmt.Sprintf(`"Caption='%s:'"`, caption)
	commandUtil.ExecCommand("wmic", []string{"where", caption_arg, "get", "FreeSpace,Size", "/value"})

}
