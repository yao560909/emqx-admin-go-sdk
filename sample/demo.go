package sample

import (
	"context"
	"fmt"

	sdk "github.com/yao560909/emqx-admin-go-sdk"
	sdkcode "github.com/yao560909/emqx-admin-go-sdk/core"
	"github.com/yao560909/emqx-admin-go-sdk/service/dashboard"
)

func DashboardListUsers(client *sdk.Client) {
	ctx := context.Background()
	loginResp, err := client.Dashboard.Login(ctx, dashboard.NewLoginReqBuilder().Username("test").Password("1qaz@WSX").Builder())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	if loginResp.APIResp.StatusCode != 200 {
		return
	}
	token := loginResp.LoginResponse.Token
	h := map[string]string{
		"Authorization": "Bearer " + token,
	}
	optionFunc := sdkcode.WithHeaders(h)
	listUsersResp, err := client.Dashboard.ListUsers(ctx, dashboard.NewListUsersReqBuilder().Build(), optionFunc)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		println(listUsersResp)
	}
	client.Dashboard.Logout(ctx, dashboard.NewLogoutReqBuilder().Build())
}
