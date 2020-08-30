// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: google/api/http.proto

package annotations

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Defines the HTTP configuration for an API service. It contains a list of
// [HttpRule][google.api.HttpRule], each specifying the mapping of an RPC method
// to one or more HTTP REST API methods.
type Http struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of HTTP configuration rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*HttpRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	// When set to true, URL path parmeters will be fully URI-decoded except in
	// cases of single segment matches in reserved expansion, where "%2F" will be
	// left encoded.
	//
	// The default behavior is to not decode RFC 6570 reserved characters in multi
	// segment matches.
	FullyDecodeReservedExpansion bool `protobuf:"varint,2,opt,name=fully_decode_reserved_expansion,json=fullyDecodeReservedExpansion,proto3" json:"fully_decode_reserved_expansion,omitempty"`
}

func (x *Http) Reset() {
	*x = Http{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http) ProtoMessage() {}

func (x *Http) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http.ProtoReflect.Descriptor instead.
func (*Http) Descriptor() ([]byte, []int) {
	return file_google_api_http_proto_rawDescGZIP(), []int{0}
}

func (x *Http) GetRules() []*HttpRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *Http) GetFullyDecodeReservedExpansion() bool {
	if x != nil {
		return x.FullyDecodeReservedExpansion
	}
	return false
}

