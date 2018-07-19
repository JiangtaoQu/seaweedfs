// Code generated by protoc-gen-go.
// source: filer.proto
// DO NOT EDIT!

/*
Package filer_pb is a generated protocol buffer package.

It is generated from these files:
	filer.proto

It has these top-level messages:
	LookupDirectoryEntryRequest
	LookupDirectoryEntryResponse
	ListEntriesRequest
	ListEntriesResponse
	Entry
	FileChunk
	FuseAttributes
	GetEntryAttributesRequest
	GetEntryAttributesResponse
	GetFileContentRequest
	GetFileContentResponse
	CreateEntryRequest
	CreateEntryResponse
	UpdateEntryRequest
	UpdateEntryResponse
	DeleteEntryRequest
	DeleteEntryResponse
	AssignVolumeRequest
	AssignVolumeResponse
	LookupVolumeRequest
	Locations
	Location
	LookupVolumeResponse
*/
package filer_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LookupDirectoryEntryRequest struct {
	Directory string `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *LookupDirectoryEntryRequest) Reset()                    { *m = LookupDirectoryEntryRequest{} }
func (m *LookupDirectoryEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupDirectoryEntryRequest) ProtoMessage()               {}
func (*LookupDirectoryEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LookupDirectoryEntryRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *LookupDirectoryEntryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LookupDirectoryEntryResponse struct {
	Entry *Entry `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
}

func (m *LookupDirectoryEntryResponse) Reset()                    { *m = LookupDirectoryEntryResponse{} }
func (m *LookupDirectoryEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupDirectoryEntryResponse) ProtoMessage()               {}
func (*LookupDirectoryEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LookupDirectoryEntryResponse) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type ListEntriesRequest struct {
	Directory string `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
}

func (m *ListEntriesRequest) Reset()                    { *m = ListEntriesRequest{} }
func (m *ListEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntriesRequest) ProtoMessage()               {}
func (*ListEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListEntriesRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

type ListEntriesResponse struct {
	Entries []*Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *ListEntriesResponse) Reset()                    { *m = ListEntriesResponse{} }
func (m *ListEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEntriesResponse) ProtoMessage()               {}
func (*ListEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListEntriesResponse) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Entry struct {
	Name        string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	IsDirectory bool            `protobuf:"varint,2,opt,name=is_directory,json=isDirectory" json:"is_directory,omitempty"`
	Chunks      []*FileChunk    `protobuf:"bytes,3,rep,name=chunks" json:"chunks,omitempty"`
	Attributes  *FuseAttributes `protobuf:"bytes,4,opt,name=attributes" json:"attributes,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Entry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entry) GetIsDirectory() bool {
	if m != nil {
		return m.IsDirectory
	}
	return false
}

func (m *Entry) GetChunks() []*FileChunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func (m *Entry) GetAttributes() *FuseAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type FileChunk struct {
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	Size   uint64 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Mtime  int64  `protobuf:"varint,4,opt,name=mtime" json:"mtime,omitempty"`
}

func (m *FileChunk) Reset()                    { *m = FileChunk{} }
func (m *FileChunk) String() string            { return proto.CompactTextString(m) }
func (*FileChunk) ProtoMessage()               {}
func (*FileChunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FileChunk) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

func (m *FileChunk) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *FileChunk) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileChunk) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

type FuseAttributes struct {
	FileSize    uint64 `protobuf:"varint,1,opt,name=file_size,json=fileSize" json:"file_size,omitempty"`
	Mtime       int64  `protobuf:"varint,2,opt,name=mtime" json:"mtime,omitempty"`
	FileMode    uint32 `protobuf:"varint,3,opt,name=file_mode,json=fileMode" json:"file_mode,omitempty"`
	Uid         uint32 `protobuf:"varint,4,opt,name=uid" json:"uid,omitempty"`
	Gid         uint32 `protobuf:"varint,5,opt,name=gid" json:"gid,omitempty"`
	Crtime      int64  `protobuf:"varint,6,opt,name=crtime" json:"crtime,omitempty"`
	Mime        string `protobuf:"bytes,7,opt,name=mime" json:"mime,omitempty"`
	Replication string `protobuf:"bytes,8,opt,name=replication" json:"replication,omitempty"`
	Collection  string `protobuf:"bytes,9,opt,name=collection" json:"collection,omitempty"`
	TtlSec      int32  `protobuf:"varint,10,opt,name=ttl_sec,json=ttlSec" json:"ttl_sec,omitempty"`
}

