package utils

import "os/user"

func GetUserDir() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	return currentUser.HomeDir
}
