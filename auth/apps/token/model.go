package token

import (
	"time"

	"github.com/rs/xid"
)

func NewToken() *Token {
	return &Token{
		AccessToken:               xid.New().String(),
		RefreshToken:              xid.New().String(),
		AccessTokenExpiredSecond:  10 * 60,
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
