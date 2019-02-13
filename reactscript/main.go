package main

import (
	"fmt"
	"github.com/gofunct/goexec"
	"os"
)

func init() {

}

var (
	exe      = goexec.NewCommand("cloudscript", "a scripting utility tool for cloud development", "0.1")
)

func main() {
	exe.Act("gen", "go generate all", func(cmd *goexec.Command) error {
		cmd.AddScript(`go generate {{ .pwd }}/...`)
		return cmd.Run()
	})
	exe.Act("js", "setup project", func(cmd *goexec.Command) error {
		cmd.AddScript("yarn install")
		return cmd.Run()
	})

	exe.Act("gobin", "Download go mod binaries", func(cmd *goexec.Command) error {
		return cmd.GoBin([]string{"github.com/jteeuwen/go-bindata/go-bindata", "github.com/olebedev/on", "github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs"})
	})
	exe.Act("bindata", "generate embeded assets", func(cmd *goexec.Command) error {
		cmd.AddScript(`go-bindata -pkg=base -prefix=modules/base/data -debug -o modules/base/base.go modules/base/data/...`)
		return cmd.Run()
	})

	if err := exe.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
