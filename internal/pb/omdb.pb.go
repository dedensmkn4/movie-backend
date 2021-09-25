// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: omdb.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchWord string `protobuf:"bytes,1,opt,name=searchWord,proto3" json:"searchWord,omitempty"`
	Page       int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetSearchWord() string {
	if x != nil {
		return x.SearchWord
	}
	return ""
}

func (x *SearchRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type SearchPaginatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search       []*SearchResponse `protobuf:"bytes,1,rep,name=search,proto3" json:"search,omitempty"`
	TotalResults string            `protobuf:"bytes,2,opt,name=totalResults,proto3" json:"totalResults,omitempty"`
	Response     string            `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SearchPaginatedResponse) Reset() {
	*x = SearchPaginatedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPaginatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPaginatedResponse) ProtoMessage() {}

func (x *SearchPaginatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPaginatedResponse.ProtoReflect.Descriptor instead.
func (*SearchPaginatedResponse) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{1}
}

func (x *SearchPaginatedResponse) GetSearch() []*SearchResponse {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *SearchPaginatedResponse) GetTotalResults() string {
	if x != nil {
		return x.TotalResults
	}
	return ""
}

func (x *SearchPaginatedResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{2}
}

func (x *SearchResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SearchResponse) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *SearchResponse) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *SearchResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SearchResponse) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type GetDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDetailRequest) Reset() {
	*x = GetDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailRequest) ProtoMessage() {}

func (x *GetDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailRequest.ProtoReflect.Descriptor instead.
func (*GetDetailRequest) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{3}
}

func (x *GetDetailRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{4}
}

func (x *Rating) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Rating) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SearchDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rated      string    `protobuf:"bytes,1,opt,name=rated,proto3" json:"rated,omitempty"`
	Released   string    `protobuf:"bytes,2,opt,name=released,proto3" json:"released,omitempty"`
	Runtime    string    `protobuf:"bytes,3,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Genre      string    `protobuf:"bytes,4,opt,name=genre,proto3" json:"genre,omitempty"`
	Director   string    `protobuf:"bytes,5,opt,name=director,proto3" json:"director,omitempty"`
	Writer     string    `protobuf:"bytes,6,opt,name=writer,proto3" json:"writer,omitempty"`
	Actors     string    `protobuf:"bytes,7,opt,name=actors,proto3" json:"actors,omitempty"`
	Plot       string    `protobuf:"bytes,8,opt,name=plot,proto3" json:"plot,omitempty"`
	Language   string    `protobuf:"bytes,9,opt,name=language,proto3" json:"language,omitempty"`
	Country    string    `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty"`
	Awards     string    `protobuf:"bytes,11,opt,name=awards,proto3" json:"awards,omitempty"`
	Ratings    []*Rating `protobuf:"bytes,12,rep,name=ratings,proto3" json:"ratings,omitempty"`
	MetaScore  string    `protobuf:"bytes,13,opt,name=metaScore,proto3" json:"metaScore,omitempty"`
	ImdbRating string    `protobuf:"bytes,14,opt,name=imdbRating,proto3" json:"imdbRating,omitempty"`
	ImdbVotes  string    `protobuf:"bytes,15,opt,name=imdbVotes,proto3" json:"imdbVotes,omitempty"`
	Dvd        string    `protobuf:"bytes,16,opt,name=dvd,proto3" json:"dvd,omitempty"`
	BoxOffice  string    `protobuf:"bytes,17,opt,name=boxOffice,proto3" json:"boxOffice,omitempty"`
	Production string    `protobuf:"bytes,18,opt,name=production,proto3" json:"production,omitempty"`
	Website    string    `protobuf:"bytes,19,opt,name=website,proto3" json:"website,omitempty"`
	Response   string    `protobuf:"bytes,20,opt,name=response,proto3" json:"response,omitempty"`
	Title      string    `protobuf:"bytes,21,opt,name=title,proto3" json:"title,omitempty"`
	Year       string    `protobuf:"bytes,22,opt,name=year,proto3" json:"year,omitempty"`
	ImdbID     string    `protobuf:"bytes,23,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type       string    `protobuf:"bytes,24,opt,name=type,proto3" json:"type,omitempty"`
	Poster     string    `protobuf:"bytes,25,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *SearchDetailResponse) Reset() {
	*x = SearchDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDetailResponse) ProtoMessage() {}

