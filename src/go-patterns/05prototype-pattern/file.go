package main

import "fmt"

type file struct {
	name string
}

func (f *file) print(s string) {
	fmt.Println(s + f.name)
}

func (f *file) clone() inode {
	return &file{
		name: f.name + "_clone",
	}
}
