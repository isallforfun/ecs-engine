package ecs_engine

import "time"

type System interface {
	Init()
	Update(duration time.Duration)
	RegisterWorld(world *World)
	RegisterEntity(entity *Entity)
	DeRegisterEntity(entity *Entity)
	HasRequirements(entity *Entity) bool
}
