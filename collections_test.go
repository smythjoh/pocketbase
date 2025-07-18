package pocketbase

import (
	"testing"

	"github.com/smythjoh/pocketbase/migrations"
	"github.com/stretchr/testify/assert"
)

func TestCollections_List(t *testing.T) {

	defaultClient := NewClient(defaultURL, WithAdminEmailPassword(migrations.AdminEmailPassword, migrations.AdminEmailPassword))
	collections := NewCollections(defaultClient)

	tests := []struct {
		name       string
		client     *Client
		params     ParamsList
		wantResult bool
		wantErr    bool
	}{
		{
			name:       "List with no params",
			client:     defaultClient,
			wantErr:    false,
			wantResult: true,
		},
		{
			name:   "List no results - query",
			client: defaultClient,
			params: ParamsList{
				Page: 1,
				Size: 1,
			},
			wantErr:    false,
			wantResult: true,
		},
		{
			name:   "List no results - invalid query",
			client: defaultClient,
			params: ParamsList{
				Page: 2,
				Size: 3,
			},
			wantErr:    false,
			wantResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := collections.List(tt.params)
			assert.Equal(t, tt.wantErr, err != nil, err)
			assert.Equal(t, tt.wantResult, got.TotalItems > 0)
		})
	}
}
