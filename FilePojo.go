package gofs_client

type FileVo struct {
	FileId     string `json:"fileId"`
	FileName   string `json:"fileName"`
	FilePath   string `json:"filePath"`
	FileType   string `json:"fileType"`
	UploadTime string `json:"uploadTime"`
}

type File2Vo struct {
	BusinessId string `json:"businessId"`
	FileId     string `json:"fileId"`
	FileName   string `json:"fileName"`
	FilePath   string `json:"filePath"`
	FileType   string `json:"fileType"`
	UploadTime string `json:"uploadTime"`
}

type WaterInfo struct {
	Address   string  `json:"address"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type PdfVo struct {
	StorePath string `json:"storePath"`
	PageNum   int    `json:"pageNum"`
}
