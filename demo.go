package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	txim "github.com/xucx/txim/api"
)

func main() {
	sdkAppId, _ := strconv.Atoi(os.Getenv("tximid"))
	sdkKey := os.Getenv("tximkey")
	txim.Init(sdkAppId, sdkKey, "administrator")

	// 生成签名
	// sig := txim.UserSig("administrator", 300)
	// fmt.Println("签名:", sig)

	var (
		data      interface{}
		dataBytes []byte
		err       error
	)

	//创建房间
	// data, err = txim.GroupCreate(&txim.GroupCreateReq{
	// 	OwnerAccount:    "administrator",
	// 	Type:            "ChatRoom",
	// 	Name:            "测试聊天室",
	// 	ApplyJoinOption: "FreeAccess",
	// })
	// dataBytes, _ = json.Marshal(data)
	// fmt.Printf("%s,%v\n", string(dataBytes), err)

	//创建用户
	// data, err = txim.AccountImport(&txim.AccountImportReq{
	// 	Identifier: "test",
	// 	Nick:       "test1",
	// })
	// dataBytes, _ = json.Marshal(data)
	// fmt.Printf("%s,%v\n", string(dataBytes), err)

	//查询用户
	data, err = txim.AccountCheck(&txim.AccountCheckReq{
		CheckItem: []*txim.AccountCheckReqItem{
			{UserID: "test"},
		},
	})
	dataBytes, _ = json.Marshal(data)
	fmt.Printf("%s,%v\n", string(dataBytes), err)

	//删除用户
	// data, err = txim.AccountDelete(&txim.AccountDeleteReq{
	// 	DeleteItem: []*txim.AccountDeleteReqItem{
	// 		{UserID: "test"},
	// 	},
	// })
	// dataBytes, _ = json.Marshal(data)
	// fmt.Printf("%s,%v\n", string(dataBytes), err)

	//设置用户资料
	// data, err = txim.ProfileSet(&txim.ProfileSetReq{
	// 	FromAccount: "test1",
	// 	ProfileItem: []*txim.ProfileSetReqItem{
	// 		{
	// 			Tag:   "Tag_Profile_IM_Nick",
	// 			Value: "test1",
	// 		},
	// 	},
	// })
	// dataStr, _ := json.Marshal(data)
	// fmt.Printf("%+v,%v\n", dataStr, err)

}
