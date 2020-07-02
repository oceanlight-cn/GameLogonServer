package main

import "testing"

func TestValidateUser(t *testing.T) {
	if !validate_user("Test", "123") {
		t.Error(`Validate user faild.`)
	}
}