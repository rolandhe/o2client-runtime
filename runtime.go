package o2client_runtime

import "net/url"

func newServiceBase(host string) *ServiceBase {
	return &ServiceBase{
		host: host,
	}
}

type ServiceBase struct {
	host string
}

func (b *ServiceBase) JoinUrl(urlPath string) (string, error) {
	return url.JoinPath(b.host, urlPath)
}
