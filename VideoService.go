package gofs_client

type VideoService interface {
	cutImages(fileId string, count int) (urls []string)
}

type VideoServiceImpl struct {
	client *Client
}

func newVideoService(client *Client) VideoService {
	return VideoServiceImpl{
		client: client,
	}
}

func (v VideoServiceImpl) cutImages(fileId string, count int) (urls []string) {
	result, err := httpPost[[]string](v.client, "video/cutImages", map[string]any{
		"fileId": fileId,
		"count":  count,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	urls = result.Data
	return
}