// `HttpRule` defines the mapping of an RPC method to one or more HTTP
// REST API methods. The mapping specifies how different portions of the RPC
// request message are mapped to URL path, URL query parameters, and
// HTTP request body. The mapping is typically specified as an
// `google.api.http` annotation on the RPC method,
// see "google/api/annotations.proto" for details.
//
// The mapping consists of a field specifying the path template and
// method kind.  The path template can refer to fields in the request
// message, as in the example below which describes a REST GET
// operation on a resource collection of messages:
//
//
//     service Messaging {
//       rpc GetMessage(GetMessageRequest) returns (Message) {
//         option (google.api.http).get = "/v1/messages/{message_id}/{sub.subfield}";
//       }
//     }
//     message GetMessageRequest {
//       message SubMessage {
//         string subfield = 1;
//       }
//       string message_id = 1; // mapped to the URL
//       SubMessage sub = 2;    // `sub.subfield` is url-mapped
//     }
//     message Message {
//       string text = 1; // content of the resource
//     }
//
// The same http annotation can alternatively be expressed inside the
// `GRPC API Configuration` YAML file.
//
//     http:
//       rules:
//         - selector: <proto_package_name>.Messaging.GetMessage
//           get: /v1/messages/{message_id}/{sub.subfield}
//
// This definition enables an automatic, bidrectional mapping of HTTP
// JSON to RPC. Example:
//
// HTTP | RPC
// -----|-----
// `GET /v1/messages/123456/foo`  | `GetMessage(message_id: "123456" sub: SubMessage(subfield: "foo"))`
//
// In general, not only fields but also field paths can be referenced
// from a path pattern. Fields mapped to the path pattern cannot be
// repeated and must have a primitive (non-message) type.
//
// Any fields in the request message which are not bound by the path
// pattern automatically become (optional) HTTP query
// parameters. Assume the following definition of the request message:
//
//
//     service Messaging {
//       rpc GetMessage(GetMessageRequest) returns (Message) {
//         option (google.api.http).get = "/v1/messages/{message_id}";
//       }
//     }
//     message GetMessageRequest {
//       message SubMessage {
//         string subfield = 1;
//       }
//       string message_id = 1; // mapped to the URL
//       int64 revision = 2;    // becomes a parameter
//       SubMessage sub = 3;    // `sub.subfield` becomes a parameter
//     }
//
//
// This enables a HTTP JSON to RPC mapping as below:
//
// HTTP | RPC
// -----|-----
// `GET /v1/messages/123456?revision=2&sub.subfield=foo` | `GetMessage(message_id: "123456" revision: 2 sub: SubMessage(subfield: "foo"))`
//
// Note that fields which are mapped to HTTP parameters must have a
// primitive type or a repeated primitive type. Message types are not
// allowed. In the case of a repeated type, the parameter can be
// repeated in the URL, as in `...?param=A&param=B`.
//
// For HTTP method kinds which allow a request body, the `body` field
// specifies the mapping. Consider a REST update method on the
// message resource collection:
//
//
//     service Messaging {
//       rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
//         option (google.api.http) = {
//           put: "/v1/messages/{message_id}"
//           body: "message"
//         };
//       }
//     }
//     message UpdateMessageRequest {
//       string message_id = 1; // mapped to the URL
//       Message message = 2;   // mapped to the body
//     }
//
//
// The following HTTP JSON to RPC mapping is enabled, where the
// representation of the JSON in the request body is determined by
// protos JSON encoding:
//
// HTTP | RPC
// -----|-----
// `PUT /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id: "123456" message { text: "Hi!" })`
//
// The special name `*` can be used in the body mapping to define that
// every field not bound by the path template should be mapped to the
// request body.  This enables the following alternative definition of
// the update method:
//
//     service Messaging {
//       rpc UpdateMessage(Message) returns (Message) {
//         option (google.api.http) = {
//           put: "/v1/messages/{message_id}"
//           body: "*"
//         };
//       }
//     }
//     message Message {
//       string message_id = 1;
//       string text = 2;
//     }
//
//
// The following HTTP JSON to RPC mapping is enabled:
//
// HTTP | RPC
// -----|-----
// `PUT /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id: "123456" text: "Hi!")`
//
// Note that when using `*` in the body mapping, it is not possible to
// have HTTP parameters, as all fields not bound by the path end in
// the body. This makes this option more rarely used in practice of
// defining REST APIs. The common usage of `*` is in custom methods
// which don't use the URL at all for transferring data.
//
// It is possible to define multiple HTTP methods for one RPC by using
// the `additional_bindings` option. Example:
//
//     service Messaging {
//       rpc GetMessage(GetMessageRequest) returns (Message) {
//         option (google.api.http) = {
//           get: "/v1/messages/{message_id}"
//           additional_bindings {
//             get: "/v1/users/{user_id}/messages/{message_id}"
//           }
//         };
//       }
//     }
//     message GetMessageRequest {
//       string message_id = 1;
//       string user_id = 2;
//     }
//
//
// This enables the following two alternative HTTP JSON to RPC
// mappings:
//
// HTTP | RPC
// -----|-----
// `GET /v1/messages/123456` | `GetMessage(message_id: "123456")`
// `GET /v1/users/me/messages/123456` | `GetMessage(user_id: "me" message_id: "123456")`
//
// # Rules for HTTP mapping
//
// The rules for mapping HTTP path, query parameters, and body fields
// to the request message are as follows:
//
// 1. The `body` field specifies either `*` or a field path, or is
//    omitted. If omitted, it indicates there is no HTTP request body.
// 2. Leaf fields (recursive expansion of nested messages in the
//    request) can be classified into three types:
//     (a) Matched in the URL template.
//     (b) Covered by body (if body is `*`, everything except (a) fields;
//         else everything under the body field)
//     (c) All other fields.
// 3. URL query parameters found in the HTTP request are mapped to (c) fields.
// 4. Any body sent with an HTTP request can contain only (b) fields.
//
// The syntax of the path template is as follows:
//
//     Template = "/" Segments [ Verb ] ;
//     Segments = Segment { "/" Segment } ;
//     Segment  = "*" | "**" | LITERAL | Variable ;
//     Variable = "{" FieldPath [ "=" Segments ] "}" ;
//     FieldPath = IDENT { "." IDENT } ;
//     Verb     = ":" LITERAL ;
//
// The syntax `*` matches a single path segment. The syntax `**` matches zero
// or more path segments, which must be the last part of the path except the
// `Verb`. The syntax `LITERAL` matches literal text in the path.
//
// The syntax `Variable` matches part of the URL path as specified by its
// template. A variable template must not contain other variables. If a variable
// matches a single path segment, its template may be omitted, e.g. `{var}`
// is equivalent to `{var=*}`.
//
// If a variable contains exactly one path segment, such as `"{var}"` or
// `"{var=*}"`, when such a variable is expanded into a URL path, all characters
// except `[-_.~0-9a-zA-Z]` are percent-encoded. Such variables show up in the
// Discovery Document as `{var}`.
//
// If a variable contains one or more path segments, such as `"{var=foo/*}"`
// or `"{var=**}"`, when such a variable is expanded into a URL path, all
// characters except `[-_.~/0-9a-zA-Z]` are percent-encoded. Such variables
// show up in the Discovery Document as `{+var}`.
//
// NOTE: While the single segment variable matches the semantics of
// [RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.2
// Simple String Expansion, the multi segment variable **does not** match
// RFC 6570 Reserved Expansion. The reason is that the Reserved Expansion
// does not expand special characters like `?` and `#`, which would lead
// to invalid URLs.
//
// NOTE: the field paths in variables and in the `body` must not refer to
// repeated fields or map fields.
type HttpRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selects methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Determines the URL pattern is matched by this rules. This pattern can be
	// used with any of the {get|put|post|delete|patch} methods. A custom method
	// can be defined using the 'custom' field.
	//
	// Types that are assignable to Pattern:
	//	*HttpRule_Get
	//	*HttpRule_Put
	//	*HttpRule_Post
	//	*HttpRule_Delete
	//	*HttpRule_Patch
	//	*HttpRule_Custom
	Pattern isHttpRule_Pattern `protobuf_oneof:"pattern"`
	// The name of the request field whose value is mapped to the HTTP body, or
	// `*` for mapping all fields not captured by the path pattern to the HTTP
	// body. NOTE: the referred field must not be a repeated field and must be
	// present at the top-level of request message type.
	Body string `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	// Optional. The name of the response field whose value is mapped to the HTTP
	// body of response. Other response fields are ignored. When
	// not set, the response message will be used as HTTP body of response.
	ResponseBody string `protobuf:"bytes,12,opt,name=response_body,json=responseBody,proto3" json:"response_body,omitempty"`
	// Additional HTTP bindings for the selector. Nested bindings must
	// not contain an `additional_bindings` field themselves (that is,
	// the nesting may only be one level deep).
	AdditionalBindings []*HttpRule `protobuf:"bytes,11,rep,name=additional_bindings,json=additionalBindings,proto3" json:"additional_bindings,omitempty"`
}

func (x *HttpRule) Reset() {
	*x = HttpRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRule) ProtoMessage() {}

func (x *HttpRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRule.ProtoReflect.Descriptor instead.
func (*HttpRule) Descriptor() ([]byte, []int) {
	return file_google_api_http_proto_rawDescGZIP(), []int{1}
}

func (x *HttpRule) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (m *HttpRule) GetPattern() isHttpRule_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (x *HttpRule) GetGet() string {
	if x, ok := x.GetPattern().(*HttpRule_Get); ok {
		return x.Get
	}
	return ""
}

func (x *HttpRule) GetPut() string {
	if x, ok := x.GetPattern().(*HttpRule_Put); ok {
		return x.Put
	}
	return ""
}

func (x *HttpRule) GetPost() string {
	if x, ok := x.GetPattern().(*HttpRule_Post); ok {
		return x.Post
	}
	return ""
}

func (x *HttpRule) GetDelete() string {
	if x, ok := x.GetPattern().(*HttpRule_Delete); ok {
		return x.Delete
	}
	return ""
}

func (x *HttpRule) GetPatch() string {
	if x, ok := x.GetPattern().(*HttpRule_Patch); ok {
		return x.Patch
	}
	return ""
}

func (x *HttpRule) GetCustom() *CustomHttpPattern {
	if x, ok := x.GetPattern().(*HttpRule_Custom); ok {
		return x.Custom
	}
	return nil
}

func (x *HttpRule) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *HttpRule) GetResponseBody() string {
	if x != nil {
		return x.ResponseBody
	}
	return ""
}

func (x *HttpRule) GetAdditionalBindings() []*HttpRule {
	if x != nil {
		return x.AdditionalBindings
	}
	return nil
}

type isHttpRule_Pattern interface {
	isHttpRule_Pattern()
}

type HttpRule_Get struct {
	// Used for listing and getting information about resources.
	Get string `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}

