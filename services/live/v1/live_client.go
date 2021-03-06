package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v1/model"
)

type LiveClient struct {
	HcClient *http_client.HcHttpClient
}

func NewLiveClient(hcClient *http_client.HcHttpClient) *LiveClient {
	return &LiveClient{HcClient: hcClient}
}

func LiveClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//可单独创建直播播放域名或推流域名，每个租户最多可配置64条域名记录。
func (c *LiveClient) CreateDomain(request *model.CreateDomainRequest) (*model.CreateDomainResponse, error) {
	requestDef := GenReqDefForCreateDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainResponse), nil
	}
}

//将用户已创建的播放域名和推流域名建立域名映射关系
func (c *LiveClient) CreateDomainMapping(request *model.CreateDomainMappingRequest) (*model.CreateDomainMappingResponse, error) {
	requestDef := GenReqDefForCreateDomainMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainMappingResponse), nil
	}
}

//创建录制配置接口
func (c *LiveClient) CreateRecordConfig(request *model.CreateRecordConfigRequest) (*model.CreateRecordConfigResponse, error) {
	requestDef := GenReqDefForCreateRecordConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordConfigResponse), nil
	}
}

//禁止直播推流
func (c *LiveClient) CreateStreamForbidden(request *model.CreateStreamForbiddenRequest) (*model.CreateStreamForbiddenResponse, error) {
	requestDef := GenReqDefForCreateStreamForbidden()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStreamForbiddenResponse), nil
	}
}

//创建直播转码模板
func (c *LiveClient) CreateTranscodingsTemplate(request *model.CreateTranscodingsTemplateRequest) (*model.CreateTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForCreateTranscodingsTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTranscodingsTemplateResponse), nil
	}
}

//删除域名。只有在域名停用（off）状态时才能删除。
func (c *LiveClient) DeleteDomain(request *model.DeleteDomainRequest) (*model.DeleteDomainResponse, error) {
	requestDef := GenReqDefForDeleteDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainResponse), nil
	}
}

//将播放域名和推流域名的域名映射关系删除
func (c *LiveClient) DeleteDomainMapping(request *model.DeleteDomainMappingRequest) (*model.DeleteDomainMappingResponse, error) {
	requestDef := GenReqDefForDeleteDomainMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainMappingResponse), nil
	}
}

//删除录制配置接口
func (c *LiveClient) DeleteRecordConfig(request *model.DeleteRecordConfigRequest) (*model.DeleteRecordConfigResponse, error) {
	requestDef := GenReqDefForDeleteRecordConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordConfigResponse), nil
	}
}

//恢复直播推流接口
func (c *LiveClient) DeleteStreamForbidden(request *model.DeleteStreamForbiddenRequest) (*model.DeleteStreamForbiddenResponse, error) {
	requestDef := GenReqDefForDeleteStreamForbidden()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStreamForbiddenResponse), nil
	}
}

//删除直播转码模板
func (c *LiveClient) DeleteTranscodingsTemplate(request *model.DeleteTranscodingsTemplateRequest) (*model.DeleteTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTranscodingsTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTranscodingsTemplateResponse), nil
	}
}

//获取直播播放日志，基于域名以5分钟粒度进行打包，日志内容以 \"|\" 进行分隔。
func (c *LiveClient) ListLiveSampleLogs(request *model.ListLiveSampleLogsRequest) (*model.ListLiveSampleLogsResponse, error) {
	requestDef := GenReqDefForListLiveSampleLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveSampleLogsResponse), nil
	}
}

//查询直播中的流信息
func (c *LiveClient) ListLiveStreamsOnline(request *model.ListLiveStreamsOnlineRequest) (*model.ListLiveStreamsOnlineResponse, error) {
	requestDef := GenReqDefForListLiveStreamsOnline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLiveStreamsOnlineResponse), nil
	}
}

//查询录制配置接口
func (c *LiveClient) ListRecordConfigs(request *model.ListRecordConfigsRequest) (*model.ListRecordConfigsResponse, error) {
	requestDef := GenReqDefForListRecordConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordConfigsResponse), nil
	}
}

//查询禁播黑名单列表
func (c *LiveClient) ListStreamForbidden(request *model.ListStreamForbiddenRequest) (*model.ListStreamForbiddenResponse, error) {
	requestDef := GenReqDefForListStreamForbidden()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStreamForbiddenResponse), nil
	}
}

//查询直播加速的播流域名网络带宽监控数据
func (c *LiveClient) ShowBandwidth(request *model.ShowBandwidthRequest) (*model.ShowBandwidthResponse, error) {
	requestDef := GenReqDefForShowBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBandwidthResponse), nil
	}
}

//查询直播域名
func (c *LiveClient) ShowDomain(request *model.ShowDomainRequest) (*model.ShowDomainResponse, error) {
	requestDef := GenReqDefForShowDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainResponse), nil
	}
}

//查询加速的直播播放在线人数
func (c *LiveClient) ShowOnlineUsers(request *model.ShowOnlineUsersRequest) (*model.ShowOnlineUsersResponse, error) {
	requestDef := GenReqDefForShowOnlineUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOnlineUsersResponse), nil
	}
}

//查询直播加速的播流域名网络流量监控数据
func (c *LiveClient) ShowTraffic(request *model.ShowTrafficRequest) (*model.ShowTrafficResponse, error) {
	requestDef := GenReqDefForShowTraffic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrafficResponse), nil
	}
}

//查询直播转码模板
func (c *LiveClient) ShowTranscodingsTemplate(request *model.ShowTranscodingsTemplateRequest) (*model.ShowTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForShowTranscodingsTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTranscodingsTemplateResponse), nil
	}
}

//修改直播播放、RTMP推流加速域名相关信息
func (c *LiveClient) UpdateDomain(request *model.UpdateDomainRequest) (*model.UpdateDomainResponse, error) {
	requestDef := GenReqDefForUpdateDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainResponse), nil
	}
}

//修改禁推属性
func (c *LiveClient) UpdateStreamForbidden(request *model.UpdateStreamForbiddenRequest) (*model.UpdateStreamForbiddenResponse, error) {
	requestDef := GenReqDefForUpdateStreamForbidden()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStreamForbiddenResponse), nil
	}
}

//修改直播转码模板
func (c *LiveClient) UpdateTranscodingsTemplate(request *model.UpdateTranscodingsTemplateRequest) (*model.UpdateTranscodingsTemplateResponse, error) {
	requestDef := GenReqDefForUpdateTranscodingsTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTranscodingsTemplateResponse), nil
	}
}
