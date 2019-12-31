package main

import (
	"ecs_engine"
	"sync"
	"time"
)

type MoveSystem struct {
	world               *ecs_engine.World
	transformComponents map[uint32]*Transform
}

func (m *MoveSystem) Init() {
}

func (m *MoveSystem) Update(duration time.Duration) {
	wg := sync.WaitGroup{}
	for _, transform := range m.transformComponents {
		wg.Add(1)
		go func(t *Transform) {
			t.x += float32(duration) / float32(time.Second)
			t.y += float32(duration) / float32(time.Second)
			wg.Done()
		}(transform)
	}
	wg.Wait()
}

func (m *MoveSystem) RegisterWorld(world *ecs_engine.World) {
	m.world = world
}

func (m *MoveSystem) RegisterEntity(entity *ecs_engine.Entity) {
	if component, has := entity.GetComponent(COMPONENT_TRANSFORM); has {
		m.transformComponents[entity.Id] = component.(*Transform)
	}
}

func (m *MoveSystem) DeRegisterEntity(entity *ecs_engine.Entity) {
	delete(m.transformComponents, entity.Id)
}

func (m *MoveSystem) HasRequirements(entity *ecs_engine.Entity) bool {
	_, has := entity.GetComponent(COMPONENT_TRANSFORM)
	return has
}

func NewMoveSystem() *MoveSystem {
	return &MoveSystem{
		transformComponents: make(map[uint32]*Transform),
	}
}