func (m *FuseAttributes) Reset()                    { *m = FuseAttributes{} }
func (m *FuseAttributes) String() string            { return proto.CompactTextString(m) }
func (*FuseAttributes) ProtoMessage()               {}
func (*FuseAttributes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FuseAttributes) GetFileSize() uint64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *FuseAttributes) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *FuseAttributes) GetFileMode() uint32 {
	if m != nil {
		return m.FileMode
	}
	return 0
}

func (m *FuseAttributes) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *FuseAttributes) GetGid() uint32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *FuseAttributes) GetCrtime() int64 {
	if m != nil {
		return m.Crtime
	}
	return 0
}

func (m *FuseAttributes) GetMime() string {
	if m != nil {
		return m.Mime
	}
	return ""
}

func (m *FuseAttributes) GetReplication() string {
	if m != nil {
		return m.Replication
	}
	return ""
}

func (m *FuseAttributes) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *FuseAttributes) GetTtlSec() int32 {
	if m != nil {
		return m.TtlSec
	}
	return 0
}

type GetEntryAttributesRequest struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ParentDir string `protobuf:"bytes,2,opt,name=parent_dir,json=parentDir" json:"parent_dir,omitempty"`
	FileId    string `protobuf:"bytes,3,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
}

func (m *GetEntryAttributesRequest) Reset()                    { *m = GetEntryAttributesRequest{} }
func (m *GetEntryAttributesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntryAttributesRequest) ProtoMessage()               {}
func (*GetEntryAttributesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetEntryAttributesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetEntryAttributesRequest) GetParentDir() string {
	if m != nil {
		return m.ParentDir
	}
	return ""
}

func (m *GetEntryAttributesRequest) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

type GetEntryAttributesResponse struct {
	Attributes *FuseAttributes `protobuf:"bytes,1,opt,name=attributes" json:"attributes,omitempty"`
	Chunks     []*FileChunk    `protobuf:"bytes,2,rep,name=chunks" json:"chunks,omitempty"`
}

func (m *GetEntryAttributesResponse) Reset()                    { *m = GetEntryAttributesResponse{} }
func (m *GetEntryAttributesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryAttributesResponse) ProtoMessage()               {}
func (*GetEntryAttributesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetEntryAttributesResponse) GetAttributes() *FuseAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *GetEntryAttributesResponse) GetChunks() []*FileChunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

type GetFileContentRequest struct {
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
}

func (m *GetFileContentRequest) Reset()                    { *m = GetFileContentRequest{} }
func (m *GetFileContentRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFileContentRequest) ProtoMessage()               {}
func (*GetFileContentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetFileContentRequest) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

type GetFileContentResponse struct {
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *GetFileContentResponse) Reset()                    { *m = GetFileContentResponse{} }
func (m *GetFileContentResponse) String() string            { return proto.CompactTextString(m) }
func (*GetFileContentResponse) ProtoMessage()               {}
func (*GetFileContentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetFileContentResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type CreateEntryRequest struct {
	Directory string `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Entry     *Entry `protobuf:"bytes,2,opt,name=entry" json:"entry,omitempty"`
}

func (m *CreateEntryRequest) Reset()                    { *m = CreateEntryRequest{} }
func (m *CreateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateEntryRequest) ProtoMessage()               {}
func (*CreateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CreateEntryRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *CreateEntryRequest) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type CreateEntryResponse struct {
}

