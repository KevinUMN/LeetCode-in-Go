package Problem0722

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	source []string
	ans    []string
}{

	{
		[]string{
			"void func(int k) {",
			"// this function does nothing /*",
			"   k = k*2/4;",
			"   k = k/2;*/",
			"}",
		},
		[]string{
			"void func(int k) {",
			"   k = k*2/4;",
			"   k = k/2;*/",
			"}",
		},
	},

	{
		[]string{
			"/*Test program */",
			"int main()",
			"{ ",
			"  // variable declaration ",
			"int a, b, c;",
			"/* This is a test",
			"   multiline  ",
			"   comment for ",
			"   testing */",
			"a = b + c;",
			"}",
		},
		[]string{
			"int main()",
			"{ ",
			"  ",
			"int a, b, c;",
			"a = b + c;",
			"}",
		},
	},

	{
		[]string{
			"a/*comment",
			"line",
			"more_comment*/b",
		},
		[]string{
			"ab",
		},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, removeComments(tc.source), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeComments(tc.source)
		}
	}
}
