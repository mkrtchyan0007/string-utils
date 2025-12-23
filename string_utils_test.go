package stringutils

import "testing"

func TestReverse(t *testing.T) {
	if Reverse("hello") != "olleh" {
		t.Error("Reverse failed")
	}
}

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("level") {
		t.Error("IsPalindrome failed for 'level'")
	}
	if IsPalindrome("hello") {
		t.Error("IsPalindrome failed for 'hello'")
	}
}
