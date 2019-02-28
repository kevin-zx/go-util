package imageUtil

import (
	"fmt"
	"os"
	"testing"
)

func TestRandomResize(t *testing.T) {

	testDataPath := "../testData/image/"
	testDataDir, err := os.Open(testDataPath)

	if err != nil {
		panic(err)
	}
	dirs, err := testDataDir.Readdir(-1)
	testDataDir.Close()
	if err != nil {
		panic(err)
	}
	for _, d := range dirs {
		fmt.Println(d.Name())
		img := LoadImage(testDataPath + d.Name())
		img = RandomResize(img)
		img = RandomFilter(img)
		SaveImage(testDataPath+"r_"+d.Name(), RandomResize(img))
	}
	//
}
