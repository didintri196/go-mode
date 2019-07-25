package main

import program "github.com/didintri196/go-mode"

func main() {
	config := program.Configuration{
		Mode: program.DebugMode,
		Tipe: "log",
	}
	program.SetMode(config)
	program.Println("Halo Didin")
	program.Printf("Halo %s", "Didin")
	program.Print("Halo Didin")
}
