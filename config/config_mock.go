package config

func NewMockConfig() Config {
	return Config{
		SecretKeyJwt:              "abc",
		TokenCurrentUserId:        "current_user_id",
		TokenCurrentUserRole:      "current_user_role",
		TokenExp:                  "",
		ConnectionStr:             "",
		InvalidTodoStatusArgument: "Invalid todo status.",
	}
}
