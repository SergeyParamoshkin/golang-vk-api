package vkapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type TargetGroupResponse struct {
	Static        string `json:"static"`
	LoaderVersion string `json:"loaderVersion"`
	LangVersion   string `json:"langVersion"`
}

func (client *VKClient) CreateTargetGroup(AccountID string, lifetime string, name string) (TargetGroupResponse, error) {
	fmt.Println("CreateTargetGroup....")
	d := TargetGroupResponse{}
	v := url.Values{}
	v.Add("account_id", AccountID)
	v.Add("lifetime", lifetime)
	v.Add("name", name)

	resp, err := client.MakeRequest("ads.createTargetGroup", v)
	if err != nil {
		return d, err
	}

	json.Unmarshal(resp.Response, &d)
	fmt.Println(d)

	return d, nil
}
