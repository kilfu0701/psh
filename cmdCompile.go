package main

import (
	//"fmt"
	"io/ioutil"
	"os"

	"langs"
)

var cmdCompile = &Command{
	UsageLine: "compile file.ext",
	Short:     "Compile a script.",
	Long:      `Compile a file into colorful script.
Usage:
    psh compile file.ext
    psh compile -style=default file.ext
    psh compile -style=mono file.ext
`,
	CustomFlags: false,
}

var style string

func init() {
	cmdCompile.Run = runInit
	cmdCompile.Flag.StringVar(&style, "style", "default", "highlight style")
	logI("Style: %v", style)
}

func runInit(cmd *Command, args []string) int {
	langs.Comment("//", "$")

	path, _ := os.Getwd()

	for _, f := range args {
		inputFile := path + "/" + f
		logI("compile \"%v\" ...", inputFile)
		_, err := ioutil.ReadFile(inputFile)
		if err != nil {
			logE("File not found: %s", inputFile)
			continue
		}
	}

	return 0
}
