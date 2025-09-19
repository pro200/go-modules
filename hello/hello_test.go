package hello_test

import (
	"fmt"
	"testing"

	"github.com/pro200/go-modules/hello"
)

func TestHello(t *testing.T) {
	fmt.Println(hello.Hello())
}
