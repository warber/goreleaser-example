package lib_test

import (
	"testing"

	"github.com/warber/goreleaser-example/pkg/lib"
)

func Test_Hello(t *testing.T) {
	lib.Hello()
}

func Test_HelloWorld(t *testing.T) {
	lib.HelloWorld()
}
