// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	BlockSigner
*/
package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bc "github.com/chainmint/protocol/bc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	IsSigner             bool           `protobuf:"varint,2,opt,name=is_signer,json=isSigner" json:"is_signer,omitempty"`
	IsGenerator          bool           `protobuf:"varint,3,opt,name=is_generator,json=isGenerator" json:"is_generator,omitempty"`
	BlockchainId         *bc.Hash       `protobuf:"bytes,4,opt,name=blockchain_id,json=blockchainId" json:"blockchain_id,omitempty"`
	GeneratorUrl         string         `protobuf:"bytes,5,opt,name=generator_url,json=generatorUrl" json:"generator_url,omitempty"`
	GeneratorAccessToken string         `protobuf:"bytes,6,opt,name=generator_access_token,json=generatorAccessToken" json:"generator_access_token,omitempty"`
	BlockHsmUrl          string         `protobuf:"bytes,7,opt,name=block_hsm_url,json=blockHsmUrl" json:"block_hsm_url,omitempty"`
	BlockHsmAccessToken  string         `protobuf:"bytes,8,opt,name=block_hsm_access_token,json=blockHsmAccessToken" json:"block_hsm_access_token,omitempty"`
	ConfiguredAt         uint64         `protobuf:"varint,9,opt,name=configured_at,json=configuredAt" json:"configured_at,omitempty"`
	BlockPub             []byte         `protobuf:"bytes,10,opt,name=block_pub,json=blockPub,proto3" json:"block_pub,omitempty"`
	Signers              []*BlockSigner `protobuf:"bytes,11,rep,name=signers" json:"signers,omitempty"`
	Quorum               uint32         `protobuf:"varint,12,opt,name=quorum" json:"quorum,omitempty"`
	MaxIssuanceWindowMs  uint64         `protobuf:"varint,13,opt,name=max_issuance_window_ms,json=maxIssuanceWindowMs" json:"max_issuance_window_ms,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Config) GetIsSigner() bool {
	if m != nil {
		return m.IsSigner
	}
	return false
}

func (m *Config) GetIsGenerator() bool {
	if m != nil {
		return m.IsGenerator
	}
	return false
}

func (m *Config) GetBlockchainId() *bc.Hash {
	if m != nil {
		return m.BlockchainId
	}
	return nil
}

func (m *Config) GetGeneratorUrl() string {
	if m != nil {
		return m.GeneratorUrl
	}
	return ""
}

func (m *Config) GetGeneratorAccessToken() string {
	if m != nil {
		return m.GeneratorAccessToken
	}
	return ""
}

func (m *Config) GetBlockHsmUrl() string {
	if m != nil {
		return m.BlockHsmUrl
	}
	return ""
}

func (m *Config) GetBlockHsmAccessToken() string {
	if m != nil {
		return m.BlockHsmAccessToken
	}
	return ""
}

func (m *Config) GetConfiguredAt() uint64 {
	if m != nil {
		return m.ConfiguredAt
	}
	return 0
}

func (m *Config) GetBlockPub() []byte {
	if m != nil {
		return m.BlockPub
	}
	return nil
}

func (m *Config) GetSigners() []*BlockSigner {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *Config) GetQuorum() uint32 {
	if m != nil {
		return m.Quorum
	}
	return 0
}

func (m *Config) GetMaxIssuanceWindowMs() uint64 {
	if m != nil {
		return m.MaxIssuanceWindowMs
	}
	return 0
}

type BlockSigner struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	Pubkey      []byte `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Url         string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
}

