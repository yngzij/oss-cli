package main

import (
	"fmt"

	"oss-cli/cmd"
	"oss-cli/configs"
)

func main() {
	configs.LoadConfig("./configs/dev/config.toml")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
	return
}
