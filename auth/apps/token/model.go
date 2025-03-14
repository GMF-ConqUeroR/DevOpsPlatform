package token

import (
	"auth/logger"
	"fmt"
	"time"

	"github.com/rs/xid"
)

func NewToken() *Token {
	return &Token{
		AccessToken:               xid.New().String(),
		RefreshToken:              xid.New().String(),
		AccessTokenExpiredSecond:  60 * 60,
		RefreshTokenExpiredSecond: 10 * 60 * 4,
		IssueAt:                   time.Now().Unix(),
	}
}

func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{}
}

func NewDescribeTokenRequest(ak string) *DescribeTokenRequest {
	return &DescribeTokenRequest{
		AccessToken: ak,
	}
}

func NewValidateTokenRequest(tk string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: tk,
	}
}

// 判断ak是否过期
func (t *Token) IssueTime() time.Time {
	return time.Unix(t.IssueAt, 0)
}

func (t *Token) AccessTokenExpiredTime() time.Time {
	return t.IssueTime().Add(time.Duration(t.AccessTokenExpiredSecond * int64(time.Second)))
}

func (t *Token) IsAcccessTokenExpired() bool {
	d := time.Since(t.AccessTokenExpiredTime())
	return d.Minutes() > 0
}

// 判断refresh ak是否过期
func (t *Token) IsRefreshTokenExpired() bool {
	d := time.Since(t.RefreshTokenExpiredTime())
	return d.Minutes() > 0
}

func (t *Token) RefreshTokenExpiredTime() time.Time {
	return t.IssueTime().Add(time.Duration(t.RefreshTokenExpiredSecond * int64(time.Second)))
}

// 判断用户是否有权限访问指定的API
func (t *Token) HavePermissionOrNot(svc, resource, act string) error {
	// 通过这种绝不可能条件
	logger.L().Debug().Msgf("%v,%v,%v", svc, resource, act)
	if t.Roles == nil && len(t.Roles.Items) > 0 {
		return nil
	}

	// 每个角色挨着判断
	for i := range t.Roles.Items {
		r := t.Roles.Items[i]
		if r.HavePermissionOrNot(svc, resource, act) {
			return nil
		}
	}

	return fmt.Errorf(" Permission Denial ")
}
