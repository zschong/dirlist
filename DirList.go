package dirlist

import (
	"encoding/json"
	"io/fs"
	"os"
	"sort"
	"time"
)

type FileInfo struct {
	Name  string      // base name of the file
	Size  int64       // length in bytes for regular files; system-dependent for others
	IsDir bool        // abbreviation for Mode().IsDir()
	Mode  fs.FileMode // file mode bits
	Mtime time.Time
}

type DirList struct {
	Path string
	List []*FileInfo
}

func New(path string) *DirList {
	return &DirList{
		Path: path,
		List: make([]*FileInfo, 0),
	}
}

func (d *DirList) Read() []*FileInfo {
	if len(d.Path) == 0 {
		return nil
	}
	dlist, err := os.ReadDir(d.Path)
	if err != nil {
		return nil
	}
	files := make([]*FileInfo, 0)
	for _, entry := range dlist {
		info, err := entry.Info()
		if err == nil {
			f := new(FileInfo)
			f.Name = info.Name()
			f.Mtime = info.ModTime()
			f.Size = info.Size()
			f.Mode = info.Mode()
			f.IsDir = info.IsDir()
			files = append(files, f)
		}
	}
	d.List = files
	return files
}

func (d *DirList) SortByName() []*FileInfo {
	files := d.List
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name > files[j].Name
	})
	return files
}

func (d *DirList) SortBySize() []*FileInfo {
	files := d.List
	sort.Slice(files, func(i, j int) bool {
		return files[i].Size > files[j].Size
	})
	return files
}

func (d *DirList) SortByTime() []*FileInfo {
	files := d.List
	sort.Slice(files, func(i, j int) bool {
		return files[i].Mtime.Before(files[j].Mtime)
	})
	return files
}

func (d *DirList) SortByType() []*FileInfo {
	files := d.List
	sort.Slice(files, func(i, j int) bool {
		return files[i].IsDir
	})
	return files
}

func (d *DirList) Json() string {
	j, err := json.MarshalIndent(d.SortByType(), "", "    ")
	if err != nil {
		return ""
	}
	return string(j)
}
