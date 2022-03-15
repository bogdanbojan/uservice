package main

import (
	"github.com/ardanlabs/conf"
	"github.com/dimfeld/httptreemux/v5"
)

func main() {
	conf.New()
	httptreemux.New()
}
