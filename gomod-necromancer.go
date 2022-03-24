package main

import (
	"debug/buildinfo"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/mod/modfile"
)

func main() {

	var targetFile = flag.String("target", "PLEASE_SET_TARGET", "the binary file we want to inspect")
	var modfileName = flag.String("modfileName", "PLEASE_SET_MODFILENAME", "")
	flag.Parse()

	bi, err := buildinfo.ReadFile(*targetFile)
	if err != nil {
		log.Fatalf("Error reading buildinfo from %s: %v", *targetFile, err)
	}

	file, err := os.Open(*modfileName)
	contents, err := io.ReadAll(file)
	gomod, err := modfile.Parse(*modfileName, contents, nil)
	_ = gomod
	mod := modfile.File{}

	mod.AddModuleStmt(bi.Path)
	mod.AddGoStmt(bi.GoVersion)
	for _, dep := range bi.Deps {
		mod.AddRequire(dep.Path, dep.Version)
		if dep.Replace != nil {
			mod.AddReplace(dep.Path, dep.Version, dep.Replace.Path, dep.Replace.Version)
		}
	}

	formatted, err := mod.Format()
	fmt.Println(string(formatted))
}
