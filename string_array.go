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
        s += v
    }
    return s
}

func (a *StringArray) String() string {
	return strings.Join(*a, ",")
}
