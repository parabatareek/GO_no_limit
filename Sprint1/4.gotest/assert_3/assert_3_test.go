package assert_3

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFamily_AddNew(t *testing.T) {
	type newPerson struct {
		r Relationship
		p Person
	}
	tests := []struct {
		name           string
		existedMembers map[Relationship]Person
		newPerson      newPerson
		wantErr        bool
	}{
		{
			name: "add father",
			existedMembers: map[Relationship]Person{
				Mother: {
					FirstName: "Maria",
					LastName:  "Popova",
					Age:       36,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			wantErr: false,
		},

		{name: "add father #",
			existedMembers: map[Relationship]Person{
				Mother: {
					FirstName: "Marina",
					LastName:  "Cvetaeva",
					Age:       34,
				},
			},
			newPerson: newPerson{
				p: Person{
					FirstName: "Anton",
					LastName:  "Palchikov",
					Age:       22,
				},
				r: Father,
			},
			wantErr: false,
		},

		{
			name: "catch error",
			existedMembers: map[Relationship]Person{
				Father: {
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Ken",
					LastName:  "Gymsohn",
					Age:       32,
				},
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := &Family{
				Members: test.existedMembers,
			}
			err := f.AddNew(test.newPerson.r, test.newPerson.p)
			if !test.wantErr {
				require.NoError(t, err)

				assert.Contains(t, f.Members, test.newPerson.r)
				return
			}

			assert.Error(t, err)
		})
	}
}
