package gofs_client

import "fmt"

type FileService interface {
	DownloadFileByUrl(url string) (fileId string)
	DownloadFileByUrls(urls []string) (fileIds []string)
	UploadTempFile(data []byte, fileName, dir string, autoRename bool) (file FileVo)
	Upload(data []byte, fileName string) (file FileVo)
	CopyFileToTemp(fileId, dir string, autoRename bool) (file FileVo)
	GetFile(fileId string) (file FileVo)
	GetFiles(fileIds []string) (files []FileVo)
	GetFilesByBusinessId(businessId, bucketCode string) (files []FileVo)
	GetFilesByBusinessIds(businessIds []string, bucketCode string) (files []File2Vo)
	GetFileBytesById(fileId string) (data []byte)
	GetFileBytesByStorePath(storePath string) (data []byte)
	DeleteById(fileId string)
	DeleteByBusinessId(businessId, bucketCode string)
	BindBucket(fileId, businessId, bucketCode string)
	BindBuckets(fileIds []string, businessId, bucketCode string)
	BindBucketWithWater(fileId, businessId, bucketCode string, waterInfo WaterInfo)
	BindBucketsWithWater(fileIds []string, businessId, bucketCode string, waterInfo WaterInfo)
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
	result, err := httpPost[string](f.client, "rpc/downloadFileByUrl", map[string]any{
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
	result, err := httpPost[[]string](f.client, "rpc/downloadFileByUrls", map[string]any{
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
	result, err := upload[FileVo](f.client, "rpc/uploadTempFile", data, fileName, map[string]string{
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

func (f FileServiceImpl) Upload(data []byte, fileName string) (file FileVo) {
	result, err := upload[FileVo](f.client, "rpc/upload", data, fileName, map[string]string{})
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
	result, err := httpPost[FileVo](f.client, "rpc/copyFileToTemp", map[string]any{
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
	result, err := httpPost[FileVo](f.client, "rpc/getFile", map[string]any{
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
	result, err := httpPost[[]FileVo](f.client, "rpc/getFiles", map[string]any{
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
	result, err := httpPost[[]FileVo](f.client, "rpc/getFilesByBusinessId", map[string]any{
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
	result, err := httpPost[[]File2Vo](f.client, "rpc/getFilesByBusinessIds", map[string]any{
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
	data, err := download(f.client, "rpc/getFileBytesById", map[string]any{
		"id": fileId,
	})
	if err != nil {
		panic(err)
	}
	return
}

func (f FileServiceImpl) GetFileBytesByStorePath(storePath string) (data []byte) {
	data, err := download(f.client, "rpc/getFileBytesByStorePath", map[string]any{
		"storePath": storePath,
	})
	if err != nil {
		panic(err)
	}
	return
}

func (f FileServiceImpl) DeleteById(fileId string) {
	result, err := httpPost[any](f.client, "rpc/deleteById", map[string]any{
		"id": fileId,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (f FileServiceImpl) DeleteByBusinessId(businessId, bucketCode string) {
	result, err := httpPost[any](f.client, "rpc/deleteByBusinessId", map[string]any{
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

func (f FileServiceImpl) BindBucket(fileId, businessId, bucketCode string) {
	result, err := httpPost[any](f.client, "rpc/bindBucket", map[string]any{
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
	result, err := httpPost[any](f.client, "rpc/bindBuckets", map[string]any{
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

func (f FileServiceImpl) BindBucketWithWater(fileId, businessId, bucketCode string, waterInfo WaterInfo) {
	result, err := httpPost[any](f.client, "rpc/bindBucketWithWater", map[string]any{
		"fileId":     fileId,
		"businessId": businessId,
		"bucketCode": bucketCode,
		"waterInfo":  waterInfo,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (f FileServiceImpl) BindBucketsWithWater(fileIds []string, businessId, bucketCode string, waterInfo WaterInfo) {
	result, err := httpPost[any](f.client, "rpc/bindBucketsWithWater", map[string]any{
		"fileIds":    fileIds,
		"businessId": businessId,
		"bucketCode": bucketCode,
		"waterInfo":  waterInfo,
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
	result, err := httpPost[FileVo](f.client, "rpc/compressToZip", map[string]any{
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
