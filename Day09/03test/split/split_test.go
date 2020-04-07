package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%#v,got:%#v", want, got)
	}
}

//测试组
func TestArrSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := []test{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
	}

	for _, val := range tests {
		got := Split(val.input, val.sep)
		if !reflect.DeepEqual(got, val.want) {
			t.Errorf("excepted:%#v,got:%#v", val.want, got)
		}
	}

}

//通过 map 完成子测试
func TestZiSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}

	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("name:%s,excepted:%#v,got:%#v\n", name, tc.want, got)
		}
	}
	// t.Run 查看更详细的子测试
	fmt.Printf("-----------------------------------------------\n")
	fmt.Println("t.Run 查看更详细的子测试")
}
func TestRunSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%s,excepted:%#v,got:%#v\n", name, tc.want, got)
			}
		})
	}

}
