package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ListEnterpriseProjectRequest struct {
	Id      *string                              `json:"id,omitempty"`
	Limit   *int32                               `json:"limit,omitempty"`
	Name    *string                              `json:"name,omitempty"`
	Offset  int32                                `json:"offset"`
	SortDir *ListEnterpriseProjectRequestSortDir `json:"sort_dir,omitempty"`
	SortKey *ListEnterpriseProjectRequestSortKey `json:"sort_key,omitempty"`
	Status  *int32                               `json:"status,omitempty"`
}

func (o ListEnterpriseProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectRequest", string(data)}, " ")
}

type ListEnterpriseProjectRequestSortDir struct {
	value string
}

type ListEnterpriseProjectRequestSortDirEnum struct {
	DESC ListEnterpriseProjectRequestSortDir
	ASC  ListEnterpriseProjectRequestSortDir
}

func GetListEnterpriseProjectRequestSortDirEnum() ListEnterpriseProjectRequestSortDirEnum {
	return ListEnterpriseProjectRequestSortDirEnum{
		DESC: ListEnterpriseProjectRequestSortDir{
			value: "desc",
		},
		ASC: ListEnterpriseProjectRequestSortDir{
			value: "asc",
		},
	}
}

func (c ListEnterpriseProjectRequestSortDir) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListEnterpriseProjectRequestSortDir) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListEnterpriseProjectRequestSortKey struct {
	value string
}

type ListEnterpriseProjectRequestSortKeyEnum struct {
	CREATED_AT ListEnterpriseProjectRequestSortKey
	UPDATED_AT ListEnterpriseProjectRequestSortKey
}

func GetListEnterpriseProjectRequestSortKeyEnum() ListEnterpriseProjectRequestSortKeyEnum {
	return ListEnterpriseProjectRequestSortKeyEnum{
		CREATED_AT: ListEnterpriseProjectRequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListEnterpriseProjectRequestSortKey{
			value: "updated_at",
		},
	}
}

func (c ListEnterpriseProjectRequestSortKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListEnterpriseProjectRequestSortKey) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
