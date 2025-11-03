package pkg

import (
	"os"
)

const (
	SharedSecret = "_sharedSecret"
)

var (
	WorkDir, _ = os.Getwd()
)
