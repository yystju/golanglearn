package main_test

import (
	"testing"
)

func TestMain_shouldnotrun(t *testing.T) {
	t.Error("Should not run...")
}

func TestMain_whatever(t *testing.T) {
	t.Skip("This is skipped...")
}
