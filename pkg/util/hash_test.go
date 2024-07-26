package util

import (
	"fmt"
	"testing"

	openfunction "github.com/openfunction/apis/core/v1beta2"
)

func TestName(t *testing.T) {

	x := openfunction.ServingSpec{}
	fmt.Println(Hash(x))
}
