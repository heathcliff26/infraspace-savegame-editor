package cli

import (
	"fmt"
	"strconv"
	"strings"
)

type IntFlag struct {
	Value int

	set bool
}

func (f *IntFlag) Set(arg string) error {
	i, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("failed to convert \"%s\" to integer: %v", arg, err)
	}
	f.Value = i
	f.set = true
	return nil
}

func (f *IntFlag) String() string {
	return strconv.Itoa(f.Value)
}

func (f *IntFlag) IsSet() bool {
	return f.set
}

type IntMapFlag map[string]int

func (f IntMapFlag) Set(arg string) error {
	keyvalue := strings.Split(arg, "=")
	if len(keyvalue) != 2 {
		return fmt.Errorf("unexcpected input '%s' for resource, expected <resource>=<value>", arg)
	}
	key := keyvalue[0]
	if _, ok := f[key]; ok {
		return fmt.Errorf("got multiple inputs for %s", keyvalue[0])
	}
	i, err := strconv.Atoi(keyvalue[1])
	if err != nil {
		return fmt.Errorf("failed to convert \"%s\" to integer: %v", arg, err)
	}
	f[key] = i
	return nil
}

func (f IntMapFlag) String() string {
	return "Can't convert IntMapFlag to string"
}
