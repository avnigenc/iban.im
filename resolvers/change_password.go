package resolvers

import (
	"context"
	"github.com/monocash/iban.im/config"

	"github.com/monocash/iban.im/handler"
	"github.com/monocash/iban.im/model"
	// "fmt"
)

// ChangePassword mutation change password
func (r *Resolvers) ChangePassword(ctx context.Context, args changePasswordMutationArgs) (*ChangePasswordResponse, error) {
	userID := ctx.Value(handler.ContextKey("UserID"))
	// fmt.Println("inside change password")
	// fmt.Println("User id :",userID)
	// fmt.Printf("ctx: %+v\n",ctx)

	if userID == nil {
		msg := "Not Authorized"
		return &ChangePasswordResponse{Status: false, Msg: &msg, User: nil}, nil
	}
	user := model.User{}

	if err := config.DB.First(&user, userID).Error; err != nil {
		msg := "Not existing user"
		return &ChangePasswordResponse{Status: false, Msg: &msg, User: nil}, nil
	}

	user.Password = args.Password
	user.HashPassword()

	if err := config.DB.Save(&user).Error; err != nil {
		msg := err.Error()
		return &ChangePasswordResponse{Status: false, Msg: &msg, User: nil}, err
	}
	return &ChangePasswordResponse{Status: true, Msg: nil, User: &UserResponse{u: &user}}, nil
}

type changePasswordMutationArgs struct {
	Password string
}

// ChangePasswordResponse is the response type
type ChangePasswordResponse struct {
	Status bool
	Msg    *string
	User   *UserResponse
}

// Ok for ChangePasswordResponse
func (r *ChangePasswordResponse) Ok() bool {
	return r.Status
}

// Error for ChangePasswordResponse
func (r *ChangePasswordResponse) Error() *string {
	return r.Msg
}
