package errors

import (
	"fmt"
	"os"
	"strings"
)

/*
VaultErrors - обрабатывает ошибки при работе с хранилищем секретов Vault.
*/
func VaultErrors(err error) {
	switch {
	case strings.Contains(err.Error(), "secret not found"):
		fmt.Println("project not found")
		os.Exit(1)
	case strings.Contains(err.Error(), "connection reset by peer"):
		fmt.Println("Storage connection unavailable. Please check your Vault configuration.")
		os.Exit(1)
	case strings.Contains(err.Error(), "Error making API request"):
		fmt.Println("Vault API bad request. Check your vault token.")
		os.Exit(1)
	default:
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
SpecErrors - обрабатывает ошибки при работе с моделями инфраструктуры.
*/
func SpecErrors(err error) {
	switch {
	case err.Error() == "unknown mode":
		fmt.Println("unknown mode")
		os.Exit(1)
	default:
		fmt.Println(err)
		os.Exit(1)
	}
}
