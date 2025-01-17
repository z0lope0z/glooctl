// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: route_extensions.proto

/*
Package routing is a generated protocol buffer package.

It is generated from these files:
	route_extensions.proto

It has these top-level messages:
	RouteExtensions
	HeaderValue
	CorsPolicy
*/
package routing

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// RouteExtensions should be placed in the route.extensions field
// RouteExtensions extend the behavior of a regular route in gloo (within a virtual service)
type RouteExtensions struct {
	// These headers will be added to the request before it is sent to the upstream
	AddRequestHeaders []*HeaderValue `protobuf:"bytes,1,rep,name=add_request_headers,json=addRequestHeaders" json:"add_request_headers,omitempty"`
	// These headers will be added to the response before it is returned to the downstream
	AddResponseHeaders []*HeaderValue `protobuf:"bytes,2,rep,name=add_response_headers,json=addResponseHeaders" json:"add_response_headers,omitempty"`
	// These headers will be removed from the request before it is sent to the upstream
	RemoveResponseHeaders []string `protobuf:"bytes,3,rep,name=remove_response_headers,json=removeResponseHeaders" json:"remove_response_headers,omitempty"`
	// The maximum number of retries to attempt for requests that get a 5xx response
	MaxRetries uint32 `protobuf:"varint,4,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// If set, time out requests on this route. If not set, this will default to the connection timeout on the upstream
	Timeout *time.Duration `protobuf:"bytes,5,opt,name=timeout,stdduration" json:"timeout,omitempty"`
	// Rewrite the host header of the request to this value, if set
	HostRewrite string `protobuf:"bytes,6,opt,name=host_rewrite,json=hostRewrite,proto3" json:"host_rewrite,omitempty"`
	// Configure Cross-Origin Resource Sharing requests
	Cors *CorsPolicy `protobuf:"bytes,7,opt,name=cors" json:"cors,omitempty"`
}

func (m *RouteExtensions) Reset()                    { *m = RouteExtensions{} }
func (m *RouteExtensions) String() string            { return proto.CompactTextString(m) }
func (*RouteExtensions) ProtoMessage()               {}
func (*RouteExtensions) Descriptor() ([]byte, []int) { return fileDescriptorRouteExtensions, []int{0} }

func (m *RouteExtensions) GetAddRequestHeaders() []*HeaderValue {
	if m != nil {
		return m.AddRequestHeaders
	}
	return nil
}

func (m *RouteExtensions) GetAddResponseHeaders() []*HeaderValue {
	if m != nil {
		return m.AddResponseHeaders
	}
	return nil
}

func (m *RouteExtensions) GetRemoveResponseHeaders() []string {
	if m != nil {
		return m.RemoveResponseHeaders
	}
	return nil
}

func (m *RouteExtensions) GetMaxRetries() uint32 {
	if m != nil {
		return m.MaxRetries
	}
	return 0
}

func (m *RouteExtensions) GetTimeout() *time.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RouteExtensions) GetHostRewrite() string {
	if m != nil {
		return m.HostRewrite
	}
	return ""
}

func (m *RouteExtensions) GetCors() *CorsPolicy {
	if m != nil {
		return m.Cors
	}
	return nil
}

// Header name/value pair
type HeaderValue struct {
	// Header name
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Header value
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Should this value be appended?
	Append bool `protobuf:"varint,3,opt,name=append,proto3" json:"append,omitempty"`
}

func (m *HeaderValue) Reset()                    { *m = HeaderValue{} }
func (m *HeaderValue) String() string            { return proto.CompactTextString(m) }
func (*HeaderValue) ProtoMessage()               {}
func (*HeaderValue) Descriptor() ([]byte, []int) { return fileDescriptorRouteExtensions, []int{1} }

func (m *HeaderValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HeaderValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *HeaderValue) GetAppend() bool {
	if m != nil {
		return m.Append
	}
	return false
}

// Configuration for Cross-Origin Resource Sharing requests
type CorsPolicy struct {
	// Specifies the origins that will be allowed to do CORS requests.
	//
	// An origin is allowed if either allow_origin matches.
	AllowOrigin []string `protobuf:"bytes,1,rep,name=allow_origin,json=allowOrigin" json:"allow_origin,omitempty"`
	// Specifies the content for the *access-control-allow-methods* header.
	AllowMethods string `protobuf:"bytes,2,opt,name=allow_methods,json=allowMethods,proto3" json:"allow_methods,omitempty"`
	// Specifies the content for the *access-control-allow-headers* header.
	AllowHeaders string `protobuf:"bytes,3,opt,name=allow_headers,json=allowHeaders,proto3" json:"allow_headers,omitempty"`
	// Specifies the content for the *access-control-expose-headers* header.
	ExposeHeaders string `protobuf:"bytes,4,opt,name=expose_headers,json=exposeHeaders,proto3" json:"expose_headers,omitempty"`
	// Specifies the content for the *access-control-max-age* header.
	MaxAge *time.Duration `protobuf:"bytes,5,opt,name=max_age,json=maxAge,stdduration" json:"max_age,omitempty"`
	// Specifies whether the resource allows credentials.
	AllowCredentials bool `protobuf:"varint,6,opt,name=allow_credentials,json=allowCredentials,proto3" json:"allow_credentials,omitempty"`
}

func (m *CorsPolicy) Reset()                    { *m = CorsPolicy{} }
func (m *CorsPolicy) String() string            { return proto.CompactTextString(m) }
func (*CorsPolicy) ProtoMessage()               {}
func (*CorsPolicy) Descriptor() ([]byte, []int) { return fileDescriptorRouteExtensions, []int{2} }

func (m *CorsPolicy) GetAllowOrigin() []string {
	if m != nil {
		return m.AllowOrigin
	}
	return nil
}

func (m *CorsPolicy) GetAllowMethods() string {
	if m != nil {
		return m.AllowMethods
	}
	return ""
}

func (m *CorsPolicy) GetAllowHeaders() string {
	if m != nil {
		return m.AllowHeaders
	}
	return ""
}

func (m *CorsPolicy) GetExposeHeaders() string {
	if m != nil {
		return m.ExposeHeaders
	}
	return ""
}

func (m *CorsPolicy) GetMaxAge() *time.Duration {
	if m != nil {
		return m.MaxAge
	}
	return nil
}

func (m *CorsPolicy) GetAllowCredentials() bool {
	if m != nil {
		return m.AllowCredentials
	}
	return false
}

func init() {
	proto.RegisterType((*RouteExtensions)(nil), "gloo.api.v1.RouteExtensions")
	proto.RegisterType((*HeaderValue)(nil), "gloo.api.v1.HeaderValue")
	proto.RegisterType((*CorsPolicy)(nil), "gloo.api.v1.CorsPolicy")
}
func (this *RouteExtensions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteExtensions)
	if !ok {
		that2, ok := that.(RouteExtensions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.AddRequestHeaders) != len(that1.AddRequestHeaders) {
		return false
	}
	for i := range this.AddRequestHeaders {
		if !this.AddRequestHeaders[i].Equal(that1.AddRequestHeaders[i]) {
			return false
		}
	}
	if len(this.AddResponseHeaders) != len(that1.AddResponseHeaders) {
		return false
	}
	for i := range this.AddResponseHeaders {
		if !this.AddResponseHeaders[i].Equal(that1.AddResponseHeaders[i]) {
			return false
		}
	}
	if len(this.RemoveResponseHeaders) != len(that1.RemoveResponseHeaders) {
		return false
	}
	for i := range this.RemoveResponseHeaders {
		if this.RemoveResponseHeaders[i] != that1.RemoveResponseHeaders[i] {
			return false
		}
	}
	if this.MaxRetries != that1.MaxRetries {
		return false
	}
	if this.Timeout != nil && that1.Timeout != nil {
		if *this.Timeout != *that1.Timeout {
			return false
		}
	} else if this.Timeout != nil {
		return false
	} else if that1.Timeout != nil {
		return false
	}
	if this.HostRewrite != that1.HostRewrite {
		return false
	}
	if !this.Cors.Equal(that1.Cors) {
		return false
	}
	return true
}
func (this *HeaderValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderValue)
	if !ok {
		that2, ok := that.(HeaderValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Append != that1.Append {
		return false
	}
	return true
}
func (this *CorsPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CorsPolicy)
	if !ok {
		that2, ok := that.(CorsPolicy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.AllowOrigin) != len(that1.AllowOrigin) {
		return false
	}
	for i := range this.AllowOrigin {
		if this.AllowOrigin[i] != that1.AllowOrigin[i] {
			return false
		}
	}
	if this.AllowMethods != that1.AllowMethods {
		return false
	}
	if this.AllowHeaders != that1.AllowHeaders {
		return false
	}
	if this.ExposeHeaders != that1.ExposeHeaders {
		return false
	}
	if this.MaxAge != nil && that1.MaxAge != nil {
		if *this.MaxAge != *that1.MaxAge {
			return false
		}
	} else if this.MaxAge != nil {
		return false
	} else if that1.MaxAge != nil {
		return false
	}
	if this.AllowCredentials != that1.AllowCredentials {
		return false
	}
	return true
}

func init() { proto.RegisterFile("route_extensions.proto", fileDescriptorRouteExtensions) }

var fileDescriptorRouteExtensions = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0xa6, 0x6b, 0x57, 0x87, 0xc2, 0x66, 0xca, 0x16, 0x76, 0xb1, 0x85, 0x22, 0xa4,
	0x48, 0x88, 0x44, 0x0c, 0x81, 0xe0, 0x92, 0x0d, 0xa4, 0x09, 0x69, 0x02, 0xf9, 0x82, 0x0b, 0x6e,
	0x22, 0xb7, 0x39, 0xb8, 0xd6, 0x92, 0x9c, 0x60, 0x3b, 0x5d, 0xf7, 0x04, 0x5c, 0xf0, 0x02, 0x3c,
	0x02, 0x6f, 0x85, 0xc4, 0x93, 0xa0, 0xd8, 0xed, 0x5a, 0xe0, 0x82, 0xdd, 0xd9, 0xdf, 0xf9, 0xcf,
	0xef, 0x93, 0xff, 0x28, 0x64, 0x4f, 0x61, 0x63, 0x20, 0x83, 0x85, 0x81, 0x4a, 0x4b, 0xac, 0x74,
	0x52, 0x2b, 0x34, 0x48, 0x03, 0x51, 0x20, 0x26, 0xbc, 0x96, 0xc9, 0xfc, 0xe9, 0xc1, 0xa1, 0x40,
	0x14, 0x05, 0xa4, 0xb6, 0x34, 0x69, 0x3e, 0xa7, 0x79, 0xa3, 0xb8, 0x91, 0x58, 0x39, 0xf1, 0xc1,
	0x48, 0xa0, 0x40, 0x7b, 0x4c, 0xdb, 0x93, 0xa3, 0xe3, 0x6f, 0x3e, 0xb9, 0xc3, 0x5a, 0xf7, 0xb7,
	0xd7, 0xe6, 0xf4, 0x8c, 0xdc, 0xe5, 0x79, 0x9e, 0x29, 0xf8, 0xd2, 0x80, 0x36, 0xd9, 0x0c, 0x78,
	0x0e, 0x4a, 0x87, 0x5e, 0xe4, 0xc7, 0xc1, 0x71, 0x98, 0x6c, 0x3c, 0x9a, 0x9c, 0xd9, 0xda, 0x47,
	0x5e, 0x34, 0xc0, 0x76, 0x79, 0x9e, 0x33, 0xd7, 0xe3, 0xb0, 0xa6, 0xef, 0xc8, 0xc8, 0x39, 0xe9,
	0x1a, 0x2b, 0x0d, 0xd7, 0x56, 0x9d, 0xff, 0x58, 0x51, 0x6b, 0xe5, 0x9a, 0x56, 0x5e, 0x2f, 0xc8,
	0xbe, 0x82, 0x12, 0xe7, 0xf0, 0xaf, 0x9d, 0x1f, 0xf9, 0xf1, 0x80, 0xdd, 0x73, 0xe5, 0xbf, 0xfb,
	0x8e, 0x48, 0x50, 0xf2, 0x45, 0xa6, 0xc0, 0x28, 0x09, 0x3a, 0xec, 0x46, 0x5e, 0x3c, 0x64, 0xa4,
	0xe4, 0x0b, 0xe6, 0x08, 0x7d, 0x45, 0xfa, 0x46, 0x96, 0x80, 0x8d, 0x09, 0xb7, 0x22, 0x2f, 0x0e,
	0x8e, 0xef, 0x27, 0x2e, 0xca, 0x64, 0x15, 0x65, 0xf2, 0x66, 0x19, 0xe5, 0x49, 0xf7, 0xfb, 0xcf,
	0x23, 0x8f, 0xad, 0xf4, 0xf4, 0x01, 0xb9, 0x35, 0x43, 0x6d, 0x32, 0x05, 0x97, 0x4a, 0x1a, 0x08,
	0x7b, 0x91, 0x17, 0x0f, 0x58, 0xd0, 0x32, 0xe6, 0x10, 0x7d, 0x4c, 0xba, 0x53, 0x54, 0x3a, 0xec,
	0x5b, 0xeb, 0xfd, 0x3f, 0x3e, 0xf9, 0x14, 0x95, 0xfe, 0x80, 0x85, 0x9c, 0x5e, 0x31, 0x2b, 0x1a,
	0x9f, 0x93, 0x60, 0x23, 0x06, 0xba, 0x43, 0xfc, 0x0b, 0xb8, 0x0a, 0x3d, 0xeb, 0xda, 0x1e, 0xe9,
	0x88, 0x6c, 0xcd, 0xdb, 0x52, 0xd8, 0xb1, 0xcc, 0x5d, 0xe8, 0x1e, 0xe9, 0xf1, 0xba, 0x86, 0x2a,
	0x0f, 0xfd, 0xc8, 0x8b, 0xb7, 0xd9, 0xf2, 0x36, 0xfe, 0xda, 0x21, 0x64, 0xfd, 0x46, 0x3b, 0x2d,
	0x2f, 0x0a, 0xbc, 0xcc, 0x50, 0x49, 0x21, 0x2b, 0xbb, 0xd0, 0x01, 0x0b, 0x2c, 0x7b, 0x6f, 0x11,
	0x7d, 0x48, 0x86, 0x4e, 0x52, 0x82, 0x99, 0x61, 0xae, 0x97, 0xef, 0xb8, 0xbe, 0x73, 0xc7, 0xd6,
	0xa2, 0x75, 0xfe, 0x6b, 0xd1, 0x2a, 0xf6, 0x47, 0xe4, 0x36, 0x2c, 0x6a, 0xdc, 0xd8, 0x52, 0xd7,
	0xaa, 0x86, 0x8e, 0xae, 0x64, 0x2f, 0x49, 0xbf, 0xdd, 0x0e, 0x17, 0x70, 0xd3, 0xf0, 0x7b, 0x25,
	0x5f, 0xbc, 0x16, 0x6d, 0xb0, 0xbb, 0x6e, 0x8a, 0xa9, 0x82, 0x1c, 0x2a, 0x23, 0x79, 0xa1, 0xed,
	0x02, 0xb6, 0xd9, 0x8e, 0x2d, 0x9c, 0xae, 0xf9, 0xc9, 0xf3, 0x1f, 0xbf, 0x0e, 0xbd, 0x4f, 0xa9,
	0x90, 0x66, 0xd6, 0x4c, 0x92, 0x29, 0x96, 0xa9, 0xc6, 0x02, 0x9f, 0x48, 0x4c, 0xdb, 0x7d, 0xa4,
	0xf5, 0x85, 0x48, 0xa7, 0xa8, 0xa0, 0x2e, 0x1a, 0x21, 0x2b, 0x9d, 0xb6, 0xff, 0x9b, 0xac, 0xc4,
	0xa4, 0x67, 0x87, 0x78, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x66, 0x85, 0x7f, 0xd7, 0x81, 0x03,
	0x00, 0x00,
}
