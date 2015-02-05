// The StringArray type is borrowed from NSQ
// https://github.com/bitly/nsq/blob/master/util/string_array.go

package util

import (
	"strings"
)

type StringArray []string

func (a *StringArray) Set(s string) error {
	*a = append(*a, s)
	return nil
}

func (a *StringArray) String() string {
	return strings.Join(*a, ",")
}