type HttpRule_Put struct {
	// Used for updating a resource.
	Put string `protobuf:"bytes,3,opt,name=put,proto3,oneof"`
}

type HttpRule_Post struct {
	// Used for creating a resource.
	Post string `protobuf:"bytes,4,opt,name=post,proto3,oneof"`
}

type HttpRule_Delete struct {
	// Used for deleting a resource.
	Delete string `protobuf:"bytes,5,opt,name=delete,proto3,oneof"`
}

type HttpRule_Patch struct {
	// Used for updating a resource.
	Patch string `protobuf:"bytes,6,opt,name=patch,proto3,oneof"`
}

type HttpRule_Custom struct {
	// The custom pattern is used for specifying an HTTP method that is not
	// included in the `pattern` field, such as HEAD, or "*" to leave the
	// HTTP method unspecified for this rule. The wild-card rule is useful
	// for services that provide content to Web (HTML) clients.
	Custom *CustomHttpPattern `protobuf:"bytes,8,opt,name=custom,proto3,oneof"`
}

func (*HttpRule_Get) isHttpRule_Pattern() {}

func (*HttpRule_Put) isHttpRule_Pattern() {}

func (*HttpRule_Post) isHttpRule_Pattern() {}

func (*HttpRule_Delete) isHttpRule_Pattern() {}

