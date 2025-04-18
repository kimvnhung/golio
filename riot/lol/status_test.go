package lol

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/kimvnhung/golio/internal"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kimvnhung/golio/api"
	"github.com/kimvnhung/golio/internal/mock"
)

func TestStatusClient_Get(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Status
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Status{},
			doer: mock.NewJSONMockDoer(Status{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&StatusClient{c: client}).Get()
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}
