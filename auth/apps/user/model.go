package user

import (
	"github.com/infraboard/mcube/pb/resource"
	"golang.org/x/crypto/bcrypt"
)

const (
	AppName = "user"
)

type Service interface {
	RPCServer
}

func (cr *CreateUserRequest) MakePasswordHash() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(cr.Password), 10)
	if err != nil {
		panic(err)
	}
	cr.Password = string(bytes)
}

func NewUser(in *CreateUserRequest) *User {
	in.MakePasswordHash()
	return &User{
		Meta: resource.NewMeta(),
		Spec: in,
	}
}

func (u *User) CheckPassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Spec.Password), []byte(pw))
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{}
}

func NewDescribeUserRequest(name string) *DescribeUserRequest {
	return &DescribeUserRequest{
		Name: name,
	}
}
