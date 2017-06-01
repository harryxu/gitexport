package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"runtime"

	"github.com/spf13/cobra"
)

var outputDir = os.TempDir()
var revison string

var repoRoot string

// RootCmd Default Command
var RootCmd = &cobra.Command{
	Use: "gitexport -r [revison]",
	Run: func(cmd *cobra.Command, args []string) {
		if strings.TrimSpace(revison) == "" {
			revison = defaultDiff()
		}
		export(filelog(revison))
	},
}

func main() {
	RootCmd.PersistentFlags().StringVarP(&revison, "revision", "r", "", "Revision as git diff")

	output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(err)
	}
	repoRoot = strings.Trim(string(output), "\n")

	os.Chdir(repoRoot)

	RootCmd.Execute()
}

// Get files list by git diff.
func filelog(diff string) []string {
	fmt.Printf("diff %s\n", diff)
	output, err := exec.Command("git", "diff", "--name-status", diff).Output()
	if err != nil {
		panic(err)
	}

	files := strings.Split(string(output), "\n")

	return files
}

func getLatestRevHash() string {
	output, err := exec.Command("git", "log", "--pretty=format:%h", "-n1").Output()
	if err != nil {
		panic(err)
	}

	rev := string(output)
	if len(rev) == 0 {
		return ""
	}

	return rev
}

func getRevisionByStep(lastStep string) string {
	output, err := exec.Command("git", "log",
		"--pretty=format:%h",
		"-n1",
		"--skip="+lastStep).Output()
	if err != nil {
		panic(err)
	}

	return string(output)
}

func defaultDiff() string {
	return getRevisionByStep("1") + ".." + getLatestRevHash()
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

	openDir(destDir)
}

func openDir(path string) {
	osname := runtime.GOOS
	var command string

	if osname == "windows" {
		command = "explorer.exe"
		path = strings.Replace(path, "/", "\\", -1)
	} else if osname == "darwin" {
		command = "open"
	} else if osname == "linux" {
		command = "xdg-open"
	}

	exec.Command(command, path).Run()
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
