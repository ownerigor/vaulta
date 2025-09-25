package main

import (
	"github.com/ownerigor/vaulta/cmd"
	"github.com/ownerigor/vaulta/pkg/msg"
)

func main() {
	if err := cmd.Execute(); err != nil {
		msg.Die(err.Error())
	}
}
