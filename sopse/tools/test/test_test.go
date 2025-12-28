package test

import "testing"

func TestAssertFile(t *testing.T) {
	orig := TempFile(t, 123)
	AssertFile(t, orig, 123)
}

func TestTempFile(t *testing.T) {
	orig := TempFile(t, 123)
	AssertFile(t, orig, 123)
}
