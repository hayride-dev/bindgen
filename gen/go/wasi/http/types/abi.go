// Code generated by wit-bindgen-go. DO NOT EDIT.

package types

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	monotonicclock "github.com/hayride-dev/bindgen/gen/go/wasi/clocks/monotonic-clock"
	"unsafe"
)

// OptionFieldSizePayloadShape is used for storage in variant or result types.
type OptionFieldSizePayloadShape struct {
	shape [unsafe.Sizeof(cm.Option[FieldSizePayload]{})]byte
}

func lower_OptionString(v cm.Option[string]) (f0 uint32, f1 *uint8, f2 uint32) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1, v2 := cm.LowerString(*some)
		f1 = (*uint8)(v1)
		f2 = (uint32)(v2)
	}
	return
}

func lower_Method(v Method) (f0 uint32, f1 *uint8, f2 uint32) {
	f0 = (uint32)(v.Tag())
	switch f0 {
	case 9: // other
		v1, v2 := cm.LowerString(*v.Other())
		f1 = (*uint8)(v1)
		f2 = (uint32)(v2)
	}
	return
}

func lower_Scheme(v Scheme) (f0 uint32, f1 *uint8, f2 uint32) {
	f0 = (uint32)(v.Tag())
	switch f0 {
	case 2: // other
		v1, v2 := cm.LowerString(*v.Other())
		f1 = (*uint8)(v1)
		f2 = (uint32)(v2)
	}
	return
}

func lower_OptionScheme(v cm.Option[Scheme]) (f0 uint32, f1 uint32, f2 *uint8, f3 uint32) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1, v2, v3 := lower_Scheme(*some)
		f1 = (uint32)(v1)
		f2 = (*uint8)(v2)
		f3 = (uint32)(v3)
	}
	return
}

func lower_OptionDuration(v cm.Option[monotonicclock.Duration]) (f0 uint32, f1 uint64) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1 := (uint64)(*some)
		f1 = (uint64)(v1)
	}
	return
}

// ErrorCodeShape is used for storage in variant or result types.
type ErrorCodeShape struct {
	shape [unsafe.Sizeof(ErrorCode{})]byte
}

func lower_OptionU16(v cm.Option[uint16]) (f0 uint32, f1 uint32) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1 := (uint32)(*some)
		f1 = (uint32)(v1)
	}
	return
}

func lower_DNSErrorPayload(v DNSErrorPayload) (f0 uint32, f1 *uint8, f2 uint32, f3 uint32, f4 uint32) {
	f0, f1, f2 = lower_OptionString(v.Rcode)
	f3, f4 = lower_OptionU16(v.InfoCode)
	return
}

func lower_OptionU8(v cm.Option[uint8]) (f0 uint32, f1 uint32) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1 := (uint32)(*some)
		f1 = (uint32)(v1)
	}
	return
}

func lower_TLSAlertReceivedPayload(v TLSAlertReceivedPayload) (f0 uint32, f1 uint32, f2 uint32, f3 *uint8, f4 uint32) {
	f0, f1 = lower_OptionU8(v.AlertID)
	f2, f3, f4 = lower_OptionString(v.AlertMessage)
	return
}

func lower_OptionU64(v cm.Option[uint64]) (f0 uint32, f1 uint64) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1 := (uint64)(*some)
		f1 = (uint64)(v1)
	}
	return
}

func lower_OptionU32(v cm.Option[uint32]) (f0 uint32, f1 uint32) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1 := (uint32)(*some)
		f1 = (uint32)(v1)
	}
	return
}

func lower_FieldSizePayload(v FieldSizePayload) (f0 uint32, f1 *uint8, f2 uint32, f3 uint32, f4 uint32) {
	f0, f1, f2 = lower_OptionString(v.FieldName)
	f3, f4 = lower_OptionU32(v.FieldSize)
	return
}

func lower_OptionFieldSizePayload(v cm.Option[FieldSizePayload]) (f0 uint32, f1 uint32, f2 *uint8, f3 uint32, f4 uint32, f5 uint32) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1, v2, v3, v4, v5 := lower_FieldSizePayload(*some)
		f1 = (uint32)(v1)
		f2 = (*uint8)(v2)
		f3 = (uint32)(v3)
		f4 = (uint32)(v4)
		f5 = (uint32)(v5)
	}
	return
}

