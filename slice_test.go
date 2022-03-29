package gtools

import (
	"reflect"
	"testing"
)

func TestSliceUNIQ(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "1,1,2",
			args: args{
				data: []uint64{1, 1, 2},
			},
			want: []uint64{1, 2},
		},
		{
			name: "a,b,c,d",
			args: args{
				data: []string{"a", "b", "c", "d"},
			},
			want: []string{"a", "b", "c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceUNIQ(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceUNIQ() = %v, want %v", got, tt.want)
			}
		})
	}
}
