// Code generated by wit-bindgen-go. DO NOT EDIT.

package outgoinghandler

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/hayride-dev/bindgen/gen/go/wasi/http/types"
	"unsafe"
)

// ErrorCodeShape is used for storage in variant or result types.
type ErrorCodeShape struct {
	shape [unsafe.Sizeof(types.ErrorCode{})]byte
}

func lower_OptionRequestOptions(v cm.Option[types.RequestOptions]) (f0 uint32, f1 uint32) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1 := cm.Reinterpret[uint32](*some)
		f1 = (uint32)(v1)
	}
	return
}
