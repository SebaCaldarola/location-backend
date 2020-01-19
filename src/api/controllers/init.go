package controllers

import (
	"fmt"
	"github.com/sebastian/location-backend/src/api/core"
)

var (
	cc *core.Core
)

func init(){
	var err error
	cc, err = core.NewCore()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = cc.InitDbTables()
	if err != nil {
		fmt.Println("error migrating :" + err.Error())
		panic(err)
	}
}

