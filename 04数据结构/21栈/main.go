package main

import "github.com/pkg/errors"

type Stack struct {
	MaxTop int
	Top int
	arr [5]int
}

func (s *Stack)push(val int)(err error) {
	if s.Top == s.MaxTop {
		err = errors.New("stack full")
		return
	}
	s.Top++
	s.arr[s.Top] = val
	return nil
}

func (s *Stack)Pop() (val int, err error) {
	if s.Top == -1 {
		err = errors.New("stack empty!")
		return
	}
	val = s.arr[s.Top]
	s.Top--
	return val,nil
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top: -1,
	}
}
