package gofs_client

import "fmt"

type FileService interface {
	DownloadFileByUrl(url string) (fileId string)
	DownloadFileByUrls(urls []string) (fileIds []string)
	UploadTempFile(data []byte, fileName, dir string, autoRename bool) (file FileVo)
	CopyFileToTemp(fileId, dir string, autoRename bool) (file FileVo)
	GetFile(fileId string) (file FileVo)
	GetFiles(fileIds []string) (files []FileVo)
	GetFilesByBusinessId(businessId, bucketCode string) (files []FileVo)
	GetFilesByBusinessIds(businessIds []string, bucketCode string) (files []File2Vo)
	GetFileBytesById(fileId string) (data []byte)
	GetFileBytesByStorePath(storePath string) (data []byte)
	BindBucket(fileId, businessId, bucketCode string)
	BindBuckets(fileIds []string, businessId, bucketCode string)
	DeleteByFileId(fileId string)
	DeleteByBusinessId(businessId, bucketCode string)
	CompressToZip(dir, zipName string, autoDelete bool) (file FileVo)
}

type FileServiceImpl struct {
	client *Client
}

func newFileService(client *Client) FileService {
	return FileServiceImpl{
		client: client,
	}
}

func (f FileServiceImpl) DownloadFileByUrl(url string) (fileId string) {
	result, err := httpPost[string](f.client, "file/downloadFileByUrl", map[string]any{
		"url": url,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	fileId = result.Data
	return
}

func (f FileServiceImpl) DownloadFileByUrls(urls []string) (fileIds []string) {
	result, err := httpPost[[]string](f.client, "file/downloadFileByUrls", map[string]any{
		"urls": urls,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	fileIds = result.Data
	return
}

func (f FileServiceImpl) UploadTempFile(data []byte, fileName, dir string, autoRename bool) (file FileVo) {
	result, err := upload[FileVo](f.client, "file/uploadFileToTemp", data, fileName, map[string]string{
		"dir":        dir,
		"autoRename": fmt.Sprint(autoRename),
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	file = result.Data
	return
}

func (f FileServiceImpl) CopyFileToTemp(fileId, dir string, autoRename bool) (file FileVo) {
	result, err := httpPost[FileVo](f.client, "file/copyFileToTempp", map[string]any{
		"fileId":     fileId,
		"dir":        dir,
		"autoRename": autoRename,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	file = result.Data
	return
}

func (f FileServiceImpl) GetFile(fileId string) (file FileVo) {
	result, err := httpPost[FileVo](f.client, "file/getFile", map[string]any{
		"id": fileId,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	file = result.Data
	return
}

func (f FileServiceImpl) GetFiles(fileIds []string) (files []FileVo) {
	result, err := httpPost[[]FileVo](f.client, "file/getFiles", map[string]any{
		"fileIds": fileIds,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	files = result.Data
	return
}

func (f FileServiceImpl) GetFilesByBusinessId(businessId, bucketCode string) (files []FileVo) {
	result, err := httpPost[[]FileVo](f.client, "file/getFilesByBusinessId", map[string]any{
		"businessId": businessId,
		"bucketCode": bucketCode,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	files = result.Data
	return
}

func (f FileServiceImpl) GetFilesByBusinessIds(businessIds []string, bucketCode string) (files []File2Vo) {
	result, err := httpPost[[]File2Vo](f.client, "file/getFilesByBusinessIds", map[string]any{
		"businessIds": businessIds,
		"bucketCode":  bucketCode,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	files = result.Data
	return
}

func (f FileServiceImpl) GetFileBytesById(fileId string) (data []byte) {
	data, err := download(f.client, "file/getFileBytesById", map[string]any{
		"id": fileId,
	})
	if err != nil {
		panic(err)
	}
	return
}

func (f FileServiceImpl) GetFileBytesByStorePath(storePath string) (data []byte) {
	data, err := download(f.client, "file/getFileBytesByStorePath", map[string]any{
		"storePath": storePath,
	})
	if err != nil {
		panic(err)
	}
	return
}

func (f FileServiceImpl) BindBucket(fileId, businessId, bucketCode string) {
	result, err := httpPost[any](f.client, "file/bindBucket", map[string]any{
		"fileId":     fileId,
		"businessId": businessId,
		"bucketCode": bucketCode,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (f FileServiceImpl) BindBuckets(fileIds []string, businessId, bucketCode string) {
	result, err := httpPost[any](f.client, "file/bindBuckets", map[string]any{
		"fileIds":    fileIds,
		"businessId": businessId,
		"bucketCode": bucketCode,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (f FileServiceImpl) DeleteByFileId(fileId string) {
	result, err := httpPost[any](f.client, "file/deleteById", map[string]any{
		"id": fileId,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
}

func (f FileServiceImpl) DeleteByBusinessId(businessId, bucketCode string) {
	result, err := httpPost[any](f.client, "file/deleteByBusinessId", map[string]any{
		"businessId": businessId,
		"bucketCode": bucketCode,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (f FileServiceImpl) CompressToZip(dir, zipName string, autoDelete bool) (file FileVo) {
	result, err := httpPost[FileVo](f.client, "file/compressToZip", map[string]any{
		"dir":        dir,
		"zipName":    zipName,
		"autoDelete": autoDelete,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	file = result.Data
	return
}
