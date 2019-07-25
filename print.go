package print

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
)

var mode = 0
var DefaultWriter io.Writer = os.Stdout

func SetMode(value string) {
	switch value {
	case DebugMode, "":
		mode = 1
	case ReleaseMode:
		mode = 2
		log.SetOutput(ioutil.Discard)

	default:
		panic("print mode unknown: " + value)
	}
}

func Print(val string) {
	log.Println(val)
}
