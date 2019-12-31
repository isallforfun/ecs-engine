package main

import (
	"ecs_engine"
	"sync"
	"time"
)

type MoveSystem struct {
	world *ecs_engine.World
}

func (m *MoveSystem) Init() {
}

func (m *MoveSystem) Update(duration time.Duration) {
	wg := sync.WaitGroup{}
	for _, transform := range m.world.GetComponents(COMPONENT_TRANSFORM) {
		wg.Add(1)
		go func(t *Transform) {
			t.x += 1 //float32(duration) / float32(time.Second)
			t.y += 1 //float32(duration) / float32(time.Second)
			wg.Done()
		}(transform.(*Transform))
	}
	wg.Wait()
}

func (m *MoveSystem) RegisterWorld(world *ecs_engine.World) {
	m.world = world
}

func (m *MoveSystem) HasRequirements(entity *ecs_engine.Entity) bool {
	_, has := entity.GetComponent(COMPONENT_TRANSFORM)
	return has
}

func NewMoveSystem() *MoveSystem {
	return &MoveSystem{}
}
