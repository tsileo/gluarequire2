package gluarequire2

import (
	"testing"

	"github.com/yuin/gopher-lua"
)

func TestRequireNoop(t *testing.T) {
	L := lua.NewState()
	defer L.Close()

	NewRequire2Module(nil).SetGlobal(L)

	if err := L.DoString(`
	local testmod = require2('_tests/testmod')
	assert(testmod.return1() == 1)
	`); err != nil {
		panic(err)
	}
}
