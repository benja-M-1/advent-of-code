package main

import (
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	r1 := One(string(input))
	fmt.Printf("puzzle 1: %v\n", r1)
	r2 := Two(string(input))
	fmt.Printf("puzzle 2: %v\n", r2)
}

func One(input string) int {
	fs := parse(strings.Trim(input, "\n"))

	sum := 0
	for _, dir := range fs.Dirs() {
		if size := dir.Size(); size < 100000 {
			sum += size
		}
	}

	return sum
}

func parse(input string) Fs {
	fs := Fs{}
	fs.AddRoot()

	for _, cmd := range strings.Split(input, "$") {
		cmd = strings.TrimSpace(cmd)
		switch {
		case strings.HasPrefix(cmd, "cd"):
			dir := strings.TrimPrefix(cmd, "cd ")
			if err := fs.ChangeDir(dir); err != nil {
				panic(err)
			}
		case strings.HasPrefix(cmd, "ls"):
			out := strings.Split(cmd, "\n")
			for _, el := range out[1:] {
				if strings.HasPrefix(el, "dir") {
					name := strings.TrimPrefix(el, "dir ")
					fs.CurrentDir().AddDir(Dir{Name: name})

					continue
				}

				parts := strings.Split(el, " ")
				size, _ := strconv.Atoi(parts[0])
				f := File{
					Name: parts[1],
					Size: size,
				}
				fs.CurrentDir().AddFile(f)
			}
		}
	}

	return fs
}

func Two(input string) int {
	fs := parse(strings.Trim(input, "\n"))

	dirs := BySize(fs.Dirs())
	sort.Sort(dirs)

	totalSize := fs.Root.Size()
	unusedSpace := 70000000 - totalSize
	neededSpace := 30000000 - unusedSpace

	for _, d := range dirs {
		if d.Size() >= neededSpace {
			return d.Size()
		}
	}

	return 0
}
