package httptest_1

import (
	"net/http"
	"reflect"
	"testing"
)

func TestUserViewHandler(t *testing.T) {
	typ
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		{
			"check Content-Type",
			args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserViewHandler(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserViewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
