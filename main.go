package main

import (
	"bytes"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/hexbotio/hex-plugin"
)

type HexLocal struct {
}

func (g *HexLocal) Perform(args hexplugin.Arguments) (resp hexplugin.Response) {
	output := ""
	success := true
	var o bytes.Buffer
	var e bytes.Buffer
	c := exec.Command("/bin/sh", "-c", args.Command)
	c.Stdout = &o
	c.Stderr = &e
	err := c.Run()
	output = o.String()
	if err != nil {
		output = output + "\n" + err.Error() + "\n" + e.String()
		success = false
	}
	resp = hexplugin.Response{
		Output:  output,
		Success: success,
	}
	return resp
}

func main() {
	var pluginMap = map[string]plugin.Plugin{
		"action": &hexplugin.HexPlugin{Impl: &HexLocal{}},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: hexplugin.GetHandshakeConfig(),
		Plugins:         pluginMap,
	})
}
