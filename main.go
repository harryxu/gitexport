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
var repo = "/Users/harry/dev/data/www/bigecko/hunt"

func main() {
	// dir, err := os.Getwd()

	os.Chdir(repo)

	export(filelog(diff1))

	fmt.Printf("repo name: %s \n", repoName())

	fmt.Printf("output dir: %s\n", outputDir)

	fmt.Printf("outdir name: %s \n", outDirName())

	fmt.Println("end")
}

func filelog(diff string) []string {
	output, err := exec.Command("git", "diff", "--name-status", diff).Output()
	if err != nil {
		panic(err)
	}

	files := strings.Split(string(output), "\n")

	return files
}

func repoName() string {
	repoRoot, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(err)
	}

	return strings.Trim(path.Base(string(repoRoot)), "\n")
}

func outDirName() string {

	t := time.Now()
	outdir := repoName() + "-" + t.Format("20060102150405")

	return outdir
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
				src := path.Join(repo, file)
				dst := path.Join(destDir, file)
				copyFileContents(src, dst)
			}
		}
	}

	if len(deletes) > 0 {
		fmt.Printf("Deleted files: \n %s \n", strings.Join(deletes, "\n"))
	}

	fmt.Printf("Exported to: %s\n", destDir)
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
