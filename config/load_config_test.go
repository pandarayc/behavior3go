package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *B3ProjectCfg
		wantErr bool
	}{
		{
			name:    "test load project config",
			args:    args{path: "../example/project.json"},
			want:    &B3ProjectCfg{},
			wantErr: false,
		},
		{
			name:    "test load tree config",
			args:    args{path: "../example/tree.json"},
			want:    &B3ProjectCfg{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
