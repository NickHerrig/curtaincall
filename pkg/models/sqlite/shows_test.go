package sqlite

import (
	"reflect"
	"testing"

	"curtaincall.tech/pkg/models"
)

func TestShowModelGet(t *testing.T) {

	tests := []struct {
		name        string
        showID      int
        wantShow    *models.Show
        wantError   error
	}{
		{
			name:      "Valid ID",
			showID: 1,
			wantShow: &models.Show{
				ID:      1,
				Name:    "Hamilton",
				Company: "Company A",
			},
			wantError: nil,
		},
		{
			name:      "Valid ID",
			showID: 1,
			wantShow:
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := ShowModel{db}

			show, err := m.Get(tt.showID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if !reflect.DeepEqual(show, tt.wantShow) {
				t.Errorf("want %v; got %v", tt.wantShow, show)
			}
		})
	}
}
