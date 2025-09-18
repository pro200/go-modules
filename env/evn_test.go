package env_test

import (
	"testing"

	"github.com/pro200/go-modules/env"
)

func TestEnv(t *testing.T) {
	env.Get("hello")

	//t.Error("Wrong result")
}
