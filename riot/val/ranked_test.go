package val

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

func TestChallengesClient_GetLeaderboardByActId(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Leaderboard
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Leaderboard{},
			doer: mock.NewJSONMockDoer(Leaderboard{}, 200),
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
				got, err := (&RankedClient{c: client}).GetLeaderboardByActID("actId", -1, 0)
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}
