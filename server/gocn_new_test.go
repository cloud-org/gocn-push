package server

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func TestGocnNew_GetNewsContent(t *testing.T) {
	type args struct {
		publishTime time.Time
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "one",
			args: args{publishTime: time.Now().Add(-24 * time.Hour)},
			want: nil,
		},
	}
	g := GocnNew{Client: resty.New()}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := g.GetNewsContent(tt.args.publishTime)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNewsContent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
