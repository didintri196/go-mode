package program

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

type Program struct {
	output io.Writer
	mu     sync.Mutex
}
type Configuration struct {
	Mode string
	Tipe string
}
type Fmt struct{}
type Log struct{}

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
)

func SetOutput(w io.Writer) {
	var m Program
	m.mu.Lock()
	defer m.mu.Unlock()
	m.output = w

}

var mode = 1
var tipe = 1

// TIPE 1 = LOG
// TIPE 2 = FMT

func SetMode(Config Configuration) {
	switch Config.Mode {
	case DebugMode, "":
		mode = 1
		switch Config.Tipe {
		case "log", "":
			tipe = 1
		case "fmt":
			tipe = 2
		default:
			panic("print mode unknown: " + Config.Tipe)
		}
	case ReleaseMode:
		mode = 2
		tipe = 2
		log.SetOutput(ioutil.Discard)
		SetOutput(ioutil.Discard)
	default:
		panic("print mode unknown: " + Config.Mode)
	}
}

func Println(val string) {
	if tipe == 1 {
		if mode == 1 {
			log.Println(val)
		} else {
		}
	} else if tipe == 2 {
		if mode == 1 {
			fmt.Println(val)
		} else {
		}
	}

}

func Printf(format string, a ...interface{}) {
	//LOG
	if tipe == 1 {
		if mode == 1 {
			log.Printf(format, a...)
		} else {
		}
		//FMT
	} else if tipe == 2 {
		if mode == 1 {
			fmt.Printf(format, a...)
		} else {
		}
	}
}

func Print(a ...interface{}) {
	//LOG
	if tipe == 1 {
		if mode == 1 {
			log.Print(a...)
		} else {
		}
		//FMT
	} else if tipe == 2 {
		if mode == 1 {
			fmt.Print(a...)
		} else {
		}
	}
}
