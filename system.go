package ecs_engine

import "time"

type System interface {
	Init()
	Update(duration time.Duration)
	RegisterWorld(world *World)
	HasRequirements(entity *Entity) bool
}
