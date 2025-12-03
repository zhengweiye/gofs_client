package gofs_client

type QrcodeService interface {
	Create(imageName, txt, dir string, width, height int) (url string)
	GetPath(imageName, dir string) (url string)
}

type QrcodeServiceImpl struct {
	client *Client
}

func newQrcodeService(client *Client) QrcodeService {
	return QrcodeServiceImpl{
		client: client,
	}
}

func (q QrcodeServiceImpl) Create(imageName, txt, dir string, width, height int) (url string) {
	result, err := httpPost[string](q.client, "qrCode/create", map[string]any{
		"imageName": imageName,
		"txt":       txt,
		"dir":       dir,
		"width":     width,
		"height":    height,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	url = result.Data
	return
}

func (q QrcodeServiceImpl) GetPath(imageName, dir string) (url string) {
	result, err := httpPost[string](q.client, "qrCode/getPath", map[string]any{
		"imageName": imageName,
		"dir":       dir,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	url = result.Data
	return
}
