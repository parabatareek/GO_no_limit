package main

import "testing"

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
		{
			"add father",
			map[Relationship]Person{
				Mother: {
					FirstName: "Maria",
					LastName:  "Sharapova",
					Age:       34,
				},
			},
			newPerson{
				p: Person{
					"Anton",
					"Kuztetsov",
					34,
				},
				r: Father,
			},
			false,
		},
		{
			name: "catch error",
			existedMembers: map[Relationship]Person{
				Father: {
					"Sergey",
					"Smornoff",
					34,
				},
			},
			newPerson: newPerson{
				p: Person{"Sergey", "Smirnoff", 45},
				r: Father,
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := &Family{Members: test.existedMembers}
			if err := f.AddNew(test.newPerson.r, test.newPerson.p); (err != nil) != test.wantErr {
				t.Errorf("AddNew() error: %v, wantErr: %v", err, test.wantErr)
			}
		})
	}
}
