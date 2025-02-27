// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package service

import (
	"fmt"
	"time"

	"qingcloud-sdk-go/config"
	"qingcloud-sdk-go/request"
	"qingcloud-sdk-go/request/data"
	"qingcloud-sdk-go/request/errors"
)

var _ fmt.State
var _ time.Time

type MiscService struct {
	Config     *config.Config
	Properties *MiscServiceProperties
}

type MiscServiceProperties struct {
}

func (s *QingCloudService) Misc() (*MiscService, error) {
	properties := &MiscServiceProperties{}

	return &MiscService{Config: s.Config, Properties: properties}, nil
}

// Documentation URL: https://docs.qingcloud.com/product/api/action/misc/get_quota_left.html
func (s *MiscService) GetQuotaLeft(i *GetQuotaLeftInput) (*GetQuotaLeftOutput, error) {
	if i == nil {
		i = &GetQuotaLeftInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "GetQuotaLeft",
		RequestMethod: "GET",
	}

	x := &GetQuotaLeftOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type GetQuotaLeftInput struct {
	ResourceTypes []*string `json:"resource_types" name:"resource_types" location:"params"`
	Zone          *string   `json:"zone" name:"zone" location:"params"` // Required
}

func (v *GetQuotaLeftInput) Validate() error {

	if v.Zone == nil {
		return errors.ParameterRequiredError{
			ParameterName: "Zone",
			ParentName:    "GetQuotaLeftInput",
		}
	}

	return nil
}

type GetQuotaLeftOutput struct {
	Message      *string      `json:"message" name:"message"`
	Action       *string      `json:"action" name:"action" location:"elements"`
	QuotaLeftSet []*QuotaLeft `json:"quota_left_set" name:"quota_left_set" location:"elements"`
	RetCode      *int         `json:"ret_code" name:"ret_code" location:"elements"`
}

// Documentation URL: https://docs.qingcloud.com/product/api/action/misc
func (s *MiscService) GetResourceLimit(i *GetResourceLimitInput) (*GetResourceLimitOutput, error) {
	if i == nil {
		i = &GetResourceLimitInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "GetResourceLimit",
		RequestMethod: "GET",
	}

	x := &GetResourceLimitOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type GetResourceLimitInput struct {
	VolumeType *int    `json:"volume_type" name:"volume_type" location:"params"`
	Zone       *string `json:"zone" name:"zone" location:"params"` // Required
}

func (v *GetResourceLimitInput) Validate() error {

	if v.Zone == nil {
		return errors.ParameterRequiredError{
			ParameterName: "Zone",
			ParentName:    "GetResourceLimitInput",
		}
	}

	return nil
}

type GetResourceLimitOutput struct {
	Message        *string         `json:"message" name:"message"`
	Action         *string         `json:"action" name:"action" location:"elements"`
	MaxSize        *int            `json:"max-size" name:"max-size" location:"elements"`
	MinSize        *int            `json:"min-size" name:"min-size" location:"elements"`
	ResourceLimits *ResourceLimits `json:"resource_limits" name:"resource_limits" location:"elements"`
	RetCode        *int            `json:"ret_code" name:"ret_code" location:"elements"`
	Step           *int            `json:"step" name:"step" location:"elements"`
	VxNetSubnets   []*string       `json:"vxnet_subnets" name:"vxnet_subnets" location:"elements"`
	VxNetVersion   *int            `json:"vxnet_version" name:"vxnet_version" location:"elements"`
}