func (m *CreateEntryResponse) Reset()                    { *m = CreateEntryResponse{} }
func (m *CreateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateEntryResponse) ProtoMessage()               {}
func (*CreateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type UpdateEntryRequest struct {
	Directory string `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Entry     *Entry `protobuf:"bytes,2,opt,name=entry" json:"entry,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UpdateEntryRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *UpdateEntryRequest) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type UpdateEntryResponse struct {
}

func (m *UpdateEntryResponse) Reset()                    { *m = UpdateEntryResponse{} }
func (m *UpdateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryResponse) ProtoMessage()               {}
func (*UpdateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type DeleteEntryRequest struct {
	Directory    string `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	IsDirectory  bool   `protobuf:"varint,3,opt,name=is_directory,json=isDirectory" json:"is_directory,omitempty"`
	IsDeleteData bool   `protobuf:"varint,4,opt,name=is_delete_data,json=isDeleteData" json:"is_delete_data,omitempty"`
	IsRecursive  bool   `protobuf:"varint,5,opt,name=is_recursive,json=isRecursive" json:"is_recursive,omitempty"`
}

func (m *DeleteEntryRequest) Reset()                    { *m = DeleteEntryRequest{} }
func (m *DeleteEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteEntryRequest) ProtoMessage()               {}
func (*DeleteEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *DeleteEntryRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *DeleteEntryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteEntryRequest) GetIsDirectory() bool {
	if m != nil {
		return m.IsDirectory
	}
	return false
}

func (m *DeleteEntryRequest) GetIsDeleteData() bool {
	if m != nil {
		return m.IsDeleteData
	}
	return false
}

func (m *DeleteEntryRequest) GetIsRecursive() bool {
	if m != nil {
		return m.IsRecursive
	}
	return false
}

type DeleteEntryResponse struct {
}

func (m *DeleteEntryResponse) Reset()                    { *m = DeleteEntryResponse{} }
func (m *DeleteEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteEntryResponse) ProtoMessage()               {}
func (*DeleteEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type AssignVolumeRequest struct {
	Count       int32  `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Collection  string `protobuf:"bytes,2,opt,name=collection" json:"collection,omitempty"`
	Replication string `protobuf:"bytes,3,opt,name=replication" json:"replication,omitempty"`
	TtlSec      int32  `protobuf:"varint,4,opt,name=ttl_sec,json=ttlSec" json:"ttl_sec,omitempty"`
	DataCenter  string `protobuf:"bytes,5,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
}

func (m *AssignVolumeRequest) Reset()                    { *m = AssignVolumeRequest{} }
func (m *AssignVolumeRequest) String() string            { return proto.CompactTextString(m) }
func (*AssignVolumeRequest) ProtoMessage()               {}
func (*AssignVolumeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *AssignVolumeRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *AssignVolumeRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *AssignVolumeRequest) GetReplication() string {
	if m != nil {
		return m.Replication
	}
	return ""
}

func (m *AssignVolumeRequest) GetTtlSec() int32 {
	if m != nil {
		return m.TtlSec
	}
	return 0
}

func (m *AssignVolumeRequest) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

type AssignVolumeResponse struct {
	FileId    string `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
	Url       string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	PublicUrl string `protobuf:"bytes,3,opt,name=public_url,json=publicUrl" json:"public_url,omitempty"`
	Count     int32  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *AssignVolumeResponse) Reset()                    { *m = AssignVolumeResponse{} }
func (m *AssignVolumeResponse) String() string            { return proto.CompactTextString(m) }
func (*AssignVolumeResponse) ProtoMessage()               {}
func (*AssignVolumeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *AssignVolumeResponse) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

func (m *AssignVolumeResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AssignVolumeResponse) GetPublicUrl() string {
	if m != nil {
		return m.PublicUrl
	}
	return ""
}

func (m *AssignVolumeResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type LookupVolumeRequest struct {
	VolumeIds []string `protobuf:"bytes,1,rep,name=volume_ids,json=volumeIds" json:"volume_ids,omitempty"`
}

func (m *LookupVolumeRequest) Reset()                    { *m = LookupVolumeRequest{} }
func (m *LookupVolumeRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupVolumeRequest) ProtoMessage()               {}
func (*LookupVolumeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *LookupVolumeRequest) GetVolumeIds() []string {
	if m != nil {
		return m.VolumeIds
	}
	return nil
}

type Locations struct {
	Locations []*Location `protobuf:"bytes,1,rep,name=locations" json:"locations,omitempty"`
}

func (m *Locations) Reset()                    { *m = Locations{} }
func (m *Locations) String() string            { return proto.CompactTextString(m) }
func (*Locations) ProtoMessage()               {}
func (*Locations) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *Locations) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

type Location struct {
	Url       string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	PublicUrl string `protobuf:"bytes,2,opt,name=public_url,json=publicUrl" json:"public_url,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *Location) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Location) GetPublicUrl() string {
	if m != nil {
		return m.PublicUrl
	}
	return ""
}

type LookupVolumeResponse struct {
	LocationsMap map[string]*Locations `protobuf:"bytes,1,rep,name=locations_map,json=locationsMap" json:"locations_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *LookupVolumeResponse) Reset()                    { *m = LookupVolumeResponse{} }
func (m *LookupVolumeResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupVolumeResponse) ProtoMessage()               {}
func (*LookupVolumeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *LookupVolumeResponse) GetLocationsMap() map[string]*Locations {
	if m != nil {
		return m.LocationsMap
	}
	return nil
}

func init() {
	proto.RegisterType((*LookupDirectoryEntryRequest)(nil), "filer_pb.LookupDirectoryEntryRequest")
	proto.RegisterType((*LookupDirectoryEntryResponse)(nil), "filer_pb.LookupDirectoryEntryResponse")
	proto.RegisterType((*ListEntriesRequest)(nil), "filer_pb.ListEntriesRequest")
	proto.RegisterType((*ListEntriesResponse)(nil), "filer_pb.ListEntriesResponse")
	proto.RegisterType((*Entry)(nil), "filer_pb.Entry")
	proto.RegisterType((*FileChunk)(nil), "filer_pb.FileChunk")
	proto.RegisterType((*FuseAttributes)(nil), "filer_pb.FuseAttributes")
	proto.RegisterType((*GetEntryAttributesRequest)(nil), "filer_pb.GetEntryAttributesRequest")
	proto.RegisterType((*GetEntryAttributesResponse)(nil), "filer_pb.GetEntryAttributesResponse")
	proto.RegisterType((*GetFileContentRequest)(nil), "filer_pb.GetFileContentRequest")
	proto.RegisterType((*GetFileContentResponse)(nil), "filer_pb.GetFileContentResponse")
	proto.RegisterType((*CreateEntryRequest)(nil), "filer_pb.CreateEntryRequest")
	proto.RegisterType((*CreateEntryResponse)(nil), "filer_pb.CreateEntryResponse")
	proto.RegisterType((*UpdateEntryRequest)(nil), "filer_pb.UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "filer_pb.UpdateEntryResponse")
	proto.RegisterType((*DeleteEntryRequest)(nil), "filer_pb.DeleteEntryRequest")
	proto.RegisterType((*DeleteEntryResponse)(nil), "filer_pb.DeleteEntryResponse")
	proto.RegisterType((*AssignVolumeRequest)(nil), "filer_pb.AssignVolumeRequest")
	proto.RegisterType((*AssignVolumeResponse)(nil), "filer_pb.AssignVolumeResponse")
	proto.RegisterType((*LookupVolumeRequest)(nil), "filer_pb.LookupVolumeRequest")
	proto.RegisterType((*Locations)(nil), "filer_pb.Locations")
	proto.RegisterType((*Location)(nil), "filer_pb.Location")
	proto.RegisterType((*LookupVolumeResponse)(nil), "filer_pb.LookupVolumeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SeaweedFiler service

type SeaweedFilerClient interface {
	LookupDirectoryEntry(ctx context.Context, in *LookupDirectoryEntryRequest, opts ...grpc.CallOption) (*LookupDirectoryEntryResponse, error)
	ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error)
	GetEntryAttributes(ctx context.Context, in *GetEntryAttributesRequest, opts ...grpc.CallOption) (*GetEntryAttributesResponse, error)
	CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*CreateEntryResponse, error)
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error)
	AssignVolume(ctx context.Context, in *AssignVolumeRequest, opts ...grpc.CallOption) (*AssignVolumeResponse, error)
	LookupVolume(ctx context.Context, in *LookupVolumeRequest, opts ...grpc.CallOption) (*LookupVolumeResponse, error)
}

type seaweedFilerClient struct {
	cc *grpc.ClientConn
}

func NewSeaweedFilerClient(cc *grpc.ClientConn) SeaweedFilerClient {
	return &seaweedFilerClient{cc}
}

func (c *seaweedFilerClient) LookupDirectoryEntry(ctx context.Context, in *LookupDirectoryEntryRequest, opts ...grpc.CallOption) (*LookupDirectoryEntryResponse, error) {
	out := new(LookupDirectoryEntryResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/LookupDirectoryEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error) {
	out := new(ListEntriesResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/ListEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) GetEntryAttributes(ctx context.Context, in *GetEntryAttributesRequest, opts ...grpc.CallOption) (*GetEntryAttributesResponse, error) {
	out := new(GetEntryAttributesResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/GetEntryAttributes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*CreateEntryResponse, error) {
	out := new(CreateEntryResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/CreateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error) {
	out := new(UpdateEntryResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/UpdateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error) {
	out := new(DeleteEntryResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/DeleteEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) AssignVolume(ctx context.Context, in *AssignVolumeRequest, opts ...grpc.CallOption) (*AssignVolumeResponse, error) {
	out := new(AssignVolumeResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/AssignVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) LookupVolume(ctx context.Context, in *LookupVolumeRequest, opts ...grpc.CallOption) (*LookupVolumeResponse, error) {
	out := new(LookupVolumeResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/LookupVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SeaweedFiler service

type SeaweedFilerServer interface {
	LookupDirectoryEntry(context.Context, *LookupDirectoryEntryRequest) (*LookupDirectoryEntryResponse, error)
	ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error)
	GetEntryAttributes(context.Context, *GetEntryAttributesRequest) (*GetEntryAttributesResponse, error)
	CreateEntry(context.Context, *CreateEntryRequest) (*CreateEntryResponse, error)
	UpdateEntry(context.Context, *UpdateEntryRequest) (*UpdateEntryResponse, error)
	DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error)
	AssignVolume(context.Context, *AssignVolumeRequest) (*AssignVolumeResponse, error)
	LookupVolume(context.Context, *LookupVolumeRequest) (*LookupVolumeResponse, error)
}

func RegisterSeaweedFilerServer(s *grpc.Server, srv SeaweedFilerServer) {
	s.RegisterService(&_SeaweedFiler_serviceDesc, srv)
}

func _SeaweedFiler_LookupDirectoryEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupDirectoryEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).LookupDirectoryEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/LookupDirectoryEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).LookupDirectoryEntry(ctx, req.(*LookupDirectoryEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/ListEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).ListEntries(ctx, req.(*ListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_GetEntryAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).GetEntryAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/GetEntryAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).GetEntryAttributes(ctx, req.(*GetEntryAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).CreateEntry(ctx, req.(*CreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).DeleteEntry(ctx, req.(*DeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_AssignVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).AssignVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/AssignVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).AssignVolume(ctx, req.(*AssignVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_LookupVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).LookupVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/LookupVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).LookupVolume(ctx, req.(*LookupVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SeaweedFiler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "filer_pb.SeaweedFiler",
	HandlerType: (*SeaweedFilerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LookupDirectoryEntry",
			Handler:    _SeaweedFiler_LookupDirectoryEntry_Handler,
		},
		{
			MethodName: "ListEntries",
			Handler:    _SeaweedFiler_ListEntries_Handler,
		},
		{
			MethodName: "GetEntryAttributes",
			Handler:    _SeaweedFiler_GetEntryAttributes_Handler,
		},
		{
			MethodName: "CreateEntry",
			Handler:    _SeaweedFiler_CreateEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _SeaweedFiler_UpdateEntry_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _SeaweedFiler_DeleteEntry_Handler,
		},
		{
			MethodName: "AssignVolume",
			Handler:    _SeaweedFiler_AssignVolume_Handler,
		},
		{
			MethodName: "LookupVolume",
			Handler:    _SeaweedFiler_LookupVolume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "filer.proto",
}

func init() { proto.RegisterFile("filer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1004 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0x5d, 0x6f, 0xdc, 0x44,
	0x14, 0x8d, 0xd7, 0xbb, 0x9b, 0xf5, 0xdd, 0x4d, 0x81, 0xd9, 0xb4, 0x98, 0x6d, 0x52, 0x16, 0xd3,
	0xa2, 0x54, 0x48, 0x51, 0x14, 0x78, 0xa8, 0x40, 0x48, 0x54, 0x49, 0xa9, 0x2a, 0xa5, 0xaa, 0x34,
	0x21, 0x48, 0x3c, 0xad, 0x1c, 0xfb, 0x66, 0x19, 0xc5, 0x6b, 0x2f, 0x9e, 0x71, 0x50, 0x78, 0xe5,
	0x91, 0xdf, 0x81, 0x78, 0x47, 0xfc, 0x03, 0xfe, 0x18, 0x9a, 0x0f, 0x7b, 0xc7, 0x6b, 0x6f, 0x3f,
	0x1e, 0xfa, 0x36, 0x73, 0xe7, 0xce, 0xb9, 0xe7, 0xd8, 0xf7, 0x1e, 0x1b, 0x86, 0x57, 0x2c, 0xc1,
	0xfc, 0x70, 0x99, 0x67, 0x22, 0x23, 0x03, 0xb5, 0x99, 0x2d, 0x2f, 0x83, 0x57, 0x70, 0xff, 0x2c,
	0xcb, 0xae, 0x8b, 0xe5, 0x29, 0xcb, 0x31, 0x12, 0x59, 0x7e, 0xfb, 0x2c, 0x15, 0xf9, 0x2d, 0xc5,
	0x5f, 0x0b, 0xe4, 0x82, 0xec, 0x81, 0x17, 0x97, 0x07, 0xbe, 0x33, 0x75, 0x0e, 0x3c, 0xba, 0x0a,
	0x10, 0x02, 0xdd, 0x34, 0x5c, 0xa0, 0xdf, 0x51, 0x07, 0x6a, 0x1d, 0x3c, 0x83, 0xbd, 0x76, 0x40,
	0xbe, 0xcc, 0x52, 0x8e, 0xe4, 0x11, 0xf4, 0x50, 0x06, 0x14, 0xda, 0xf0, 0xf8, 0x83, 0xc3, 0x92,
	0xca, 0xa1, 0xce, 0xd3, 0xa7, 0xc1, 0x31, 0x90, 0x33, 0xc6, 0x85, 0x8c, 0x31, 0xe4, 0x6f, 0x45,
	0x27, 0xf8, 0x1e, 0xc6, 0xb5, 0x3b, 0xa6, 0xe2, 0x63, 0xd8, 0x46, 0x1d, 0xf2, 0x9d, 0xa9, 0xdb,
	0x56, 0xb3, 0x3c, 0x0f, 0xfe, 0x72, 0xa0, 0xa7, 0x42, 0x95, 0x34, 0x67, 0x25, 0x8d, 0x7c, 0x06,
	0x23, 0xc6, 0x67, 0x2b, 0x02, 0x52, 0xf6, 0x80, 0x0e, 0x19, 0xaf, 0xa4, 0x92, 0x2f, 0xa1, 0x1f,
	0xfd, 0x52, 0xa4, 0xd7, 0xdc, 0x77, 0x55, 0xa9, 0xf1, 0xaa, 0xd4, 0x0f, 0x2c, 0xc1, 0x13, 0x79,
	0x46, 0x4d, 0x0a, 0x79, 0x02, 0x10, 0x0a, 0x91, 0xb3, 0xcb, 0x42, 0x20, 0xf7, 0xbb, 0xea, 0x79,
	0xf8, 0xd6, 0x85, 0x82, 0xe3, 0xd3, 0xea, 0x9c, 0x5a, 0xb9, 0xc1, 0x15, 0x78, 0x15, 0x1c, 0xf9,
	0x18, 0xb6, 0xe5, 0x9d, 0x19, 0x8b, 0x0d, 0xdb, 0xbe, 0xdc, 0xbe, 0x88, 0xc9, 0x3d, 0xe8, 0x67,
	0x57, 0x57, 0x1c, 0x85, 0x62, 0xea, 0x52, 0xb3, 0x93, 0xda, 0x38, 0xfb, 0x1d, 0x7d, 0x77, 0xea,
	0x1c, 0x74, 0xa9, 0x5a, 0x93, 0x5d, 0xe8, 0x2d, 0x04, 0x5b, 0xa0, 0xa2, 0xe1, 0x52, 0xbd, 0x09,
	0xfe, 0xec, 0xc0, 0x9d, 0x3a, 0x0d, 0x72, 0x1f, 0x3c, 0x55, 0x4d, 0x21, 0x38, 0x0a, 0x41, 0x75,
	0xd3, 0x79, 0x0d, 0xa5, 0x63, 0xa1, 0x54, 0x57, 0x16, 0x59, 0xac, 0x8b, 0xee, 0xe8, 0x2b, 0x2f,
	0xb3, 0x18, 0xc9, 0x87, 0xe0, 0x16, 0x2c, 0x56, 0x65, 0x77, 0xa8, 0x5c, 0xca, 0xc8, 0x9c, 0xc5,
	0x7e, 0x4f, 0x47, 0xe6, 0x4c, 0x09, 0x89, 0x72, 0x85, 0xdb, 0xd7, 0x42, 0xf4, 0x4e, 0x0a, 0x59,
	0xc8, 0xe8, 0xb6, 0x7e, 0x49, 0x72, 0x4d, 0xa6, 0x30, 0xcc, 0x71, 0x99, 0xb0, 0x28, 0x14, 0x2c,
	0x4b, 0xfd, 0x81, 0x3a, 0xb2, 0x43, 0xe4, 0x01, 0x40, 0x94, 0x25, 0x09, 0x46, 0x2a, 0xc1, 0x53,
	0x09, 0x56, 0x44, 0x3e, 0x4f, 0x21, 0x92, 0x19, 0xc7, 0xc8, 0x87, 0xa9, 0x73, 0xd0, 0xa3, 0x7d,
	0x21, 0x92, 0x73, 0x8c, 0x82, 0x39, 0x7c, 0xf2, 0x1c, 0x55, 0x7b, 0xdd, 0x5a, 0xef, 0xc5, 0xb4,
	0x66, 0x5b, 0xc3, 0xec, 0x03, 0x2c, 0xc3, 0x1c, 0x53, 0x21, 0x9b, 0xc6, 0x4c, 0x89, 0xa7, 0x23,
	0xa7, 0x2c, 0xb7, 0x5f, 0x9c, 0x6b, 0xbf, 0xb8, 0xe0, 0x0f, 0x07, 0x26, 0x6d, 0x95, 0x4c, 0x43,
	0xd7, 0xfb, 0xc6, 0x79, 0xfb, 0xbe, 0xb1, 0xda, 0xb3, 0xf3, 0xc6, 0xf6, 0x0c, 0x8e, 0xe0, 0xee,
	0x73, 0x14, 0x2a, 0x9e, 0xa5, 0x02, 0x53, 0x51, 0x4a, 0xdd, 0xd4, 0x70, 0xc1, 0x31, 0xdc, 0x5b,
	0xbf, 0x61, 0x28, 0xfb, 0xb0, 0x1d, 0xe9, 0x90, 0xba, 0x32, 0xa2, 0xe5, 0x36, 0xf8, 0x19, 0xc8,
	0x49, 0x8e, 0xa1, 0xc0, 0x77, 0xf0, 0x9d, 0xca, 0x43, 0x3a, 0xaf, 0xf5, 0x90, 0xbb, 0x30, 0xae,
	0x41, 0x6b, 0x2e, 0xb2, 0xe2, 0xc5, 0x32, 0x7e, 0x5f, 0x15, 0x6b, 0xd0, 0xa6, 0xe2, 0x3f, 0x0e,
	0x90, 0x53, 0x4c, 0xf0, 0x9d, 0x4a, 0xb6, 0x98, 0x6b, 0xc3, 0x81, 0xdc, 0xa6, 0x03, 0x3d, 0x84,
	0x3b, 0x32, 0x45, 0x55, 0x9b, 0xc5, 0xa1, 0x08, 0xd5, 0x68, 0x0d, 0xe8, 0x88, 0x71, 0x4d, 0xe1,
	0x34, 0x14, 0xa1, 0x01, 0xca, 0x31, 0x2a, 0x72, 0xce, 0x6e, 0x50, 0x0d, 0x9b, 0x02, 0xa2, 0x65,
	0x48, 0x6a, 0xa9, 0x71, 0x36, 0x5a, 0xfe, 0x76, 0x60, 0xfc, 0x94, 0x73, 0x36, 0x4f, 0x7f, 0xca,
	0x92, 0x62, 0x81, 0xa5, 0x98, 0x5d, 0xe8, 0x45, 0x59, 0x61, 0xde, 0x6f, 0x8f, 0xea, 0xcd, 0xda,
	0xac, 0x75, 0x1a, 0xb3, 0xb6, 0x36, 0xad, 0x6e, 0x73, 0x5a, 0xad, 0x69, 0xec, 0xda, 0xd3, 0x48,
	0x3e, 0x85, 0xa1, 0x94, 0x37, 0x8b, 0x30, 0x15, 0x98, 0x2b, 0x05, 0x1e, 0x05, 0x19, 0x3a, 0x51,
	0x91, 0xe0, 0x06, 0x76, 0xeb, 0x44, 0x4d, 0x2f, 0x6e, 0xf4, 0x4b, 0x69, 0x45, 0x79, 0x62, 0x58,
	0xca, 0xa5, 0x1a, 0xe0, 0xe2, 0x32, 0x61, 0xd1, 0x4c, 0x1e, 0xb8, 0x66, 0x80, 0x55, 0xe4, 0x22,
	0x4f, 0x56, 0x9a, 0xbb, 0x96, 0xe6, 0xe0, 0x6b, 0x18, 0xeb, 0x2f, 0x60, 0xfd, 0x01, 0xed, 0x03,
	0xdc, 0xa8, 0xc0, 0x8c, 0xc5, 0xfa, 0x4b, 0xe4, 0x51, 0x4f, 0x47, 0x5e, 0xc4, 0x3c, 0xf8, 0x0e,
	0xbc, 0xb3, 0x4c, 0x6b, 0xe6, 0xe4, 0x08, 0xbc, 0xa4, 0xdc, 0x98, 0x8f, 0x16, 0x59, 0xb5, 0x5c,
	0x99, 0x47, 0x57, 0x49, 0xc1, 0xb7, 0x30, 0x28, 0xc3, 0xa5, 0x0e, 0x67, 0x93, 0x8e, 0xce, 0x9a,
	0x8e, 0xe0, 0x3f, 0x07, 0x76, 0xeb, 0x94, 0xcd, 0xa3, 0xba, 0x80, 0x9d, 0xaa, 0xc4, 0x6c, 0x11,
	0x2e, 0x0d, 0x97, 0x23, 0x9b, 0x4b, 0xf3, 0x5a, 0x45, 0x90, 0xbf, 0x0c, 0x97, 0xba, 0x7b, 0x46,
	0x89, 0x15, 0x9a, 0xfc, 0x08, 0x1f, 0x35, 0x52, 0x24, 0xeb, 0x6b, 0x2c, 0xe7, 0x40, 0x2e, 0xc9,
	0x63, 0xe8, 0xdd, 0x84, 0x49, 0x81, 0x66, 0xe8, 0xc6, 0xcd, 0x27, 0xc0, 0xa9, 0xce, 0xf8, 0xa6,
	0xf3, 0xc4, 0x39, 0xfe, 0xb7, 0x07, 0xa3, 0x73, 0x0c, 0x7f, 0x43, 0x8c, 0xa5, 0x05, 0xe5, 0x64,
	0x5e, 0xaa, 0xaa, 0xff, 0x8a, 0x90, 0x47, 0xeb, 0xf4, 0x5b, 0xff, 0x7d, 0x26, 0x5f, 0xbc, 0x29,
	0xcd, 0x4c, 0xc4, 0x16, 0x39, 0x83, 0xa1, 0xf5, 0xe3, 0x41, 0xf6, 0xac, 0x8b, 0x8d, 0x7f, 0x98,
	0xc9, 0xfe, 0x86, 0xd3, 0x0a, 0x2d, 0x04, 0xd2, 0x34, 0x7f, 0xf2, 0xf9, 0xea, 0xda, 0xc6, 0x8f,
	0xd0, 0xe4, 0xe1, 0xeb, 0x93, 0x6c, 0xc2, 0x96, 0x33, 0xda, 0x84, 0x9b, 0x5e, 0x6c, 0x13, 0x6e,
	0xb3, 0x53, 0x85, 0x66, 0xb9, 0x9e, 0x8d, 0xd6, 0xf4, 0x59, 0x1b, 0xad, 0xcd, 0x2a, 0x15, 0x9a,
	0xe5, 0x3b, 0x36, 0x5a, 0xd3, 0x42, 0x6d, 0xb4, 0x36, 0xb3, 0xda, 0x22, 0xaf, 0x60, 0x64, 0x9b,
	0x00, 0xb1, 0x2e, 0xb4, 0xb8, 0xd8, 0xe4, 0xc1, 0xa6, 0x63, 0x1b, 0xd0, 0xee, 0x79, 0x1b, 0xb0,
	0x65, 0xea, 0x6d, 0xc0, 0xb6, 0x51, 0x09, 0xb6, 0x2e, 0xfb, 0xea, 0x97, 0xfc, 0xab, 0xff, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x27, 0x89, 0x58, 0x67, 0xa1, 0x0b, 0x00, 0x00,
}
