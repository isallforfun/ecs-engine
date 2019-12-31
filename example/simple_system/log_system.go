package main

import (
	"ecs_engine"
	"log"
	"reflect"
	"sync"
	"time"
)

type LogSystem struct {
	world               *ecs_engine.World
	transformComponents map[uint32]*Transform
}

func (l LogSystem) Init() {
}

func (l LogSystem) Update(duration time.Duration) {
	wg := sync.WaitGroup{}
	for _, transform := range l.transformComponents {
		wg.Add(1)
		go func(t *Transform) {
			log.Println(t)
			wg.Done()
		}(transform)
	}
	wg.Wait()
}

func (l LogSystem) RegisterWorld(world *ecs_engine.World) {
	l.world = world
}

func (l LogSystem) RegisterEntity(entity *ecs_engine.Entity) {
	if _, has := entity.Components[COMPONENT_TRANSFORM]; has {
		l.transformComponents[entity.Id] = reflect.ValueOf(entity.Components[COMPONENT_TRANSFORM]).Interface().(*Transform)
	}
}

func (l LogSystem) DeRegisterEntity(entity *ecs_engine.Entity) {
	delete(l.transformComponents, entity.Id)
}

func (l LogSystem) HasRequirements(entity *ecs_engine.Entity) bool {
	_, has := entity.Components[COMPONENT_TRANSFORM]
	return has
}

func NewLogSystem() *LogSystem {
	return &LogSystem{
		transformComponents: make(map[uint32]*Transform),
	}
}
