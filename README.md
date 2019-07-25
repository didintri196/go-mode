# go-mode 
GOLANG EASY MODE DEBUG OR PRODUCTION

This package is made so that it is easier to create two project modes so that it is not difficult to find the data that you have printed on the coding page, you just need to change the status to release. All testing posts will disappear.

## Install

Install the package with:

```bash
go get github.com/didintri196/go-mode
```

Import it with:

```go
import program "github.com/didintri196/go-mode"
```

and use `program` as the package name inside the code.

## Example

Please check the example folder for details.

```go
package main

import (
	program "github.com/didintri196/go-mode"
)

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

```
## MODE

| Mode          | Description            |
| ------------ | :--------------- |
| program.DebugMode     | DEBUG MODE           |
| program.ReleaseMode     | RELEASE MODE           |

## TIPE

| Tipe          | Description            |
| ------------ | :--------------- |
| log     | Using log to Print          |
| fmt     | Using fmt to Print          |