func (x *SearchDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDetailResponse.ProtoReflect.Descriptor instead.
func (*SearchDetailResponse) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{5}
}

func (x *SearchDetailResponse) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *SearchDetailResponse) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *SearchDetailResponse) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *SearchDetailResponse) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *SearchDetailResponse) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *SearchDetailResponse) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *SearchDetailResponse) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *SearchDetailResponse) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *SearchDetailResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *SearchDetailResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *SearchDetailResponse) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *SearchDetailResponse) GetRatings() []*Rating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *SearchDetailResponse) GetMetaScore() string {
	if x != nil {
		return x.MetaScore
	}
	return ""
}

func (x *SearchDetailResponse) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *SearchDetailResponse) GetImdbVotes() string {
	if x != nil {
		return x.ImdbVotes
	}
	return ""
}

func (x *SearchDetailResponse) GetDvd() string {
	if x != nil {
		return x.Dvd
	}
	return ""
}

func (x *SearchDetailResponse) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *SearchDetailResponse) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *SearchDetailResponse) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *SearchDetailResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *SearchDetailResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SearchDetailResponse) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *SearchDetailResponse) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *SearchDetailResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SearchDetailResponse) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

var File_omdb_proto protoreflect.FileDescriptor

var file_omdb_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x43, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x6f,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x7e, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x9f, 0x05,
	0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x69, 0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x76, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x76, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x62, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x62, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x77,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64,
	0x62, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x32,
	0xb2, 0x01, 0x0a, 0x0b, 0x4f, 0x6d, 0x64, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x52, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x6f, 0x6d, 0x64, 0x62, 0x2f,
	0x61, 0x6c, 0x6c, 0x12, 0x4f, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f,
	0x6f, 0x6d, 0x64, 0x62, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x64, 0x65, 0x6e, 0x73, 0x6d, 0x6b, 0x6e, 0x34, 0x2f, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omdb_proto_rawDescOnce sync.Once
	file_omdb_proto_rawDescData = file_omdb_proto_rawDesc
)

func file_omdb_proto_rawDescGZIP() []byte {
	file_omdb_proto_rawDescOnce.Do(func() {
		file_omdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_omdb_proto_rawDescData)
	})
	return file_omdb_proto_rawDescData
}

var file_omdb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_omdb_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),           // 0: proto.SearchRequest
	(*SearchPaginatedResponse)(nil), // 1: proto.SearchPaginatedResponse
	(*SearchResponse)(nil),          // 2: proto.SearchResponse
	(*GetDetailRequest)(nil),        // 3: proto.GetDetailRequest
	(*Rating)(nil),                  // 4: proto.Rating
	(*SearchDetailResponse)(nil),    // 5: proto.SearchDetailResponse
}
var file_omdb_proto_depIdxs = []int32{
	2, // 0: proto.SearchPaginatedResponse.search:type_name -> proto.SearchResponse
	4, // 1: proto.SearchDetailResponse.ratings:type_name -> proto.Rating
	0, // 2: proto.OmdbService.FindAll:input_type -> proto.SearchRequest
	3, // 3: proto.OmdbService.FindById:input_type -> proto.GetDetailRequest
	1, // 4: proto.OmdbService.FindAll:output_type -> proto.SearchPaginatedResponse
	5, // 5: proto.OmdbService.FindById:output_type -> proto.SearchDetailResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_omdb_proto_init() }
func file_omdb_proto_init() {
	if File_omdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_omdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPaginatedResponse); i {
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
		file_omdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_omdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDetailRequest); i {
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
		file_omdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
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
		file_omdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchDetailResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_omdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_omdb_proto_goTypes,
		DependencyIndexes: file_omdb_proto_depIdxs,
		MessageInfos:      file_omdb_proto_msgTypes,
	}.Build()
	File_omdb_proto = out.File
	file_omdb_proto_rawDesc = nil
	file_omdb_proto_goTypes = nil
	file_omdb_proto_depIdxs = nil
}
