package engine

import (
	"log"
	"reflect"
	"testing"
)

func test_node() *Node {
	n := NewNode()
	n.Insert("/home")
	n.Insert("/myhome/say")
	n.Insert("/home/:name")
	n.Insert("/file/*filepath")
	return n
}

func TestNode_Insert(t *testing.T) {
	type args struct {
		path string
	}
	n := NewNode()
	tests := []struct {
		name string
		n    *Node
		args args
	}{
		// TODO: Add test cases.
		{
			n: n,
			args: args{
				path: "/home",
			},
		},
		{
			n: n,
			args: args{
				path: "/myhome/say",
			},
		},
		{
			n: n,
			args: args{
				path: "/home/:name",
			},
		},
		{
			n: n,
			args: args{
				path: "/file/*filepath",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.Insert(tt.args.path)

		})
	}
	log.Println(n)

}

func Test_isMatchNode(t *testing.T) {
	type args struct {
		path string
		node *Node
	}
	n1 := NewNode()
	n1.part = ":name"

	n2 := NewNode()
	n2.part = "*filepath"

	n3 := NewNode()
	n3.part = "/myhome/say"

	tests := []struct {
		name  string
		args  args
		want  map[string]string
		want1 bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				path: "sz",
				node: n1,
			},
			want: map[string]string{
				"name": "sz",
			},
			want1: true,
		},
		{
			args: args{
				path: "index/style.css",
				node: n2,
			},
			want: map[string]string{
				"filepath": "index/style.css",
			},
			want1: true,
		},
		{
			name: "正常情况匹配节点",
			args: args{
				path: "/hello/say",
				node: n3,
			},
			want:  map[string]string{},
			want1: false,
		},
		{
			name: "直接返回节点",
			args: args{
				path: "/",
				node: n3,
			},
			want:  map[string]string{},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isMatchNode(tt.args.path, tt.args.node)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isMatchNode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isMatchNode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNode_Search(t *testing.T) {
	type args struct {
		path string
	}
	n := test_node()
	tests := []struct {
		name string
		n    *Node
		args args
		want *NodeInfo
	}{
		// TODO: Add test cases.
		{
			name: "hello路由不在，直接返回nil",
			n:    n,
			args: args{
				path: "/hello/say",
			},
			want: nil,
		},
		{
			name: "",
			n:    n,
			args: args{
				path: "/myhome/say",
			},
			want: &NodeInfo{
				Param: map[string]string{},
				Node:  n.children[1].children[0],
			},
		},
		{
			name: "",
			n:    n,
			args: args{
				path: "/home/shenz",
			},
			want: &NodeInfo{
				Param: map[string]string{
					"name": "shenz",
				},
				Node: n.children[0].children[0],
			},
		},
		{
			name: "",
			n:    n,
			args: args{
				path: "/file/index/style.css",
			},
			want: &NodeInfo{
				Param: map[string]string{
					"filepath": "index/style.css",
				},
				Node: n.children[2].children[0],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Search(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
