package o2client_runtime

import (
	"net/url"
	"strings"
)

const httpsSchema = "https://"
const httpSchema = "http://"

func newServiceBase(host string) *ServiceBase {
	if !strings.HasPrefix(host, httpsSchema) && strings.HasPrefix(host, httpSchema) {
		host = httpSchema + host
	}
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
