package main

import (
	"fmt"
	"strings"
	"sync"

	obsws "github.com/christopher-dG/go-obs-websocket"
	streamdeck "github.com/magicmonkey/go-streamdeck"
	"github.com/magicmonkey/go-streamdeck/buttons"
	_ "github.com/magicmonkey/go-streamdeck/devices"
)

func main() {
	sd, err := streamdeck.New()
	if err != nil {
		panic(err)
	}

	obsClient := obsws.Client{Host: "localhost", Port: 4444}
	err = obsClient.Connect()
	if err != nil {
		panic(err)
	}

	sceneReq := obsws.NewGetSceneListRequest()
	scenes, err := sceneReq.SendReceive(obsClient)
	if err != nil {
		panic(err)
	}
	current_secene := strings.ToLower(scenes.CurrentScene)

	b := buttons.NewTextButton(current_secene)
	action := &OBSSceneAction{}
	b.SetActionHandler(action)

	sd.AddButton(1, buttons.NewTextButton(current_secene))

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

type OBSSceneAction struct{}

func (a *OBSSceneAction) Pressed(btn streamdeck.Button) {
	fmt.Println("Pressed")
}
