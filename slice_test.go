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

func TestSliceExists(t *testing.T) {
	type args struct {
		data interface{}
		ele  interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1,2,3,4",
			args: args{
				data: []uint64{1, 2, 3, 4, 5},
				ele:  uint64(1),
			},
			want: true,
		},
		{
			name: "a,b,c,d",
			args: args{
				data: []string{"a", "b", "e", "f"},
				ele:  "c",
			},
			want: false,
		},
		{
			name: "map",
			args: args{
				data: []map[string]string{{"a": "b"}},
				ele:  map[string]string{"a": "b"},
			},
			want: true,
		},
		{
			name: "map2",
			args: args{
				data: []map[string]string{{"a": "b"}},
				ele:  map[string]string{"a": "c"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceExists(tt.args.data, tt.args.ele); got != tt.want {
				t.Errorf("SliceExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
