package test

import (
	"CurrencyExchangeApp/util"
	"fmt"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	token, err := util.GenerateJWT("123456")
	if err != nil {
		t.Fatalf("GenerateJWT returned error: %v", err)
	}
	fmt.Println(token)
}
