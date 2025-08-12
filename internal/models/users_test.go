package models

import (
	"testing"

	"github.com/print-agent/snippet-box/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	testCases := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			name:   "ValidID",
			userID: 1,
			want:   true,
		},
		{
			name:   "ZeroID",
			userID: 0,
			want:   false,
		},
		{
			name:   "Non-existentID",
			userID: 2,
			want:   false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}

			exists, err := m.Exists(tC.userID)

			assert.Equal(t, exists, tC.want)
			assert.NilError(t, err)
		})
	}
}
