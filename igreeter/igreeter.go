package igreeter

import build "greetertest/src/greeter/lib"

type IGreeter interface {
	Initialize()
	Greet()
	SayHi(string)
	Version() string
	GetBinary() []byte
	DeInitialize()
	BuildInfo() build.BuildInfo
}
