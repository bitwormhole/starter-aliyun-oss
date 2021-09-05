package main

import (
	"github.com/bitwormhole/starter"
	ossstarter "github.com/bitwormhole/starter-aliyun-oss"
)

func main() {
	starter.InitApp().Use(ossstarter.Module()).Run()
}
