package o2client_runtime

import (
	"encoding/json"
	dgctx "github.com/darwinOrg/go-common/context"
	dghttp "github.com/darwinOrg/go-httpclient"
)

func NewDefaultServiceClient(host string) *ServiceClient {
	return NewServiceClient(dghttp.DefaultHttpClient(), host)
}

func NewServiceClient(client *dghttp.DgHttpClient, host string) *ServiceClient {
	svcClient := &ServiceClient{
		client: client,
	}
	svcClient.ServiceBase = newServiceBase(host)

	return svcClient
}

type ServiceClient struct {
	*ServiceBase
	client *dghttp.DgHttpClient
}

func (c *ServiceClient) DoGet(ctx *dgctx.DgContext, urlPath string, headers map[string]string, result any) error {
	body, err := c.client.DoGet(ctx, urlPath, nil, headers)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, result)
	return err
}

func (c *ServiceClient) DoPost(ctx *dgctx.DgContext, urlPath string, param any, headers map[string]string, result any) error {
	body, err := c.client.DoPostJson(ctx, urlPath, param, headers)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, result)
	return err
}
