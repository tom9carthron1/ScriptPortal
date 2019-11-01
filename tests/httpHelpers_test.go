package main

import (
	"github.com/homedepot/ScriptPortal/endpoints/httpHelpers"
	"testing"
)

func TestGetRelativePath(t *testing.T) {
	if httpHelpers.GetRelativePath("abc", "/abc/xyz") != "xyz" {
		t.Error("GetRelativePath failed")
	}
	if httpHelpers.GetRelativePath("abc", "/abc/") != "" {
		t.Error("GetRelativePath failed")
	}
	if httpHelpers.GetRelativePath("abc/xyz", "/abc/xyz/qwerty.html") != "qwerty" {
		t.Error("GetRelativePath Failed")
	}
}
