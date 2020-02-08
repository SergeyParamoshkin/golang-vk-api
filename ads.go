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

//
func (a *Ads) CreateTargetGroup(AccountID string, lifetime string, name string) (CreateTargetGroupResponse, error) {
	d := CreateTargetGroupResponse{}
	v := url.Values{}
	v.Add("account_id", AccountID)
	v.Add("lifetime", lifetime)
	v.Add("name", name)

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

func (a *Ads) GetTargetGroup(AccountID string) ([]TargetGroupResponse, error) {
	d := []TargetGroupResponse{}
	v := url.Values{}
	v.Add("account_id", AccountID)

	resp, err := a.client.MakeRequest("ads.getTargetGroups", v)
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

func (a *Ads) UpdateTargetGroup(AccountID string, TargetGroupID int) (int, error) {
	d := 0
	return d, nil
}

//
func (a *Ads) DeleteTargetGroup(AccountID string, TargetGroupID int) (int, error) {
	d := 0
	v := url.Values{}
	v.Add("account_id", AccountID)
	v.Add("target_group_id", strconv.Itoa(TargetGroupID))
	resp, err := a.client.MakeRequest("ads.deleteTargetGroup", v)

	if err != nil {
		return d, err
	}

	err = json.Unmarshal(resp.Response, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

// Возвращает количество обработанных контактов.
type ImportTargetContactsResponse struct {
	Response int `json:"response"`
}

// Examples:
// ImportTargetContacts
//
func (a *Ads) ImportTargetContacts(
	accountID string,
	targetGroupID int,
	contacts string) (ImportTargetContactsResponse, error) {
	d := ImportTargetContactsResponse{}
	v := url.Values{}
	v.Add("account_id", accountID)
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
