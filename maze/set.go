package maze

import (
	"fmt"
	"strings"
)

type Set struct {
	items map[interface{}]struct{}
}

var itemExists = struct{}{}

func NewSet(items ...interface{}) *Set {
	this := &Set{items: make(map[interface{}]struct{})}
	if len(items) > 0 {
		this.Add(items)
	}
	return this
}

func (this *Set) Add(items ...interface{}) {
	for _, item := range items {
		this.items[item] = itemExists
	}
}

func (this *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(this.items, item)
	}
}

func (this *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, contains := this.items[item]; !contains {
			return false
		}
	}
	return true
}

func (this *Set) Size() int {
	return len(this.items)
}

func (this *Set) Empty() bool {
	return this.Size() == 0
}

func (this *Set) Clear() {
	this.items = make(map[interface{}]struct{})
}

func (this *Set) Values() []interface{} {
	values := make([]interface{}, this.Size())
	count := 0
	for item := range this.items {
		values[count] = item
		count++
	}
	return values
}

func (this *Set) String() string {
	str := "Set{ "
	items := []string{}
	for item := range this.items {
		items = append(items, fmt.Sprintf("%v", item))
	}
	str += strings.Join(items, ", ")
	str += " }"
	return str
}
