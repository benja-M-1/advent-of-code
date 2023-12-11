package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

type File struct {
	Name string
	Size int
}

type Dir struct {
	Fs     *Fs
	Name   string
	Files  map[string]File
	Dirs   map[string]*Dir
	Parent *Dir
}

func (d *Dir) Identifier() string {
	if d.Parent == nil {
		return d.Name
	}

	return fmt.Sprintf("%s/%s", d.Parent.Identifier(), d.Name)
}

func (d *Dir) AddFile(f File) {
	if d.Files == nil {
		d.Files = make(map[string]File)
	}

	d.Files[f.Name] = f
}

func (d *Dir) AddDir(dir Dir) {
	if d.Dirs == nil {
		d.Dirs = make(map[string]*Dir)
	}

	d.Dirs[dir.Name] = &dir
	dir.Parent = d
	d.Fs.AddDir(&dir)
}

func (d *Dir) Size() int {
	var size int
	for _, f := range d.Files {
		size += f.Size
	}

	for _, dir := range d.Dirs {
		size += dir.Size()
	}

	return size
}

func (d *Dir) String() string {
	var out string

	size := d.Size()
	if size < 100000 {
		out += fmt.Sprintf("- ✅ %s (dir, size=%d)\n", d.Name, size)
	} else {
		out += fmt.Sprintf("- ❌ %s (dir, size=%d)\n", d.Name, size)
	}

	for _, f := range d.Files {
		out += fmt.Sprintf("\t- %s (file, size=%d)\n", f.Name, f.Size)
	}

	for _, dir := range d.Dirs {
		for _, s := range strings.Split(dir.String(), "\n") {
			out += fmt.Sprintf("\t%s\n", s)
		}
	}

	return out
}

type Fs struct {
	Root    *Dir
	Current *Dir
	all     map[string]*Dir
}

func (fs *Fs) AddRoot() {
	r := Dir{Name: "/", Fs: fs}
	fs.Root = &r
	fs.AddDir(&r)
}

func (fs *Fs) AddDir(dir *Dir) {
	if fs.all == nil {
		fs.all = make(map[string]*Dir)
	}

	fs.all[dir.Identifier()] = dir
	dir.Fs = fs
}

func (fs *Fs) ChangeDir(name string) error {
	if name == ".." {
		fs.Current = fs.CurrentDir().Parent

		return nil
	}

	if name == "/" {
		fs.Current = fs.Root

		return nil
	}

	id := fmt.Sprintf("%s/%s", fs.CurrentDir().Identifier(), name)
	dir, ok := fs.all[id]
	if !ok {
		return fmt.Errorf("dir %v not found", name)
	}

	fs.Current = dir

	return nil
}

func (fs *Fs) CurrentDir() *Dir {
	if fs.Current == nil {
		fs.Current = fs.Root
	}
	return fs.Current
}

func (fs *Fs) Dirs() []*Dir {
	return maps.Values(fs.all)
}

func (fs *Fs) String() string {
	return fs.Root.String()
}

type BySize []*Dir

func (s BySize) Len() int           { return len(s) }
func (s BySize) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s BySize) Less(i, j int) bool { return s[i].Size() < s[j].Size() }
