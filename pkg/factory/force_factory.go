package main

type matrix struct {
	name string
}

func NewMatrix(name string) *matrix {
	m := new(matrix)
	return m
}
