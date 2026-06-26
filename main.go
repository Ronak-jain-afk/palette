package main

import "github.com/Ronak-jain-afk/palette/cmd"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.Version = version
	cmd.Execute()
}
