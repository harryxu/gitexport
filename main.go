package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

var outputDir = "/tmp"
var diff1 = "60d8383..dd0b2ff"
var repo = "/Users/harry/dev/data/www/bigecko/hunt"

func main() {
	// dir, err := os.Getwd()

	os.Chdir(repo)

	filelog(diff1)

	fmt.Printf("repo name: %s \n", repoName())

	fmt.Printf("outdir name: %s \n", outDirName())

	fmt.Println("end")
}

func filelog(diff string) {
	args := []string{"diff", "--name-status", diff}

	output, err := exec.Command("git", args...).Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
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
