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
	components map[uint16]map[uint32]Component

	mutexCounter    sync.Mutex
	mutexSystem     sync.RWMutex
	mutexEntities   sync.RWMutex
	mutexComponents sync.RWMutex
}

func (w *World) Init() {
	w.mutexSystem.RLock()
	for _, system := range w.systems {
		system.RegisterWorld(w)
	}
	for _, system := range w.systems {
		system.Init()
	}
	w.mutexSystem.RUnlock()
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
	w.mutexSystem.Lock()
	w.systems = append(w.systems, sys)
	w.mutexSystem.Unlock()
}

func (w *World) addComponentFromEntity(entity *Entity, component Component) {
	w.mutexComponents.Lock()
	if _, has := w.components[component.GetComponentType()]; !has {
		w.components[component.GetComponentType()] = make(map[uint32]Component)
	}
	w.components[component.GetComponentType()][entity.Id] = component
	w.mutexComponents.Unlock()
}

func (w *World) removeComponentFromEntity(entity *Entity, componentId uint16) {
	w.mutexComponents.Lock()
	delete(w.components[componentId], entity.Id)
	w.mutexComponents.Unlock()
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
	if _, has := w.entities[entity.Id]; has {

		for componentType, _ := range entity.Components {
			w.mutexComponents.Lock()
			delete(w.components[componentType], entity.Id)
			w.mutexComponents.Unlock()
		}
		w.mutexEntities.Lock()
		delete(w.entities, entity.Id)
		w.mutexEntities.Unlock()
	}
}

func (w *World) GetEntitiesWithComponent(componentType ...uint16) []*Entity {
	entities := make([]*Entity, 0, 0)
	for _, tp := range componentType {
		w.mutexComponents.RLock()
		if components, has := w.components[tp]; has {
			for u, _ := range components {
				w.mutexEntities.RLock()
				if entity, has := w.entities[u]; has {
					entities = append(entities, entity)
				}
				w.mutexEntities.RUnlock()
			}
		}
		w.mutexComponents.RUnlock()
	}
	return entities
}

func (w *World) GetComponents(componentType uint16) map[uint32]Component {
	w.mutexComponents.RLock()
	components, has := w.components[componentType]
	w.mutexComponents.RUnlock()
	if !has {
		return make(map[uint32]Component, 0)
	}
	return components
}

func (w *World) GetComponentFromEntity(componentType uint16, entityId uint32) (Component, bool) {
	w.mutexComponents.RLock()
	component, has := w.components[componentType][entityId]
	w.mutexComponents.RUnlock()
	return component, has
}

func (w *World) AllEntities() map[uint32]*Entity {
	return w.entities
}

func NewWorld() *World {
	return &World{
		counter:         0,
		systems:         make([]System, 0),
		mutexCounter:    sync.Mutex{},
		mutexSystem:     sync.RWMutex{},
		mutexEntities:   sync.RWMutex{},
		mutexComponents: sync.RWMutex{},
		lastUpdate:      time.Now(),
		entities:        make(map[uint32]*Entity),
		components:      make(map[uint16]map[uint32]Component),
	}
}
