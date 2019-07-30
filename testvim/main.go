package main

import "fmt"

func main() {
	nvim := NewNeoVim()
	fmt.Println(nvim.Parent.Parent)
}

// Editor indicates the type of editor.
type Editor struct {
	Name   string
	Parent *Editor
}

// NewEmacs is emacs constructor
func NewEmacs() *Editor {
	return &Editor{Name: "Emacs"}
}

// NewVi is vi constructor
func NewVi() *Editor {
	return &Editor{Name: "Vi"}
}

// NewVim is vim constructor
func NewVim() *Editor {
	return &Editor{
		Name:   "Vim",
		Parent: NewVi(),
	}
}

// NewNeoVim is neovim constructor
func NewNeoVim() *Editor {
	return &Editor{
		Name:   "Vim",
		Parent: NewVim(),
	}
}
