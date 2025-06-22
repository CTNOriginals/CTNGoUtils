package tools

import (
	"fmt"
	"slices"

	cstring "github.com/CTNOriginals/CTNGoUtils/string"
	cstruct "github.com/CTNOriginals/CTNGoUtils/struct"
)

type Cursor[T comparable] struct {
	Original *[]T
	Content  []T
}

func NewCursor[T comparable](content *[]T) Cursor[T] {
	var clone = make([]T, len(*content))
	copy(clone, *content)

	var cur = Cursor[T]{
		Original: content,
		Content:  clone,
	}

	slices.Reverse(cur.Content)

	return cur
}

// #region Getters
func (this Cursor[T]) String() string {
	return fmt.Sprintf("Cursor[%T] {\n%s\n}", *this.Original, cstring.Indent(cstruct.ToString(this), 1, " "))
}

func (this Cursor[T]) IsStart() bool {
	return len(this.Content) == len(*this.Original)
}
func (this Cursor[T]) IsEnd() bool {
	return len(this.Content) == 0
}

// Peek the current item the cursor is on without consumingit.
func (this Cursor[T]) Peek() T {
	return this.Content[len(this.Content)-1]
}
func (this Cursor[T]) PeekBack() T {
	this.Backup()
	return this.Read()
}

// Returns the index of the current item.
// The index is based on the original content.
func (this Cursor[T]) CurrentIndex() int {
	// fmt.Println(this.Peek(), len(this.Content), cap(this.Content))
	return cap(this.Content) - len(this.Content)
}

//#endregion

// Consume the current item, effectively setting the contents length to -1 of its current lengt.
func (this *Cursor[T]) Consume() {
	this.Content = this.Content[:len(this.Content)-1]
}

// Goes back one read item.
// For this to work, the cursor has to have been read at least once before.
func (this *Cursor[T]) Backup() {
	this.Content = this.Content[:len(this.Content)+1]
}

// Reads and returns the current item before consuming it
func (this *Cursor[T]) Read() (item T) {
	item = this.Peek()
	this.Consume()
	return item
}
