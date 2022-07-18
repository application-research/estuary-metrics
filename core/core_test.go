package core

import (
	"github.com/whyrusleeping/memo"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	type args struct {
		db      *gorm.DB
		cacherm *memo.Cacher
	}
	tests := []struct {
		name    string
		args    args
		want    *Metrics
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Init(tt.args.db, tt.args.cacherm)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() got = %v, want %v", got, tt.want)
			}
		})
	}
}
