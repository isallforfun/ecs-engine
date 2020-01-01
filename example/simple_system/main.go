package main

import (
	"github.com/isallforfun/ecs_engine"
)

const (
	COMPONENT_TRANSFORM uint16 = iota
)

func main() {
	world := ecs_engine.NewWorld()
	world.AddSystem(NewMoveSystem())
	world.AddSystem(NewLogSystem())
	player := world.GetEntity()
	player.AddComponent(NewTransform())

	player2 := world.GetEntity()
	player2.AddComponent(NewTransform())
	//ticker := time.NewTicker(1 * time.Second)
	world.Init()
	for {
		//<-ticker.C
		world.Update()
	}
}
