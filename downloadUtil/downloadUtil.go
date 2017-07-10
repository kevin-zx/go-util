package downloadUtil

import (
	"github.com/cavaliercoder/grab"
	"fmt"
	"time"
	//"os"
	"os"
)

func Download(url string,dst string) (bool,error) {
	// create client
	client := grab.NewClient()
	req,_ := grab.NewRequest(dst, url)

	// start download
	fmt.Printf("Downloading %v...\n",req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	Loop:
		for {
			select {
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