package util

// GetRequest 定义了客户端向服务器发送的下载请求结构
// 用于在断点续传过程中告知服务器从哪个位置开始传输文件
type GetRequest struct {
	FileName   string `json:"file_name"`   // 要下载的文件名
	StartPoint int    `json:"start_point"` // 断点位置（已下载的字节数）
}

// PartFile 定义了文件分片的结构
// 在需要分片下载大文件时使用，当前项目中未实际使用此结构
type PartFile struct {
	FileName string `json:"file_name"` // 文件名
	Begin    int    `json:"begin"`     // 分片的起始位置（字节偏移量）
	End      int    `json:"end"`       // 分片的结束位置（字节偏移量）
}
