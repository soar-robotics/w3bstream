package wasm

type ExportsHandler interface {
	Start()
	Malloc()
	Free()
}