func lower_ErrorCode(v ErrorCode) (f0 uint32, f1 uint32, f2 uint64, f3 uint32, f4 uint32, f5 uint32, f6 uint32) {
	f0 = (uint32)(v.Tag())
	switch f0 {
	case 1: // DNS-error
		v1, v2, v3, v4, v5 := lower_DNSErrorPayload(*v.DNSError())
		f1 = (uint32)(v1)
		f2 = cm.PointerToU64(v2)
		f3 = (uint32)(v3)
		f4 = (uint32)(v4)
		f5 = (uint32)(v5)
	case 14: // TLS-alert-received
		v1, v2, v3, v4, v5 := lower_TLSAlertReceivedPayload(*v.TLSAlertReceived())
		f1 = (uint32)(v1)
		f2 = (uint64)(v2)
		f3 = (uint32)(v3)
		f4 = cm.PointerToU32(v4)
		f5 = (uint32)(v5)
	case 17: // HTTP-request-body-size
		v1, v2 := lower_OptionU64(*v.HTTPRequestBodySize())
		f1 = (uint32)(v1)
		f2 = (uint64)(v2)
	case 21: // HTTP-request-header-section-size
		v1, v2 := lower_OptionU32(*v.HTTPRequestHeaderSectionSize())
		f1 = (uint32)(v1)
		f2 = (uint64)(v2)
	case 22: // HTTP-request-header-size
		v1, v2, v3, v4, v5, v6 := lower_OptionFieldSizePayload(*v.HTTPRequestHeaderSize())
		f1 = (uint32)(v1)
		f2 = (uint64)(v2)
		f3 = cm.PointerToU32(v3)
		f4 = (uint32)(v4)
		f5 = (uint32)(v5)
		f6 = (uint32)(v6)
	case 23: // HTTP-request-trailer-section-size
		v1, v2 := lower_OptionU32(*v.HTTPRequestTrailerSectionSize())
		f1 = (uint32)(v1)
		f2 = (uint64)(v2)
	case 24: // HTTP-request-trailer-size
		v1, v2, v3, v4, v5 := lower_FieldSizePayload(*v.HTTPRequestTrailerSize())
		f1 = (uint32)(v1)
		f2 = cm.PointerToU64(v2)
		f3 = (uint32)(v3)
		f4 = (uint32)(v4)
		f5 = (uint32)(v5)
	case 26: // HTTP-response-header-section-size
		v1, v2 := lower_OptionU32(*v.HTTPResponseHeaderSectionSize())
		f1 = (uint32)(v1)
		f2 = (uint64)(v2)
	case 27: // HTTP-response-header-size
		v1, v2, v3, v4, v5 := lower_FieldSizePayload(*v.HTTPResponseHeaderSize())
		f1 = (uint32)(v1)
		f2 = cm.PointerToU64(v2)
		f3 = (uint32)(v3)
		f4 = (uint32)(v4)
		f5 = (uint32)(v5)
	case 28: // HTTP-response-body-size
		v1, v2 := lower_OptionU64(*v.HTTPResponseBodySize())
		f1 = (uint32)(v1)
		f2 = (uint64)(v2)
	case 29: // HTTP-response-trailer-section-size
		v1, v2 := lower_OptionU32(*v.HTTPResponseTrailerSectionSize())
		f1 = (uint32)(v1)
		f2 = (uint64)(v2)
	case 30: // HTTP-response-trailer-size
		v1, v2, v3, v4, v5 := lower_FieldSizePayload(*v.HTTPResponseTrailerSize())
		f1 = (uint32)(v1)
		f2 = cm.PointerToU64(v2)
		f3 = (uint32)(v3)
		f4 = (uint32)(v4)
		f5 = (uint32)(v5)
	case 31: // HTTP-response-transfer-coding
		v1, v2, v3 := lower_OptionString(*v.HTTPResponseTransferCoding())
		f1 = (uint32)(v1)
		f2 = cm.PointerToU64(v2)
		f3 = (uint32)(v3)
	case 32: // HTTP-response-content-coding
		v1, v2, v3 := lower_OptionString(*v.HTTPResponseContentCoding())
		f1 = (uint32)(v1)
		f2 = cm.PointerToU64(v2)
		f3 = (uint32)(v3)
	case 38: // internal-error
		v1, v2, v3 := lower_OptionString(*v.InternalError())
		f1 = (uint32)(v1)
		f2 = cm.PointerToU64(v2)
		f3 = (uint32)(v3)
	}
	return
}

func lower_ResultOutgoingResponseErrorCode(v cm.Result[ErrorCodeShape, OutgoingResponse, ErrorCode]) (f0 uint32, f1 uint32, f2 uint32, f3 uint64, f4 uint32, f5 uint32, f6 uint32, f7 uint32) {
	if v.IsOK() {
		v1 := cm.Reinterpret[uint32](*v.OK())
		f1 = (uint32)(v1)
	} else {
		f0 = 1
		v1, v2, v3, v4, v5, v6, v7 := lower_ErrorCode(*v.Err())
		f1 = (uint32)(v1)
		f2 = (uint32)(v2)
		f3 = (uint64)(v3)
		f4 = (uint32)(v4)
		f5 = (uint32)(v5)
		f6 = (uint32)(v6)
		f7 = (uint32)(v7)
	}
	return
}

// ResultOptionTrailersErrorCodeShape is used for storage in variant or result types.
type ResultOptionTrailersErrorCodeShape struct {
	shape [unsafe.Sizeof(cm.Result[ErrorCodeShape, cm.Option[Fields], ErrorCode]{})]byte
}

func lower_OptionTrailers(v cm.Option[Fields]) (f0 uint32, f1 uint32) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1 := cm.Reinterpret[uint32](*some)
		f1 = (uint32)(v1)
	}
	return
}

// ResultIncomingResponseErrorCodeShape is used for storage in variant or result types.
type ResultIncomingResponseErrorCodeShape struct {
	shape [unsafe.Sizeof(cm.Result[ErrorCodeShape, IncomingResponse, ErrorCode]{})]byte
}
