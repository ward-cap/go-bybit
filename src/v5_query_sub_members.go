package src

import (
	"context"
	"net/url"
)

type V5APISubUIDList struct {
	CommonV5Response `json:",inline"`
	Result           V5APISubUIDListResult `json:"result"`
}

type V5APISubUIDListResult struct {
	SubMembers []V5SubMemberInfo `json:"subMembers"`
}

type V5SubMemberInfo struct {
	UID         string `json:"uid"`         // UID субакаунта
	Username    string `json:"username"`    // Назва субакаунта
	MemberType  int    `json:"memberType"`  // Тип акаунта: 1 - звичайний, 6 - кастодіальний
	Status      int    `json:"status"`      // Статус акаунта: 1 - активний
	Remark      string `json:"remark"`      // Примітка до акаунта
	AccountMode int    `json:"accountMode"` // Режим акаунта: 1 - класичний, 3 - Unified
}

func (s *V5UserService) GetSubUIDList(ctx context.Context) (res *V5APISubUIDList, err error) {
	err = s.client.getV5PrivatelyCtx(ctx, "/v5/user/query-sub-members", url.Values{}, &res)

	return
}
