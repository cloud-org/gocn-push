package server

import (
	"strings"
	"testing"
	"time"
)

func TestPushDeer_Send(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "one",
			args: args{
				content: "今天是个好日子",
			},
			wantErr: false,
		},
	}
	p := NewPushDeer("your pushdeer token")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := p.Send(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPushDeer_SendNews(t *testing.T) {
	type args struct {
		content string
	}
	_, contents := NewGocnNew(nil).GetNewsContent(time.Now().Add(-24 * time.Hour))
	content := strings.Join(contents, "")
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "one",
			args: args{
				content: content,
			},
			wantErr: false,
		},
	}
	p := NewPushDeer("your pushdeer token")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := p.Send(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
