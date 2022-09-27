package assert_3

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFamily_AddNew(t *testing.T) {
	type newPerson struct {
		p Person
		r Relationship
	}

	tests := []struct {
		name           string
		existedMembers map[Relationship]Person
		newPerson      newPerson
		wantErr        bool
	}{
		{name: "add father",
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
		{name: "catch error",
			existedMembers: map[Relationship]Person{
				Father: {
					FirstName: "Anton",
					LastName:  "Palchikov",
					Age:       22,
				},
			},
			newPerson: newPerson{
				p: Person{
					FirstName: "Bogdan",
					LastName:  "Hmeknitskiy",
					Age:       25,
				},
				r: Father,
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
			}

			assert.Error(t, err)
		})
	}
}
