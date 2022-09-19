package talk

import (
	"github.com/mb-14/gomarkov"
)

var (
	Order    = 1
	Chain    = gomarkov.NewChain(Order)
	Messages []string
)
