package sayhello

import (
	"github.com/songjinli8/hello"
	"rsc.io/quote"
)

func Say() string {
	quote.Hello()
	return hello.Hello()
}
