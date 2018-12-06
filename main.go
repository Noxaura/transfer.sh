package main

import "github.com/Noxaura/transfer.sh/cmd"

func main() {
	app := cmd.New()
	app.RunAndExitOnError()
}
