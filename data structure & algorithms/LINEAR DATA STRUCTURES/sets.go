package main

import "fmt"

type Set struct {
	elements map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		elements: make(map[int]struct{}),
	}
}

func (s *Set) Add(element int) {
	s.elements[element] = struct{}{}
}

func (s *Set) Remove(element int) {
	delete(s.elements, element)
}

func (s *Set) Contains(element int) bool {
	_, exists := s.elements[element]
	return exists
}

func (s *Set) Display() {
	for key := range s.elements {
		fmt.Printf("%d ", key)
	}
	fmt.Println()
}
