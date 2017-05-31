package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

var outputDir = os.TempDir()
var diff1 = "60d8383..dd0b2ff"

var repoRoot string

func main() {
	output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(err)
	}
	repoRoot = strings.Trim(string(output), "\n")

	os.Chdir(repoRoot)

	export(filelog(diff1))
}

func filelog(diff string) []string {
	output, err := exec.Command("git", "diff", "--name-status", diff).Output()
	if err != nil {
		panic(err)
	}

	files := strings.Split(string(output), "\n")

	return files
}

func outDirName() string {
	repoName := strings.Trim(path.Base(string(repoRoot)), "\n")
	dirName := repoName + "-" + time.Now().Format("20060102150405")
	return dirName
}

func export(files []string) {
	var deletes []string
	destDir := path.Join(outputDir, outDirName())
	os.MkdirAll(destDir, os.ModePerm)
	for _, f := range files {
		if len(f) > 0 {
			slices := strings.Split(f, "\t")
			opt := slices[0]
			file := slices[1]

			if opt == "D" {
				deletes = append(deletes, file)
			} else {
				src := path.Join(repoRoot, file)
				dst := path.Join(destDir, file)
				copyFileContents(src, dst)
			}
		}
	}

	fmt.Printf("Exported to: %s\n", destDir)

	if len(deletes) > 0 {
		fmt.Printf("Deleted files: \n %s \n", strings.Join(deletes, "\n"))
	}
}

// https://stackoverflow.com/a/21067803/157811
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	os.MkdirAll(path.Dir(dst), os.ModePerm)

	out, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		panic(err)
	}
	err = out.Sync()
	return
}
