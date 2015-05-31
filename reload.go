package listener

import (
	"log"
	"os/exec"
)

type Reload struct {
}

// Snippet from 'http://nathanleclaire.com/blog/2014/08/17/automagical-deploys-from-docker-hub/'

func (m *Reload) Call(hubMsg HubMessage) {
	log.Println("received message to reload ...")
	out, err := exec.Command("../reload.sh").Output()
	if err != nil {
		log.Println("ERROR EXECUTING COMMAND IN RELOAD HANDLER!!")
		log.Println(err)
		return
	}
	log.Println("output of reload.sh is", string(out))
}

