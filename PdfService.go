package gofs_client

type PdfService interface {
	ToPdfByFileId(fileId string) (file FileVo)
	ToPdfByStorePath(storePath string) (file FileVo)
}

type PdfServiceImpl struct {
	client *Client
}

func newPdfService(client *Client) PdfService {
	return PdfServiceImpl{
		client: client,
	}
}

func (p PdfServiceImpl) ToPdfByFileId(fileId string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "rpc/pdf/toPdfByFileId", map[string]any{
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

func (p PdfServiceImpl) ToPdfByStorePath(storePath string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "rpc/pdf/toPdfByStorePath", map[string]any{
		"storePath": storePath,
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
