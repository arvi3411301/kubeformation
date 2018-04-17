package main

import (
	"github.com/hasura/kubeformation/pkg/kubeformation/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}