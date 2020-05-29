package krassedinge

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

func DestroyMe() {
	fmt.Println("Hello from ilib/krasseDinge/DestroyMe Version 2.2.1")
	result := cmp.Diff("","")
	fmt.Println(result)
}