func (m *BlockSigner) Reset()                    { *m = BlockSigner{} }
func (m *BlockSigner) String() string            { return proto.CompactTextString(m) }
func (*BlockSigner) ProtoMessage()               {}
func (*BlockSigner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BlockSigner) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *BlockSigner) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *BlockSigner) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "config.Config")
	proto.RegisterType((*BlockSigner)(nil), "config.BlockSigner")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x4f, 0x6f, 0xd4, 0x30,
	0x10, 0xc5, 0x95, 0x4d, 0x49, 0xb3, 0x93, 0x04, 0x21, 0x2f, 0x5a, 0x59, 0xe5, 0x12, 0xb6, 0x97,
	0x5c, 0xba, 0x2b, 0xb5, 0x7c, 0x81, 0xc2, 0x81, 0xf6, 0x80, 0x84, 0x02, 0x08, 0x89, 0x8b, 0x65,
	0x3b, 0x66, 0xd7, 0xda, 0x4d, 0xbc, 0xd8, 0xb1, 0x5a, 0x3e, 0x3a, 0x37, 0xe4, 0x49, 0xf6, 0x4f,
	0x6f, 0x99, 0xf7, 0x7e, 0x79, 0x1e, 0x8f, 0x07, 0x72, 0x69, 0xba, 0xdf, 0x7a, 0xbd, 0xdc, 0x5b,
	0xd3, 0x1b, 0x92, 0x0c, 0xd5, 0xd5, 0x95, 0xdc, 0x70, 0xdd, 0xad, 0x50, 0x94, 0x66, 0xb7, 0x12,
	0x72, 0x25, 0xe4, 0xc0, 0x2c, 0xfe, 0xc5, 0x90, 0x7c, 0x42, 0x8c, 0xbc, 0x86, 0x89, 0x6e, 0x68,
	0x54, 0x46, 0xd5, 0xb4, 0x9e, 0xe8, 0x86, 0xbc, 0x83, 0xa9, 0x76, 0xcc, 0xe9, 0x75, 0xa7, 0x2c,
	0x9d, 0x94, 0x51, 0x95, 0xd6, 0xa9, 0x76, 0xdf, 0xb0, 0x26, 0xef, 0x21, 0xd7, 0x8e, 0xad, 0x55,
	0xa7, 0x2c, 0xef, 0x8d, 0xa5, 0x31, 0xfa, 0x99, 0x76, 0x9f, 0x0f, 0x12, 0xb9, 0x81, 0x42, 0xec,
	0x8c, 0xdc, 0xe2, 0xe9, 0x4c, 0x37, 0xf4, 0xa2, 0x8c, 0xaa, 0xec, 0x36, 0x5d, 0x0a, 0xb9, 0x7c,
	0xe0, 0x6e, 0x53, 0xe7, 0x27, 0xfb, 0xb1, 0x21, 0xd7, 0x50, 0x1c, 0xe3, 0x98, 0xb7, 0x3b, 0xfa,
	0x0a, 0x3b, 0xc9, 0x8f, 0xe2, 0x0f, 0xbb, 0x23, 0x1f, 0x60, 0x7e, 0x82, 0xb8, 0x94, 0xca, 0x39,
	0xd6, 0x9b, 0xad, 0xea, 0x68, 0x82, 0xf4, 0xdb, 0xa3, 0x7b, 0x8f, 0xe6, 0xf7, 0xe0, 0x91, 0xc5,
	0xd8, 0x09, 0xdb, 0xb8, 0x16, 0xa3, 0x2f, 0x11, 0xce, 0x50, 0x7c, 0x70, 0x6d, 0x48, 0xbe, 0x83,
	0xf9, 0x89, 0x79, 0x91, 0x9c, 0x22, 0x3c, 0x3b, 0xc0, 0xe7, 0xc1, 0xd7, 0x50, 0x0c, 0x33, 0xf6,
	0x56, 0x35, 0x8c, 0xf7, 0x74, 0x5a, 0x46, 0xd5, 0x45, 0x9d, 0x9f, 0xc4, 0xfb, 0x3e, 0xcc, 0x71,
	0x48, 0xde, 0x7b, 0x41, 0xa1, 0x8c, 0xaa, 0xbc, 0x4e, 0x51, 0xf8, 0xea, 0x05, 0xb9, 0x81, 0xcb,
	0x61, 0xc2, 0x8e, 0x66, 0x65, 0x5c, 0x65, 0xb7, 0xb3, 0xe5, 0xf8, 0x86, 0x1f, 0x03, 0x32, 0x4c,
	0xbb, 0x3e, 0x30, 0x64, 0x0e, 0xc9, 0x1f, 0x6f, 0xac, 0x6f, 0x69, 0x5e, 0x46, 0x55, 0x51, 0x8f,
	0x55, 0xe8, 0xbe, 0xe5, 0xcf, 0x4c, 0x3b, 0xe7, 0x79, 0x27, 0x15, 0x7b, 0xd2, 0x5d, 0x63, 0x9e,
	0x58, 0xeb, 0x68, 0x81, 0x1d, 0xcd, 0x5a, 0xfe, 0xfc, 0x38, 0x9a, 0x3f, 0xd1, 0xfb, 0xe2, 0x16,
	0xbf, 0x20, 0x3b, 0x3b, 0x24, 0x3c, 0xe9, 0x8b, 0x7b, 0x0f, 0x9b, 0x90, 0xf1, 0xb3, 0xfb, 0xce,
	0x21, 0xd9, 0x7b, 0xb1, 0x55, 0x7f, 0x71, 0x1f, 0xf2, 0x7a, 0xac, 0xc8, 0x1b, 0x88, 0xc3, 0x58,
	0x63, 0xfc, 0x23, 0x7c, 0x8a, 0x04, 0xd7, 0xeb, 0xee, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7,
	0xf2, 0xbf, 0xbe, 0x92, 0x02, 0x00, 0x00,
}
