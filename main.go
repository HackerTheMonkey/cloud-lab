package main

import (
	"github.com/HackerTheMonkey/codematters/cloud"
	"github.com/HackerTheMonkey/codematters/core"	
)

func main() {
	println("init...")
	core.Init("codematters.io", cloud.DigitalOceanAdapter{})
}
