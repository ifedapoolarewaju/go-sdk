package main

import (
	"fmt"
	"github.com/transloadit/go-sdk"
)

func main() {

	// Create client
	options := transloadit.DefaultConfig
	options.AuthKey = "TRANSLOADIT_KEY"
	options.AuthSecret = "TRANSLOADIT_SECRET"
	client, err := transloadit.NewClient(options)
	if err != nil {
		panic(err)
	}

	// Initialize new assembly
	assembly := client.CreateAssembly()

	// Add a file to upload
	assembly.AddFile("image", "../../fixtures/lol_cat.jpg")

	// Add instructions, e.g. resize image to 75x75px
	assembly.AddStep("resize", map[string]interface{}{
		"robot":           "/image/resize",
		"width":           75,
		"height":          75,
		"resize_strategy": "pad",
		"background":      "#000000",
	})

	// Start the upload
	info, err := assembly.Upload()
	if err != nil {
		panic(err)
	}

	// All files have now been uploaded and the assembly has started but no
	// results are available yet since the conversion has not finished.
	// The AssemblyWatcher provides functionality for polling until the assembly
	// has ended.
	waiter := client.WaitForAssembly(info.AssemblyUrl)
	info = <-waiter.Response

	fmt.Printf("You can view the result at: %s\n", info.Results["resize"][0].Url)

}
