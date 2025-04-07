package gofs_client

type PdfService interface {
	ToPdfByFileId(fileId, tempDir string) (file FileVo)
	ToPdfByStorePath(storePath, tempDir string) (file FileVo)
}

type PdfServiceImpl struct {
	client *Client
}

func newPdfService(client *Client) PdfService {
	return PdfServiceImpl{
		client: client,
	}
}

func (p PdfServiceImpl) ToPdfByFileId(fileId, tempDir string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "rpc/pdf/toPdfByFileId", map[string]any{
		"id":      fileId,
		"tempDir": tempDir,
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

func (p PdfServiceImpl) ToPdfByStorePath(storePath, tempDir string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "rpc/pdf/toPdfByStorePath", map[string]any{
		"storePath": storePath,
		"tempDir":   tempDir,
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
