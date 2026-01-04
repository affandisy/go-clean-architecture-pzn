package middleware

import (
	"go-clean-architecture-pzn/internal/entity"
	"go-clean-architecture-pzn/internal/model"
	"go-clean-architecture-pzn/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// func NewAuth(db *gorm.DB, log *logrus.Logger) fiber.Handler {
func NewAuth(userUserCase *usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// token := ctx.Get("Authorization", "NOT_FOUND")
		// log.Debugf("Authorization: %s", token)

		request := &model.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		userUserCase.Log.Debugf("Authorization: %s", request.Token)

		// user := new(entity.User)
		// err := db.Take(user, "token = ?", token).Error
		auth, err := userUserCase.Verify(ctx.UserContext(), request)
		if err != nil {
			// log.Warnf("Failed to find user by token: %+v", err)
			// return err
			userUserCase.Log.Warnf("Failed find user by token: %+v", err)
			return fiber.ErrUnauthorized
		}

		// log.Debugf("User: %+v", user)
		// userUserCase.Log.Debugf("User: %+v", user.ID)
		// ctx.Locals("user", user)

		userUserCase.Log.Debugf("User: %+v", auth.ID)
		ctx.Locals("auth", auth)

		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *entity.User {
	return ctx.Locals("user").(*entity.User)
}
