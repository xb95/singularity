// Code generated by protoc-gen-go from "singularity.proto"
// DO NOT EDIT!

package singularity

import proto "code.google.com/p/goprotobuf/proto"
import "math"

// Reference proto and math imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf

type Command struct {
	Command          []byte   `protobuf:"bytes,1,req,name=command" json:"command,omitempty"`
	Args             [][]byte `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *Command) Reset()         { *this = Command{} }
func (this *Command) String() string { return proto.CompactTextString(this) }

type StillAlive struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *StillAlive) Reset()         { *this = StillAlive{} }
func (this *StillAlive) String() string { return proto.CompactTextString(this) }

type Response struct {
	ExitCode         *int32 `protobuf:"varint,1,req,name=exit_code" json:"exit_code,omitempty"`
	Stdout           []byte `protobuf:"bytes,2,opt,name=stdout" json:"stdout,omitempty"`
	Stderr           []byte `protobuf:"bytes,3,opt,name=stderr" json:"stderr,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *Response) Reset()         { *this = Response{} }
func (this *Response) String() string { return proto.CompactTextString(this) }

func init() {
}
