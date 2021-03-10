package gee

import (
	"reflect"
	"testing"
)

func TestContext_Param(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		c    *Context
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Param(tt.args.key); got != tt.want {
				t.Errorf("Context.Param() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_param(t *testing.T) {
	type args struct {
		pattern string
		path    string
	}
	c := NewContext(nil, nil)
	tests := []struct {
		name string
		c    *Context
		args args
		want string
		key  string
	}{
		// TODO: Add test cases.
		{
			c: c,
			args: args{
				pattern: "/hello/:name",
				path:    "/hello/wanglei",
			},
			key:  "name",
			want: "wanglei",
		},
		{
			c: c,
			args: args{
				pattern: "/hello/*filename",
				path:    "/hello/css/style.css",
			},
			key:  "filename",
			want: "css/style.css",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.param(tt.args.pattern, tt.args.path)
			if !reflect.DeepEqual(tt.c.Param(tt.key), tt.want) {
				t.Logf("want:%s, but get %s", tt.want, tt.c.Param(tt.key))
			}
		})
	}
}
