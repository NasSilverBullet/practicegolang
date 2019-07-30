package main

import "fmt"

func main() {
	nvim := NewNeoVim()
	fmt.Println(nvim.Parent.Parent)

	hogeHuga := "HOGE"
	fmt.Println(hogeHuga)

	text := `
These commands move slowly the cursor position.
	`
	fmt.Println(text)

	var hoge = "HOGE"                                             // 何かのコメント
	var fuga = "FUGA"                                             // 素敵なコメント
	var fuganogehogehogehogehogehoge = "HOHOGAHOGEHOGEHOGEHOGHEO" // 素敵なコメント
	fmt.Println(hoge, fuga, fuganogehogehogehogehogehoge)

	longText := `
こんな

    改行と行頭空白文字混じりの

        文章があったとしても一発で1行にできます
	`
	fmt.Println(longText)
}

// Editor indicates the type of editor.
type Editor struct {
	Name   string
	Parent *Editor
}

// NewEmacs is emacs constructor
func NewEmacs() *Editor {
	return &Editor{
		Name: "Emacs",
	}
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
