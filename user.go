package main

import (
	"fmt"
	"strings"
)

type User struct {
	Id string
	NickName string
	Name string
	Email string
	AvatarURL string
}

// TODO Merge all these into one struct
type NickNameHolder struct {
	NickName string
}

type NameHolder struct {
	Name string
}

type FirstAndLastNameHolder struct {
	FirstName string
	LastName string
}

type EmailHolder struct {
	Email string
}

type AvatarURLHolder struct {
	AvatarURL string
}

func (user *User) CopyFrom(source interface{}) {
	if s, ok := source.(NickNameHolder); ok {
		user.NickName = s.NickName
	} else {
		logger.Print("source did not cast to NickNameHolder")
	}

	if s, ok := source.(NameHolder); ok && s.Name != "" {
		user.Name = s.Name
	}

	if s, ok := source.(FirstAndLastNameHolder); ok {
		if n := strings.TrimSpace(fmt.Sprintf("%s %s", s.FirstName, s.LastName)); n != "" {
			user.Name = n
		}
	}

	if s, ok := source.(EmailHolder); ok && s.Email != "" {
		user.Email = s.Email
	}

	if s, ok := source.(AvatarURLHolder); ok && s.AvatarURL != "" {
		user.AvatarURL = s.AvatarURL
	}
}
