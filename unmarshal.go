package any

import (
	"bytes"
	"strings"
	"sync"
)

// 反序列化对象池
var umPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

func (v *Value) UnmarshalJSON(b []byte) error {
	if b[0] == '"' && b[len(b)-1] == '"' {
		*v = Value(strings.Trim(string(b), `"`))
		return nil
	}
	buf := umPool.Get().(*bytes.Buffer)
	buf.Reset()
	buf.Write(b)
	*v = Value(buf.String())
	// 一定要记得放回
	umPool.Put(buf)
	return nil
}
