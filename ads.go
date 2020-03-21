package vkapi

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
)

type Ads struct {
	client *VKClient
}

func NewAds(client *VKClient) *Ads {
	return &Ads{
		client: client,
	}
}

type CreateTargetGroupResponse struct {
	ID int `json:"id"`
}

// id (integer) — идентификатор аудитории;
// name (string) — название аудитории ретаргетинга;
// last_updated (integer) — дата и время последнего пополнения аудитории в формате Unixtime;
// is_audience (integer, 1) — 1, если группа является аудиторией (т.е.может быть пополнена при помощи пикселя);
// is_shared (integer, 1) — 1, если группа является копией (см. метод ads.shareTargetGroup);
// audience_count (integer) — приблизительное количество пользователей, входящих в аудиторию;
// lifetime (integer) — количество дней, через которое пользователи, добавляемые в аудиторию ретаргетинга, будут автоматически исключены из неё;
// file_source (integer, 1) — признак пополнения аудитории через файл;
// api_source (integer, 1) — признак пополнения аудитории через метод ads.importTargetContacts;
// lookalike_source (integer, 1) — признак аудитории, полученной при помощи Look-a-Like.
//
// pixel (string) — код для размещения на сайте рекламодателя. Возвращается, если параметр extended = 1 (только для старых групп).
// domain (string) — домен сайта, где размещен код учета пользователей (только для старых групп).
type TargetGroupResponse struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	LastUpdated     int    `json:"last_updated"`
	IsAudience      bool   `json:"is_audience"`
	IsShared        bool   `json:"is_shared"`
	AudienceCount   int    `json:"audience_count"`
	Lifetime        int    `json:"lifetime"`
	FileSource      bool   `json:"file_source"`
	ApiSource       bool   `json:"api_source"`
	LookalikeSource bool   `json:"lookalike_source"`
	Pixel           string `json:"pixel"`
	Domain          string `json:"domain"`
}

type GetClientResponse []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	DayLimit string `json:"day_limit"`
	AllLimit string `json:"all_limit"`
}

type GetAccountResponse struct {
	ID     int    `json:"account_id"`
	Type   string `json:"account_type"`
	Status int    `json:"account_status"`
	Role   string `json:"access_role"`
}

//
func (a *Ads) CreateTargetGroup(AccountID int, Lifetime string, Name string, ClientIDs ...int) ([]CreateTargetGroupResponse, error) {

	v := url.Values{}
	v.Add("account_id", strconv.Itoa(AccountID))
	v.Add("lifetime", Lifetime)
	v.Add("name", Name)

	create := func() (CreateTargetGroupResponse, error) {
		d := CreateTargetGroupResponse{}

		resp, err := a.client.MakeRequest("ads.createTargetGroup", v)
		if err != nil {
			return d, err
		}

		err = json.Unmarshal(resp.Response, &d)
		if err != nil {
			log.Println(err, string(resp.Response))
			return d, err
		}

		return d, nil
	}

	if len(ClientIDs) == 0 {
		d, err := create()
		if err != nil {
			return nil, err
		}

		return []CreateTargetGroupResponse{d}, nil
	}

	ds := make([]CreateTargetGroupResponse, len(ClientIDs))

	for i, clientID := range ClientIDs {
		v.Set("client_id", strconv.Itoa(clientID))

		d, err := create()
		if err != nil {
			return nil, err
		}

		ds[i] = d
	}

	return ds, nil
}

func (a *Ads) GetTargetGroup(AccountID int, ClientIDs ...int) ([]TargetGroupResponse, error) {
	d := []TargetGroupResponse{}
	v := url.Values{}
	v.Add("account_id", strconv.Itoa(AccountID))

	get := func() error {
		resp, err := a.client.MakeRequest("ads.getTargetGroups", v)
		if err != nil {
			return err
		}

		ret := []TargetGroupResponse{}

		err = json.Unmarshal(resp.Response, &ret)
		if err != nil {
			log.Println(err, string(resp.Response))
			return err
		}

		d = append(d, ret...)

		return nil
	}

	if len(ClientIDs) == 0 {
		if err := get(); err != nil {
			return d, err
		}

		return d, nil
	}

	for _, ClientID := range ClientIDs {
		v.Set("client_id", strconv.Itoa(ClientID))

		if err := get(); err != nil {
			return d, err
		}
	}

	return d, nil
}

func (a *Ads) UpdateTargetGroup(AccountID string, TargetGroupID int) (int, error) {
	d := 0
	return d, nil
}

//
func (a *Ads) DeleteTargetGroup(AccountID int, TargetGroupID int, ClientIDs ...int) ([]int, error) {
	v := url.Values{}
	v.Add("account_id", strconv.Itoa(AccountID))
	v.Add("target_group_id", strconv.Itoa(TargetGroupID))

	delete := func() (int, error) {
		resp, err := a.client.MakeRequest("ads.deleteTargetGroup", v)
		if err != nil {
			return 0, err
		}

		var d int

		err = json.Unmarshal(resp.Response, &d)
		if err != nil {
			return 0, err
		}

		return d, nil
	}

	if len(ClientIDs) == 0 {
		d, err := delete()
		if err != nil {
			return nil, err
		}

		return []int{d}, nil
	}

	ds := make([]int, len(ClientIDs))

	for i, clientID := range ClientIDs {
		v.Set("client_id", strconv.Itoa(clientID))

		d, err := delete()
		if err != nil {
			return nil, err
		}

		ds[i] = d
	}

	return ds, nil
}

// Возвращает количество обработанных контактов.
type ImportTargetContactsResponse struct {
	Response int `json:"response"`
}

// Examples:
// ImportTargetContacts
//
func (a *Ads) ImportTargetContacts(
	accountID int,
	targetGroupID int,
	contacts string) (ImportTargetContactsResponse, error) {
	d := ImportTargetContactsResponse{}
	v := url.Values{}
	v.Add("account_id", strconv.Itoa(accountID))
	v.Add("target_group_id", strconv.Itoa(targetGroupID))
	v.Add("contacts", contacts)
	resp, err := a.client.MakeRequest("ads.importTargetContacts", v)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(resp.Response, &d.Response)
	if err != nil {
		return d, err
	}

	return d, nil
}

// Examples:
// GetClients
//

func (a *Ads) GetClients(AccountID int) (GetClientResponse, error) {
	d := GetClientResponse{}

	v := url.Values{}
	v.Add("account_id", strconv.Itoa(AccountID))

	resp, err := a.client.MakeRequest("ads.getClients", v)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(resp.Response, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

func (a *Ads) GetAccounts() ([]GetAccountResponse, error) {
	resp, err := a.client.MakeRequest("ads.getAccounts", nil)
	if err != nil {
		return nil, err
	}

	ret := make([]GetAccountResponse, 0)

	err = json.Unmarshal(resp.Response, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
