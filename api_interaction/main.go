package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func printFile(filename string) {

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

func addString(sh *shell.Shell, sendText string) string {
	cid, err := sh.Add(strings.NewReader(sendText))
	if err != nil {
		log.Fatal(err)
	}
	return cid
}

func main() {
	var outfile string = "./outfile.txt"
	var target string = "QmP8jTG1m9GSDJLCbeWhVSVgEzCPPwXRdCRuJtQ5Tz9Kc9"

	sh := shell.NewShell("localhost:5001")

	// get and display from target CID
	sh.Get(target, outfile)
	printFile(outfile)

	// add string to ipfs
	cid := addString(sh, "This is a string to add")
	println("added string to %s", cid)

	// get and display from added CID
	sh.Get(cid, outfile)
	printFile(outfile)

}

// remote working?
