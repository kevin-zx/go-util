package stringUtil

import (
	"testing"
)

func TestClear(t *testing.T) {
	testInstances := [][]string{
		{`	        ​Environment Test Chamber for Air Conditioner_Househ
old/Commercial Air Conditioner Labs_Products_Shanghai Satake Cool-heat & Control Technique Co., Ltd.	`, "Environment Test Chamber for Air Conditioner_Household/Commercial Air Conditioner Labs_Products_Shanghai Satake Cool-heat & Control Technique Co., Ltd."},
		{`        空调压缩机寿命台_家用及商用空调试验室_产品中心_上海佐竹冷热控制技术有限公司
`, "空调压缩机寿命台_家用及商用空调试验室_产品中心_上海佐竹冷热控制技术有限公司"},
	}
	for _, ti := range testInstances {
		if Clear(ti[0]) != ti[1] {
			t.Error(ti[0] + " :clear err")
		}
	}
}
