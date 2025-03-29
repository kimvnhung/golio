package val

import (
	"testing"

	"github.com/kimvnhung/golio/internal"
	log "github.com/sirupsen/logrus"

	"github.com/kimvnhung/golio/api"
	"github.com/kimvnhung/golio/internal/mock"
)

func TestNewClient(t *testing.T) {
	c := NewClient(internal.NewClient(api.RegionBrasil, "key", mock.NewStatusMockDoer(200), log.StandardLogger()))
	if c == nil {
		t.Error("returned nil")
	}
}
