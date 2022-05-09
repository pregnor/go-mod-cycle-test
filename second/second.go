package second

import "github.com/pregnor/go-mod-cycle-test/first"

func Second() {
	first.First()
}
