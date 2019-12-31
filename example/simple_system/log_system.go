package main

import (
	"ecs_engine"
	"log"
	"sync"
	"time"
)

type LogSystem struct {
	world *ecs_engine.World
}

func (l LogSystem) Init() {
}

func (l *LogSystem) Update(duration time.Duration) {
	wg := sync.WaitGroup{}
	for _, transform := range l.world.GetComponents(COMPONENT_TRANSFORM) {
		wg.Add(1)
		go func(t *Transform) {
			log.Println(t)
			wg.Done()
		}(transform.(*Transform))
	}
	wg.Wait()
}

func (l *LogSystem) RegisterWorld(world *ecs_engine.World) {
	l.world = world
}

func (l *LogSystem) HasRequirements(entity *ecs_engine.Entity) bool {
	_, has := entity.Components[COMPONENT_TRANSFORM]
	return has
}

func NewLogSystem() *LogSystem {
	return &LogSystem{}
}
