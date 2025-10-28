package gofs_client

type PdfService interface {
	OfficeToPdfByFileId(fileId, outputDir string) (file FileVo)
	OfficeToPdfByStorePath(storePath, outputDir string) (file FileVo)
	PdfToImgByFileId(fileId, outputDir string) (file FileVo)
	PdfToImgByStorePath(storePath, outputDir string) (file FileVo)
	ImgToPdfByFileId(fileId, outputDir string) (file FileVo)
	ImgToPdfByStorePath(storePath, outputDir string) (file FileVo)
	SplitPdfByFileId(fileId string) (pdfs []PdfVo)
	SplitPdfByStorePath(storePath string) (pdfs []PdfVo)
}

type PdfServiceImpl struct {
	client *Client
}

func newPdfService(client *Client) PdfService {
	return PdfServiceImpl{
		client: client,
	}
}

func (p PdfServiceImpl) OfficeToPdfByFileId(fileId, outputDir string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "pdf/officeToPdfByFileId", map[string]any{
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

func (p PdfServiceImpl) OfficeToPdfByStorePath(storePath, outputDir string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "pdf/officeToPdfByStorePath", map[string]any{
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

func (p PdfServiceImpl) PdfToImgByFileId(fileId, outputDir string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "pdf/pdfToImgByFileId", map[string]any{
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

func (p PdfServiceImpl) PdfToImgByStorePath(storePath, outputDir string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "pdf/pdfToImgByStorePath", map[string]any{
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

func (p PdfServiceImpl) ImgToPdfByFileId(fileId, outputDir string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "pdf/imgToPdfByFileId", map[string]any{
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

func (p PdfServiceImpl) ImgToPdfByStorePath(storePath, outputDir string) (file FileVo) {
	result, err := httpPost[FileVo](p.client, "pdf/imgToPdfByStorePath", map[string]any{
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

func (p PdfServiceImpl) SplitPdfByFileId(fileId string) (pdfs []PdfVo) {
	result, err := httpPost[[]PdfVo](p.client, "pdf/splitPdfByFileId", map[string]any{
		"id": fileId,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	pdfs = result.Data
	return
}

func (p PdfServiceImpl) SplitPdfByStorePath(storePath string) (pdfs []PdfVo) {
	result, err := httpPost[[]PdfVo](p.client, "pdf/splitPdfByStorePath", map[string]any{
		"storePath": storePath,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	pdfs = result.Data
	return
}
