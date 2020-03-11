package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		fmt.Fprint(os.Stdout, "%s", buf)
	}
}

func main() {
	flag.Parse()
	//没有默认参数，从标准输入读取内容
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdout))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading fromr %s failed,err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}

}
