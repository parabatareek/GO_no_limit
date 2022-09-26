package main

import "testing"

type fields struct {
	FirstName string
	LastName  string
}

func TestUser_FullName(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Testmain simple test #1",
			fields: fields{
				FirstName: "Misha",
				LastName:  "Popov",
			},
			want: "Misha Popov"},
		{
			name: "Testmain simple test #2",
			fields: fields{
				FirstName: "Andrey",
				LastName:  "Markov",
			},
			want: "Andrey Markov",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := User{
				FirstName: test.fields.FirstName,
				LastName:  test.fields.LastName,
			}
			if got := u.FullName(); got != test.want {
				t.Errorf("Result:%v. Want: %v", got, test.want)
			}
		})
	}
}
