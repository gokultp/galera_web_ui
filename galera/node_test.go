package galera_test

import (
	"testing"

	"github.com/gokultp/galera_web_ui/galera"
)

func TestGetNodes(t *testing.T) {
	_, err := galera.GetNodes(nil)

	if err == nil {
		t.Error("Expected an error to return while connection is nil")
	} else if err.Error() != galera.ErrNoClient {
		t.Error("Expected an error to return while connection is nil")
	}
}
