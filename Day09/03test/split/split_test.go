package split

import (
	"fmt"
	"os"
	"reflect"
	"runtime/pprof"
	"testing"
)

//进行额外的配置，运行在 主 goroutine 中，可以在调用 m.Run 前后做任何设置Setup和拆卸TearDown
func TestMain(m *testing.M) {

	file, _ := os.Create("./cpu.pprof")
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	fmt.Println("Write Setup code Here....")
	retCode := m.Run() //在这里会执行测试用例
	fmt.Println("Write teardown ode here...")
	os.Exit(retCode)
}

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
		"simple":    {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":  {input: "abcd", sep: "bc", want: []string{"a", "d"}},
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
		"simple":    {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":  {input: "abcd", sep: "bc", want: []string{"a", "d"}},
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

//基准测试
func BenchmarkSplit(b *testing.B) {
	//时间的重置
	b.ResetTimer() //重新计时
	for i := 0; i < b.N; i++ {
		Split("a:fsa:v:sasa", ":")
	}
}

//并行的基准测试
func BenchmarkSplitParallel(b *testing.B) {
	fmt.Println("-------------------并行的基准测试------------------")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("a:fsa:v:sasa", ":")
		}
	})
}
