package uuid

import (
	"testing"

	"github.com/bxsmart/bxcore/utils/goSnowFlake"
)

func TestIdGen(t *testing.T) {
	id := IdGen()
	println(id)
	time, ts, cid, wid, seq := goSnowFlake.ParseId(id)
	t.Log(time, ts, cid, wid, seq)
}

func TestIdGenEx(t *testing.T) {
	id := IdGenEx(1, 1)
	println(id)
	time, ts, cid, wid, seq := goSnowFlake.ParseId(id)
	t.Log(time, ts, cid, wid, seq)
}

func TestIdGenDef(t *testing.T) {
	workerId := 2 << 6
	id := IdGenDef(int64(workerId))
	println(id)
	time, ts, cid, wid, seq := goSnowFlake.ParseId(id)
	t.Log(time, ts, cid, wid, seq)
}
