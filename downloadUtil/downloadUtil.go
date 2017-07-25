package downloadUtil

import (
	"github.com/cavaliercoder/grab"
	"fmt"
	"time"
	//"os"
	"os"
	"github.com/pkg/errors"
)

func Download(url string,dst string) (bool,error) {
	// create client
	client := grab.NewClient()
	req,err := grab.NewRequest(dst, url)
	if err!=nil{
		//println(err.Error())
		return false,err
	}
	// start download
	fmt.Printf("Downloading %v...\n",req.URL())

	resp := client.Do(req)
	if resp == nil || resp.HTTPResponse == nil {
		return false,errors.New("response is nil")
	}
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	timeout := make(chan bool, 1)
	// 并发执行一个函数，等待1s后向timeout写入true
	go func() {
		time.Sleep(4 * time.Minute)
		timeout <- true
	}()

	Loop:
		for {
			select {
			case <-timeout:
				//time out
				println("time out ")
				//resp.Cancel()

				return false,nil
				//break Loop
			case <-t.C:
				fmt.Printf("  transferred %v / %v bytes (%.2f%%), %s\n",
					resp.BytesComplete(),
					resp.Size,
					100*resp.Progress(),resp.Filename)

			case <-resp.Done:
				// download is complete
				break Loop

			}
		}

		// check for errors
		if err := resp.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
			return false,err
		}

		fmt.Printf("Download saved to ./%v \n", resp.Filename)
		return true,nil
}