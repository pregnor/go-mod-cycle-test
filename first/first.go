package first

import (
	"github.com/pregnor/go-mod-cycle-test/second"
)

func First() {
	second.Second()
}
