package tft

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/kimvnhung/golio/api"
	"github.com/kimvnhung/golio/internal"
	"github.com/kimvnhung/golio/internal/mock"
)

func TestNewClient(t *testing.T) {
	t.Parallel()
	c := NewClient(internal.NewClient(api.RegionEuropeNorthEast, "key", mock.NewStatusMockDoer(200), log.StandardLogger()))
	if c == nil {
		t.Error("returned nil")
	}
}
