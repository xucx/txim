package api

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/imroc/req"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
)

const (
	API_HOST    = "console.tim.qq.com"
	API_VERSION = "v4"
)

var g struct {
	sdkAppID   int
	identifier string
	key        string
}

func Init(sdkAppID int, key, identifier string) {
	g.sdkAppID = sdkAppID
	g.key = key
	g.identifier = identifier
}

func UserSig(identifier string, expire int) string {
	sig, _ := tencentyun.GenSig(g.sdkAppID, g.key, identifier, expire)
	return sig
}

//单个账号导入
type AccountImportReq struct {
	Identifier string `json:"Identifier"` //:"test",
	Nick       string `json:"Nick"`       //:"test",
	FaceUrl    string `json:"FaceUrl"`    //:"http://www.qq.com"
}
type AccountImportAns struct{}

func AccountImport(r *AccountImportReq) (*AccountImportAns, error) {
	a := AccountImportAns{}
	err := Api("im_open_login_svc", "account_import", r, &a)
	return &a, err
}

//建群
//https://cloud.tencent.com/document/product/269/1615
type GroupCreateReq struct {
	OwnerAccount    string `json:"Owner_Account"`
	Type            string `json:"Type"`
	Name            string `json:"Name"`
	ApplyJoinOption string `json:"ApplyJoinOption"`
}
type GroupCreateAns struct {
	GroupId string `json:"GroupId"`
}

func GroupCreate(r *GroupCreateReq) (*GroupCreateAns, error) {
	a := GroupCreateAns{}
	err := Api("group_open_http_svc", "create_group", r, &a)
	return &a, err
}

//资料设置
type ProfileSetReq struct {
	FromAccount string               `json:"From_Account"`
	ProfileItem []*ProfileSetReqItem `json:"ProfileItem"`
}
type ProfileSetReqItem struct {
	Tag   string      `json:"Tag"`
	Value interface{} `json:"Value"`
}

type ProfileSetAns struct{}

func ProfileSet(r *ProfileSetReq) (*ProfileSetAns, error) {
	a := ProfileSetAns{}
	err := Api("profile", "portrait_set", r, &a)
	return &a, err
}

func Api(servicename, command string, in, out interface{}) error {
	host := fmt.Sprintf("https://%s/%s/%s/%s", API_HOST, API_VERSION, servicename, command)

	query := req.QueryParam{
		"sdkappid":    g.sdkAppID,
		"identifier":  g.identifier,
		"usersig":     UserSig(g.identifier, 5*60),
		"random":      rand.Uint32(),
		"contenttype": "json",
	}

	resp, err := req.Post(host, query, req.BodyJSON(in))
	if err != nil {
		return err
	}

	respBytes := resp.Bytes()

	errAns := struct {
		ActionStatus string `json:"ActionStatus"`
		ErrorInfo    string `json:"ErrorInfo"`
		ErrorCode    int    `json:"ErrorCode"`
	}{}

	if err := json.Unmarshal(respBytes, &errAns); err != nil {
		return err
	}

	if errAns.ErrorCode != 0 {
		return fmt.Errorf("ActionStatus:%s,ErrorInfo:%s,ErrorCode:%d", errAns.ActionStatus, errAns.ErrorInfo, errAns.ErrorCode)
	}

	if err := json.Unmarshal(respBytes, out); err != nil {
		return err
	}

	return nil
}
