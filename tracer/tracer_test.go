package tracer

import (
	"context"
	"gopkg.in/go-playground/stats.v1"
	"reflect"
	"testing"
)

func TestNewTracerServer(t *testing.T) {
	type args struct {
		tracerParams TracerParams
	}
	tests := []struct {
		name string
		args args
		want *TracerServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTracerServer(tt.args.tracerParams); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTracerServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTracerServer_Start(t1 *testing.T) {
	type fields struct {
		Ctx    context.Context
		Server *stats.ServerStats
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TracerServer{
				Ctx:    tt.fields.Ctx,
				Server: tt.fields.Server,
			}
			t.Start()
		})
	}
}
