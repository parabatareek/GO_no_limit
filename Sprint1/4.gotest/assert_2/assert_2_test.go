package assert_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type fields struct {
	FirstName string
	LastName  string
}

func TestUser_FullName(t *testing.T) {
	tests := []struct {
		name   string
		values fields
		want   string
	}{
		{name: "simple test #1",
			values: fields{
				FirstName: "Anton",
				LastName:  "Palchikov",
			},
			want: "Anton Palchikov"},
		{
			name: "simple long test #1",
			values: fields{
				FirstName: "Pablo Diego KHoze Frantsisko de Paula KHuan Nepomukeno Krispin Krispiano de la Santisima Trinidad Ruiz",
				LastName:  "Picasso",
			},
			want: "Pablo Diego KHoze Frantsisko de Paula KHuan Nepomukeno Krispin Krispiano de la Santisima Trinidad Ruiz Picasso",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := User{
				FirstName: test.values.FirstName,
				LastName:  test.values.LastName,
			}
			assert.Equal(t, test.want, u.FullName())
		})
	}
}