func (*HttpRule_Patch) isHttpRule_Pattern() {}

func (*HttpRule_Custom) isHttpRule_Pattern() {}

// A custom pattern is used for defining custom HTTP verb.
type CustomHttpPattern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this custom HTTP verb.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// The path matched by this custom verb.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CustomHttpPattern) Reset() {
	*x = CustomHttpPattern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomHttpPattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomHttpPattern) ProtoMessage() {}

func (x *CustomHttpPattern) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomHttpPattern.ProtoReflect.Descriptor instead.
func (*CustomHttpPattern) Descriptor() ([]byte, []int) {
	return file_google_api_http_proto_rawDescGZIP(), []int{2}
}

func (x *CustomHttpPattern) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CustomHttpPattern) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_google_api_http_proto protoreflect.FileDescriptor

var file_google_api_http_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x22, 0x79, 0x0a, 0x04, 0x48, 0x74, 0x74, 0x70, 0x12, 0x2a, 0x0a, 0x05, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x1f, 0x66, 0x75, 0x6c, 0x6c, 0x79,
	0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x5f, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x1c, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xda,
	0x02, 0x0a, 0x08, 0x48, 0x74, 0x74, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x03, 0x70,
	0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x70, 0x75, 0x74, 0x12,
	0x14, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x37, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x74, 0x74, 0x70, 0x50,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x45, 0x0a, 0x13, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x12, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x22, 0x3b, 0x0a, 0x11, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x74, 0x74, 0x70, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42, 0x6a, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x09, 0x48, 0x74, 0x74, 0x70,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x04,
	0x47, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_http_proto_rawDescOnce sync.Once
	file_google_api_http_proto_rawDescData = file_google_api_http_proto_rawDesc
)

func file_google_api_http_proto_rawDescGZIP() []byte {
	file_google_api_http_proto_rawDescOnce.Do(func() {
		file_google_api_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_http_proto_rawDescData)
	})
	return file_google_api_http_proto_rawDescData
}

var file_google_api_http_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_api_http_proto_goTypes = []interface{}{
	(*Http)(nil),              // 0: google.api.Http
	(*HttpRule)(nil),          // 1: google.api.HttpRule
	(*CustomHttpPattern)(nil), // 2: google.api.CustomHttpPattern
}
var file_google_api_http_proto_depIdxs = []int32{
	1, // 0: google.api.Http.rules:type_name -> google.api.HttpRule
	2, // 1: google.api.HttpRule.custom:type_name -> google.api.CustomHttpPattern
	1, // 2: google.api.HttpRule.additional_bindings:type_name -> google.api.HttpRule
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_api_http_proto_init() }
func file_google_api_http_proto_init() {
	if File_google_api_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_api_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpRule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_api_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomHttpPattern); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_api_http_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*HttpRule_Get)(nil),
		(*HttpRule_Put)(nil),
		(*HttpRule_Post)(nil),
		(*HttpRule_Delete)(nil),
		(*HttpRule_Patch)(nil),
		(*HttpRule_Custom)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_api_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_http_proto_goTypes,
		DependencyIndexes: file_google_api_http_proto_depIdxs,
		MessageInfos:      file_google_api_http_proto_msgTypes,
	}.Build()
	File_google_api_http_proto = out.File
	file_google_api_http_proto_rawDesc = nil
	file_google_api_http_proto_goTypes = nil
	file_google_api_http_proto_depIdxs = nil
}
