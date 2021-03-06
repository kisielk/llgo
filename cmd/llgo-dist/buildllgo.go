// Copyright 2012 Andrew Wilkins.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

const llgopkgpath = "github.com/axw/llgo/llgo"

var (
	// llgobin is the path to the llgo command.
	llgobin string
)

func buildLlgo() error {
	log.Println("Building llgo")

	cgoCflags := llvmcflags
	cgoLdflags := fmt.Sprintf("%s -Wl,-L%s -lLLVM-%s",
		llvmldflags, llvmlibdir, llvmversion)

	var output []byte
	var err error
	ldflags := "-r " + llvmlibdir
	args := []string{"get", "-ldflags", ldflags}
	if strings.HasSuffix(llvmversion, "svn") {
		args = append(args, "-tags")
		args = append(args, "llvmsvn")
	}
	args = append(args, llgopkgpath)
	cmd := exec.Command("go", args...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "CGO_CFLAGS="+cgoCflags)
	cmd.Env = append(cmd.Env, "CGO_LDFLAGS="+cgoLdflags)

	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", string(output))
		return err
	}

	pkg, err := build.Import(llgopkgpath, "", build.FindOnly)
	if err != nil {
		return err
	}
	llgobin = path.Join(pkg.BinDir, "llgo")
	log.Printf("Built %s", llgobin)

	return nil
}
