package utils

import (
	"container/list"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	baseStr        = "0123456789abcdefghijkmnpqrstuvwxyABCDEFGHIJKLMNPQRSTUVWXYZ=<>[]{|}"
	baseSuffixChar = ':'
	baseCodeLen    = 10
	baseLen        = int32(len(baseStr))
	baseByte       = []byte(baseStr)
)
var baseMap map[byte]int

func init() {
	rand.Seed(time.Now().UnixNano())

	baseMap = make(map[byte]int)
	for i, v := range baseByte {
		baseMap[v] = i
	}
}

// 用户自定义
func InitBaseCfg(userBase string, userBaseLen int32, suffixChar ...int32) {
	baseStr = userBase
	baseLen = userBaseLen

	if len(suffixChar) > 0 {
		baseSuffixChar = suffixChar[0]
	}
}

func Base2Code(n uint64) string {
	quotient := n
	mod := uint64(0)

	l := list.New()
	for quotient > 0 {
		mod = quotient % uint64(baseLen)
		quotient = quotient / uint64(baseLen)
		l.PushFront(baseByte[int(mod)])
	}

	if l.Len() < baseCodeLen {
		l.PushFront(byte(baseSuffixChar))
	}

	sb := strings.Builder{}
	sb.Grow(baseCodeLen)
	lLen := l.Len()
	for i := 0; i < baseCodeLen; i++ {
		if i < baseCodeLen-lLen {
			rnd := rand.Int31n(baseLen)
			sb.WriteByte(baseByte[rnd])
		} else {
			sb.WriteByte(l.Front().Value.(byte))
			l.Remove(l.Front())
		}
	}

	return sb.String()
}

func BaseCode2Num(str []byte) (res uint64, err error) {
	if str == nil || len(str) == 0 {
		return 0, errors.New("parameter is nil or empty")
	}

	var r uint64 = 0
	for i := len(str) - 1; i >= 0; i-- {
		if int32(str[i]) == baseSuffixChar {
			break
		}

		v, ok := baseMap[str[i]]
		if !ok {
			fmt.Printf("")
			return 0, errors.New("character is not base")
		}

		var b uint64 = 1
		for j := uint64(0); j < r; j++ {
			b *= uint64(baseLen)
		}

		res += b * uint64(v)
		r++
	}
	return res, nil
}
