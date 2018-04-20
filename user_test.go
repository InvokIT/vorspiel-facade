package main

import (
  "testing"
)

func TestUserCopyFrom(t *testing.T) {
  newValues := struct{
	  NickName string
	  Name string
  }{
  	"TestUser",
  	"Test User",
  }

  user := &User{
  	"testUserId",
  	"OldNickName",
  	"OldName",
  	"OldEmail@example.com",
  	"http://example.com/oldavatar.png",
  }

  user.CopyFrom(newValues)
  if user.Id != "testUserId" {
  	t.Errorf("Id was changed, expected '%s' got '%s'.", "testUserId", user.Id)
  }
  if user.NickName != newValues.NickName {
  	t.Errorf("user.NickName != newValues.NickName, expected '%s' got '%s'", newValues.NickName, user.NickName)
  }
  if user.Name != newValues.Name {
	  t.Errorf("user.Name != newValues.Name, expected '%s' got '%s'", newValues.Name, user.Name)
  }
}
