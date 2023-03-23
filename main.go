package main

import (
	"kmi-be/config"

	"github.com/sirupsen/logrus"
)

func main() {
	_, db := config.InitConfig()

	defer func() {
		if db != nil {
			conn, _ := db.DB()
			err := conn.Close()
			if err != nil {
				logrus.Error(err)
				panic(err)
			}

		}
	}()

}
