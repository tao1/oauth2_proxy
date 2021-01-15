package main

import (
	"strings"
)

type StringArray []string

func (a *StringArray) Set(s string) error {
	*a = append(*a, s)
	return nil
}

func (a *StringArray) Get() interface{} {
	var s string
	for _, v := range *a {
<<<<<<< HEAD
		s += v
=======
		if v != "" {
			s += v
		}
>>>>>>> f2baac60436f8ea6d02c3291f5ddd372c14f6437
	}
	if s != "" {
		return s
	}
	return nil
}

func (a *StringArray) String() string {
	return strings.Join(*a, ",")
}
