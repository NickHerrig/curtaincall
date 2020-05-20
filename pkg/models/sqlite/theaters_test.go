package sqlite

import (
	"reflect"
	"testing"

	"curtaincall.tech/pkg/models"
)

func TestTheaterModelGet(t *testing.T) {

	tests := []struct {
		name        string
        theaterID   int
        wantTheater *models.Theater
        wantError   error
	}{
		{
			name:      "Valid ID",
			theaterID: 1,
			wantTheater: &models.Theater{
				ID:   1,
				Name: "Des Moines Civic Center",
			},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TheaterModel{db}

			theater, err := m.Get(tt.theaterID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if !reflect.DeepEqual(theater, tt.wantTheater) {
				t.Errorf("want %v; got %v", tt.wantTheater, theater)
			}
		})
	}
}
