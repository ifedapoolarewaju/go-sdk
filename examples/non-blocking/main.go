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
	assembly := client.NewAssembly()

	// Add an file to upload
	assembly.AddFile("image", "../../fixtures/lol_cat.jpg")

	// Instructions will be read from the template
	// `02a8693053cd11e49b9ba916b58830db` stored on Transloadit's servers.
	assembly.TemplateId = "02a8693053cd11e49b9ba916b58830db"

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
	assemblyInfo := <-waiter.Response

	fmt.Printf("You can view the result at: %s\n", assemblyInfo.Results["resize"][0].Url)

}
