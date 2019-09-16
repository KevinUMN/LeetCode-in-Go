package problem0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

// 测试只包含ASCII的string
func Test_Ascii(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "abcabcbb",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "bbbbbbbb",
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: "pwwkew",
			},
			a: ans{
				one: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, lengthOfLongestSubstringAscii(p.one), "Input:%v", p)
	}
}

// 测试只包含ASCII的string
func Test_Unicode(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "bbbbbbbb",
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: "pwwkew",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "时间世界时见件",
			},
			a: ans{
				one: 6,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, lengthOfLongestSubstringUnicode(p.one), "输入:%v", p)
	}
}
