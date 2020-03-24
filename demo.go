package main

import (
	"fmt"

	txim "github.com/xucx/txim/api"
)

func main() {
	txim.Init(1, "xxx", "xxx")

	//生成签名
	sig := txim.UserSig("xxx", 300)
	fmt.Println("签名:", sig)

	//创建房间
	ans, err := txim.GroupCreate(&txim.GroupCreateReq{
		OwnerAccount:    "administrator",
		Type:            "ChatRoom",
		Name:            "测试聊天室",
		ApplyJoinOption: "FreeAccess",
	})
	fmt.Printf("创建房间结果：%x,%+v\n", err, ans)

}
