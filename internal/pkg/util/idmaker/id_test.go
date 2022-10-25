package idmaker

import "testing"

func TestGenId(t *testing.T) {
	serial := GenId()
	t.Logf("GenId():%s\n", serial)
	t.Logf("len:%d\n", len(serial))
}
