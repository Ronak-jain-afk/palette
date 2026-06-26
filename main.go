package main

import "github.com/Ronak-jain-afk/palette/cmd"

var version = "dev"

func main() {
	cmd.Version = version
	cmd.Execute()
}
