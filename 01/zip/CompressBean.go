package main

// 压缩模式
const (
	CompressMode   = iota
	UnCompressMode = iota
)

var (
	TargetPath string
	Key        string
	SourcePath string
	// 默认是压缩模式
	mode = CompressMode
)

type Compress struct {
	SourcePath, TargetPath string
	Mode                   uint
	Key                    string
}
