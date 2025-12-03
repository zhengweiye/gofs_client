package gofs_client

import "sync"

type Client struct {
	http        string
	ip          string
	port        int
	contextPath string
	appKey      string
	appSecret   string
	env         string
}

var clientObj *Client
var clientOnce sync.Once

type Option struct {
	Http        string
	Ip          string
	Port        int
	ContextPath string
	AppKey      string
	AppSecret   string
	Env         string
}

func Create(opt Option) *Client {
	clientOnce.Do(func() {
		clientObj = &Client{
			http:        opt.Http,
			ip:          opt.Ip,
			port:        opt.Port,
			contextPath: opt.ContextPath,
			appKey:      opt.AppKey,
			appSecret:   opt.AppSecret,
			env:         opt.Env,
		}
	})
	return clientObj
}

func (c *Client) GetFileService() FileService {
	return newFileService(c)
}

func (c *Client) GetPdfService() PdfService {
	return newPdfService(c)
}

func (c *Client) GetImageService() ImageService {
	return newImageService(c)
}

func (c *Client) GetVideoService() VideoService {
	return newVideoService(c)
}

func (c *Client) GetQrcodeService() QrcodeService {
	return newQrcodeService(c)
}
