// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.34
// source: provisionersdk/proto/provisioner.proto

package proto

import (
	context "context"
	errors "errors"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_provisionersdk_proto_provisioner_proto struct{}

func (drpcEncoding_File_provisionersdk_proto_provisioner_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_provisionersdk_proto_provisioner_proto) MarshalAppend(buf []byte, msg drpc.Message) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(buf, msg.(proto.Message))
}

func (drpcEncoding_File_provisionersdk_proto_provisioner_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_provisionersdk_proto_provisioner_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	return protojson.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_provisionersdk_proto_provisioner_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return protojson.Unmarshal(buf, msg.(proto.Message))
}

type DRPCProvisionerClient interface {
	DRPCConn() drpc.Conn

	Session(ctx context.Context) (DRPCProvisioner_SessionClient, error)
}

type drpcProvisionerClient struct {
	cc drpc.Conn
}

func NewDRPCProvisionerClient(cc drpc.Conn) DRPCProvisionerClient {
	return &drpcProvisionerClient{cc}
}

func (c *drpcProvisionerClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcProvisionerClient) Session(ctx context.Context) (DRPCProvisioner_SessionClient, error) {
	stream, err := c.cc.NewStream(ctx, "/provisioner.Provisioner/Session", drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcProvisioner_SessionClient{stream}
	return x, nil
}

type DRPCProvisioner_SessionClient interface {
	drpc.Stream
	Send(*Request) error
	Recv() (*Response, error)
}

type drpcProvisioner_SessionClient struct {
	drpc.Stream
}

func (x *drpcProvisioner_SessionClient) GetStream() drpc.Stream {
	return x.Stream
}

func (x *drpcProvisioner_SessionClient) Send(m *Request) error {
	return x.MsgSend(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
}

func (x *drpcProvisioner_SessionClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.MsgRecv(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcProvisioner_SessionClient) RecvMsg(m *Response) error {
	return x.MsgRecv(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
}

type DRPCProvisionerServer interface {
	Session(DRPCProvisioner_SessionStream) error
}

type DRPCProvisionerUnimplementedServer struct{}

func (s *DRPCProvisionerUnimplementedServer) Session(DRPCProvisioner_SessionStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCProvisionerDescription struct{}

func (DRPCProvisionerDescription) NumMethods() int { return 1 }

func (DRPCProvisionerDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/provisioner.Provisioner/Session", drpcEncoding_File_provisionersdk_proto_provisioner_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCProvisionerServer).
					Session(
						&drpcProvisioner_SessionStream{in1.(drpc.Stream)},
					)
			}, DRPCProvisionerServer.Session, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterProvisioner(mux drpc.Mux, impl DRPCProvisionerServer) error {
	return mux.Register(impl, DRPCProvisionerDescription{})
}

type DRPCProvisioner_SessionStream interface {
	drpc.Stream
	Send(*Response) error
	Recv() (*Request, error)
}

type drpcProvisioner_SessionStream struct {
	drpc.Stream
}

func (x *drpcProvisioner_SessionStream) Send(m *Response) error {
	return x.MsgSend(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
}

func (x *drpcProvisioner_SessionStream) Recv() (*Request, error) {
	m := new(Request)
	if err := x.MsgRecv(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcProvisioner_SessionStream) RecvMsg(m *Request) error {
	return x.MsgRecv(m, drpcEncoding_File_provisionersdk_proto_provisioner_proto{})
}
