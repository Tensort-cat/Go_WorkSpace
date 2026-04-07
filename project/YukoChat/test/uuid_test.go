package test

import (
	"fmt"
	"testing"
	"yuko_chat/pkg/util"
)

func TestGenerateUUID(t *testing.T) {
	uuid := util.GenUserUUID()
	fmt.Printf("生成的UUID: %s\n", uuid)
}
