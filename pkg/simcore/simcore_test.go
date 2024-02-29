package simcore

import (
	"reflect"
	"testing"
)

func Test_wrapFloat(t *testing.T) {
	type args struct {
		num float32
		min float32
		max float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			"wraps num less than min to max",
			args{-0.5, 0.0, 1.0},
			1.0,
		},
		{
			"wraps num greater than max to min",
			args{1.5, 0.0, 1.0},
			0.0,
		},
		{
			"num stays same if between min and max",
			args{0.5, 0.0, 1.0},
			0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrapFloat(tt.args.num, tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
