/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type GlanceListImagesResponse struct {
	// 查询首页的URL。
	First *string `json:"first,omitempty"`
	// 资源类型。
	Images []GlanceShowImageResponseBody `json:"images,omitempty"`
	// 描述镜像列表模式的URL。
	Schema *string `json:"schema,omitempty"`
	// 查询下一页的URL。当查询镜像列表最后一页时，不存在next。
	Next *string `json:"next,omitempty"`
}

func (o GlanceListImagesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"GlanceListImagesResponse", string(data)}, " ")
}
