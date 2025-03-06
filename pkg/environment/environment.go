package environment

import "github.com/scienceandcode/habits-todo-backend/pkg/common"

func IsDevelopment() bool {
	return common.GetEnv("GIN_MODE") == "debug" && common.GetEnv("HTTP_SERVER_HANDLER") == "default"
}
