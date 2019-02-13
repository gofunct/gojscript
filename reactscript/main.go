package main

import (
	"fmt"
	"github.com/gofunct/goexec"
	"os"
)

func init() {

}

var (
	binflags = "-pkg=main -prefix=server/data -debug -o server/data/..."
	exe      = goexec.NewCommand("cloudscript", "a scripting utility tool for cloud development", "0.1")
)

func main() {
	exe.Act("gen", "go generate all", func(cmd *goexec.Command) error {
		cmd.AddScript(`go generate {{ .pwd }}/...`)
		return cmd.Run()
	})
	exe.Act("setup", "setup project", func(cmd *goexec.Command) error {
		cmd.AddScript(`go mod vendor`)
		cmd.AddScript("yarn install")
		return cmd.Run()
	})

	exe.Act("gobin", "Download go mod binaries", func(cmd *goexec.Command) error {
		return cmd.GoBin([]string{"github.com/jteeuwen/go-bindata/go-bindata", "github.com/olebedev/on", "github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs"})
	})
	exe.Act("gobindata", "Download go mod binaries", func(cmd *goexec.Command) error {
		cmd.AddScript(`go-bindata -pkg=main -prefix=server/data -debug server/data/...`)
		return cmd.Run()
	})
	exe.Act("lint", "lint go and js files", func(cmd *goexec.Command) error {
		cmd.AddScript(`yarn run eslint || true`)
		cmd.AddScript(`golint ./... || true`)
		return cmd.Run()
	})

	if err := exe.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
