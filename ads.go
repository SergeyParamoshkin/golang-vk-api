package vkapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type TargetGroupResponse struct {
	ID int64 `json:"id"`
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

	err = json.Unmarshal(resp.Response, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

//func (client *VKClient) DeleteTargetGroup(AccountID string, TargetGroupID string) error {
//	v := url.Values{}
//	v.Add("account_id", AccountID)
//	v.Add("lifetime", TargetGroupID)
//	client.MakeRequest("ads.deleteTargetGroup", v)
//	return nil
//}

/// Возвращает количество обработанных контактов.
type ImportTargetContactsResponse struct {
	ID int64 `json:"response"`
}

func (client *VKClient) ImportTargetContacts(
	accountID string,
	targetGroupID string,
	contacts string) (ImportTargetContactsResponse, error) {
	d := ImportTargetContactsResponse{}
	v := url.Values{}
	v.Add("account_id", accountID)
	v.Add("target_group_id", targetGroupID)
	v.Add("contacts", contacts)
	resp, err := client.MakeRequest("ads.importTargetContacts", v)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(resp.Response, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}
