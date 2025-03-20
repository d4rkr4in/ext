package main

import (
	_ "github.com/d4rkr4in/ext/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
