package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	targetUser, exist := users[name]
	if !exist {
		return false, errors.New("not found")
	}
	if !targetUser.scheduledForDeletion {
		return false, nil
	}
	delete(users, name)
	return true, nil
}

type user struct {
	name	string
	number	int
	scheduledForDeletion bool
}
