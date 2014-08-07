package template

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file_2 := &embedded.EmbeddedFile{
		Filename:    `404.html`,
		FileModTime: time.Unix(1404142342, 0),
		Content:     string([]byte{0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x20, 0x7d, 0x7d, 0x50, 0x61, 0x67, 0x65, 0x20, 0x4e, 0x6f, 0x74, 0x20, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x20, 0xe2, 0x80, 0xa2, 0x20, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0xa, 0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x20, 0x7d, 0x7d, 0xa, 0x3c, 0x68, 0x31, 0x3e, 0x54, 0x68, 0x65, 0x20, 0x70, 0x61, 0x67, 0x65, 0x20, 0x79, 0x6f, 0x75, 0x20, 0x61, 0x72, 0x65, 0x20, 0x6c, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x61, 0x6e, 0x27, 0x74, 0x20, 0x62, 0x65, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x21, 0x3c, 0x2f, 0x68, 0x31, 0x3e, 0xa, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0xa, 0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x20, 0x7d, 0x7d, 0xa, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file_3 := &embedded.EmbeddedFile{
		Filename:    `base.html`,
		FileModTime: time.Unix(1404685144, 0),
		Content:     string([]byte{0x3c, 0x21, 0x44, 0x4f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x3d, 0x22, 0x65, 0x6e, 0x22, 0x3e, 0xa, 0x9, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x3d, 0x22, 0x75, 0x74, 0x66, 0x2d, 0x38, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3d, 0x22, 0x52, 0x79, 0x61, 0x6e, 0x20, 0x48, 0x61, 0x72, 0x74, 0x65, 0x72, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x7b, 0x7b, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x20, 0x2e, 0x20, 0x7d, 0x7d, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0xa, 0xa, 0x9, 0x9, 0x3c, 0x21, 0x2d, 0x2d, 0x20, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x20, 0x2d, 0x2d, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x2f, 0x63, 0x64, 0x6e, 0x2e, 0x6a, 0x73, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x35, 0x2e, 0x32, 0x2e, 0x33, 0x2f, 0x63, 0x73, 0x73, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x73, 0x73, 0x22, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x63, 0x73, 0x73, 0x22, 0x20, 0x2f, 0x3e, 0xa, 0xa, 0xa, 0x9, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x9, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0xa, 0x9, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x72, 0x6f, 0x77, 0x22, 0x3e, 0xa, 0x9, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x2d, 0x31, 0x32, 0x20, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0x3e, 0xa, 0xa, 0x9, 0x7b, 0x7b, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x20, 0x2e, 0x20, 0x7d, 0x7d, 0xa, 0xa, 0x9, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x9, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0xa, 0x9, 0x7b, 0x7b, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x22, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x20, 0x2e, 0x20, 0x7d, 0x7d, 0xa, 0xa, 0x9, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file_4 := &embedded.EmbeddedFile{
		Filename:    `list_libraries.html`,
		FileModTime: time.Unix(1407419126, 0),
		Content:     string([]byte{0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x20, 0x7d, 0x7d, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x20, 0xe2, 0x80, 0xa2, 0x20, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0xa, 0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x3c, 0x68, 0x31, 0x3e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x3c, 0x2f, 0x68, 0x31, 0x3e, 0xa, 0x9, 0x3c, 0x75, 0x6c, 0x3e, 0xa, 0x9, 0x9, 0x7b, 0x7b, 0x20, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x3c, 0x6c, 0x69, 0x3e, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x7b, 0x2e, 0x49, 0x64, 0x7d, 0x7d, 0x22, 0x3e, 0x7b, 0x7b, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x3c, 0x2f, 0x61, 0x3e, 0x3c, 0x2f, 0x6c, 0x69, 0x3e, 0xa, 0x9, 0x9, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x3c, 0x2f, 0x75, 0x6c, 0x3e, 0xa, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0xa, 0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x20, 0x7d, 0x7d, 0xa, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file_5 := &embedded.EmbeddedFile{
		Filename:    `movies_list.html`,
		FileModTime: time.Unix(1407417646, 0),
		Content:     string([]byte{0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x20, 0x7d, 0x7d, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x20, 0xe2, 0x80, 0xa2, 0x20, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0xa, 0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x3c, 0x68, 0x31, 0x3e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x3c, 0x2f, 0x68, 0x31, 0x3e, 0xa, 0x9, 0x3c, 0x75, 0x6c, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x2d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2d, 0x67, 0x72, 0x69, 0x64, 0x2d, 0x35, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x7b, 0x7b, 0x20, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x3c, 0x6c, 0x69, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x2f, 0x7b, 0x7b, 0x2e, 0x49, 0x64, 0x7d, 0x7d, 0x22, 0x3e, 0x3c, 0x69, 0x6d, 0x67, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x2f, 0x70, 0x2f, 0x77, 0x33, 0x34, 0x32, 0x2f, 0x7b, 0x7b, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x7d, 0x7d, 0x22, 0x2f, 0x3e, 0x3c, 0x2f, 0x61, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x70, 0x3e, 0x7b, 0x7b, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x7d, 0x7d, 0x3c, 0x2f, 0x70, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x2f, 0x6c, 0x69, 0x3e, 0xa, 0x9, 0x9, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x3c, 0x2f, 0x75, 0x6c, 0x3e, 0xa, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0xa, 0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x20, 0x7d, 0x7d, 0xa, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file_6 := &embedded.EmbeddedFile{
		Filename:    `movies_show.html`,
		FileModTime: time.Unix(1407417661, 0),
		Content:     string([]byte{0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x20, 0x7d, 0x7d, 0x7b, 0x7b, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x7d, 0x7d, 0x20, 0xe2, 0x80, 0xa2, 0x20, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0xa, 0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x3c, 0x68, 0x31, 0x3e, 0x7b, 0x7b, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x7d, 0x7d, 0x3c, 0x2f, 0x68, 0x31, 0x3e, 0xa, 0x9, 0x3c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3d, 0x22, 0x31, 0x32, 0x38, 0x30, 0x22, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3d, 0x22, 0x37, 0x32, 0x30, 0x22, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x22, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x2f, 0x7b, 0x7b, 0x2e, 0x49, 0x64, 0x7d, 0x7d, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x76, 0x6e, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x70, 0x65, 0x67, 0x55, 0x52, 0x4c, 0x22, 0x20, 0x2f, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x21, 0x2d, 0x2d, 0x20, 0x3c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x22, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x7b, 0x2e, 0x49, 0x44, 0x7d, 0x7d, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x6d, 0x70, 0x34, 0x22, 0x20, 0x2f, 0x3e, 0x20, 0x2d, 0x2d, 0x3e, 0xa, 0x9, 0x3c, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x3e, 0xa, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0xa, 0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x20, 0x7d, 0x7d, 0xa, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir_1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1407419126, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file_2, // 404.html
			file_3, // base.html
			file_4, // list_libraries.html
			file_5, // movies_list.html
			file_6, // movies_show.html

		},
	}

	// link ChildDirs
	dir_1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`pages`, &embedded.EmbeddedBox{
		Name: `pages`,
		Time: time.Unix(1407420022, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir_1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"404.html":            file_2,
			"base.html":           file_3,
			"list_libraries.html": file_4,
			"movies_list.html":    file_5,
			"movies_show.html":    file_6,
		},
	})
}
