package gofs_client

type ImageService interface {
	addWaterInfo(fileId, address string, longitude, latitude float64)
}

type ImageServiceImpl struct {
	client *Client
}

func newImageService(client *Client) ImageService {
	return ImageServiceImpl{
		client: client,
	}
}

func (i ImageServiceImpl) addWaterInfo(fileId, address string, longitude, latitude float64) {
	result, err := httpPost[any](i.client, "image/addWater", map[string]any{
		"fileId":    fileId,
		"address":   address,
		"longitude": longitude,
		"latitude":  latitude,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
}
