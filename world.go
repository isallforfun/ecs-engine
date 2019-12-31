package ecs_engine

import (
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

type World struct {
	counter    uint32
	systems    []System
	lastUpdate time.Time
	entities   map[uint32]*Entity

	mutexCounter  sync.Mutex
	mutexSystem   sync.RWMutex
	mutexEntities sync.RWMutex
}

func (w *World) Init() {
	for _, system := range w.systems {
		system.RegisterWorld(w)
	}
	for _, system := range w.systems {
		system.Init()
	}
}

func (w *World) Update() {
	now := time.Now()
	timePass := now.Sub(w.lastUpdate)
	for _, system := range w.systems {
		system.Update(timePass)
	}
	w.lastUpdate = now
}

func (w *World) AddSystem(system interface{}) {
	sys := reflect.ValueOf(system).Interface().(System)
	w.systems = append(w.systems, sys)
}

func (w *World) updateRegistry(entity *Entity) {
	for _, system := range w.systems {
		if system.HasRequirements(entity) {
			system.RegisterEntity(entity)
		} else {
			system.DeRegisterEntity(entity)
		}
	}
}

func (w *World) GetEntity() *Entity {
	var id uint32
	w.mutexCounter.Lock()
	atomic.AddUint32(&w.counter, 1)
	id = w.counter
	w.mutexCounter.Unlock()
	w.mutexEntities.Lock()
	defer w.mutexEntities.Unlock()
	w.entities[id] = &Entity{Id: id, Components: make(map[uint16]Component), World: w}
	return w.entities[id]
}

func (w *World) RemoveEntity(entity *Entity) {
	w.mutexEntities.Lock()
	if _, has := w.entities[entity.Id]; has {
		delete(w.entities, entity.Id)
		for _, system := range w.systems {
			system.DeRegisterEntity(entity)
		}
	}
	w.mutexEntities.Unlock()
}

func (w *World) GetEntitiesWithComponent(componentType ...uint16) []*Entity {
	entities := make([]*Entity, 0, 0)
	for _, entity := range w.entities {
		hasAllComponent := true
		for _, ct := range componentType {
			if _, has := entity.GetComponent(ct); !has {
				hasAllComponent = false
				break
			}
		}
		if hasAllComponent {
			entities = append(entities, entity)
		}
	}
	return entities
}

func (w *World) AllEntities() map[uint32]*Entity {
	return w.entities
}

func NewWorld() *World {
	return &World{
		counter:      0,
		systems:      make([]System, 0),
		mutexCounter: sync.Mutex{},
		mutexSystem:  sync.RWMutex{},
		lastUpdate:   time.Now(),
		entities:     make(map[uint32]*Entity),
	}
}
