package main

import "fmt"

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(s string) {
	fmt.Println(s + f.name)
	for _, i := range f.children {
		i.print(s + s)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
