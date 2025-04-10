package gofs_client

type PdfService interface {
	ToPdfByFileId(fileId, outputDir string) (file FileVo)
	ToPdfByStorePath(storePath, outputDir string) (file FileVo)
}

type PdfServiceImpl struct {
	client *Client
}

func newPdfService(client *Client) PdfService {
	return PdfServiceImpl{
		client: client,
	}
}

func (p PdfServiceImpl) ToPdfByFileId(fileId, outputDir string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "rpc/pdf/toPdfByFileId", map[string]any{
		"id":        fileId,
		"outputDir": outputDir,
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

func (p PdfServiceImpl) ToPdfByStorePath(storePath, outputDir string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "rpc/pdf/toPdfByStorePath", map[string]any{
		"storePath": storePath,
		"outputDir": outputDir,
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
