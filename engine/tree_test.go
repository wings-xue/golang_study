package engine

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	url1 = "/index/baby"
	url2 = "/file/css/style.css"
	url3 = "/index/:name/say"
	url4 = "/file/*filepath"
)

var n = NewRoot()

func TestGoodNode_insert(t *testing.T) {
	type args struct {
		pattern string
		paths   []string
		height  int
	}

	tests := []struct {
		name string
		n    *GoodNode
		args args
	}{
		// TODO: Add test cases.
		{
			n: n,
			args: args{
				pattern: url1,
				paths:   parsePath(url1),
				height:  0,
			},
		},
		{
			n: n,
			args: args{
				pattern: url2,
				paths:   parsePath(url2),
				height:  0,
			},
		},
		{
			n: n,
			args: args{
				pattern: url3,
				paths:   parsePath(url3),
				height:  0,
			},
		},
		{
			n: n,
			args: args{
				pattern: url4,
				paths:   parsePath(url4),
				height:  0,
			},
		},
	}

	searchTests := []struct {
		name    string
		pattern string
		want    string
	}{
		{
			pattern: "/file/app/style.css",
			want:    "/file/*filepath",
		},
		{
			pattern: "/say",
			want:    "",
		},
		{
			pattern: "/",
			want:    "/",
		},
		{
			pattern: "/index/sz/say",
			want:    "/index/:name/say",
		},
		{
			pattern: "/index/body",
			want:    "/index/body",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.insert(tt.args.pattern, tt.args.paths, tt.args.height)

		})
	}
	for _, stt := range searchTests {
		t.Run("", func(t *testing.T) {
			out := n.search(stt.pattern, parsePath(stt.pattern), 0)
			if out != nil {
				if !reflect.DeepEqual(out.pattern, stt.want) {
					t.Logf("want:%s, search rst: %s\n", stt.want, out.pattern)
				}
			} else {
				if stt.want != "" {
					t.Logf("want:%s, search rst: %s\n", stt.want, "nil")
				}
			}

		})

	}
	fmt.Println(n)
}
