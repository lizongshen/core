package rady

import (
	"fmt"
	"os"
	"strings"
)

const (
	ModeEnv         = "RADY_MODE"
	AutoRollbackEnv = "RADY_ROLLBACK"
	TestMod         = "test"
	AutoRollback    = "true"
)

func GetModeEnv() string {
	return os.Getenv(ModeEnv)
}

func IsTestMode() bool {
	return GetModeEnv() == TestMod
}

func ResetEnv(key string) {
	os.Setenv(key, "")
}

func GetConfigFileByMode(filePath string) string {
	mode := GetModeEnv()
	if mode == "" {
		return filePath
	}
	index := strings.LastIndexByte(filePath, os.PathSeparator)
	return fmt.Sprintf("%s%s.%s", filePath[:index+1], mode, filePath[index+1:])
}

func IsAutoRollback() bool {
	return os.Getenv(AutoRollbackEnv) == AutoRollback
}
