package state

import (
	"testing"
)

func TestDaemonCreate(t *testing.T) {
	if Running.String() != "Running" {
		t.Fatal("Running state should be 'Running'")
	}
	if Error.String() != "Error" {
		t.Fatal("Error state should be 'Error'")
	}
}
