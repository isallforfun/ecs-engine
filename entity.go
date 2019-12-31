package ecs_engine

import "reflect"

type Entity struct {
	Id         uint32
	Components map[uint16]Component
	World      *World
}

func (e *Entity) AddComponent(component interface{}) {
	comp := component.(Component)
	e.Components[comp.GetComponentType()] = comp
	comp.SetEntityId(e.Id)
	e.World.updateRegistry(e)
}

func (e *Entity) RemoveComponent(componentType uint16) {
	delete(e.Components, componentType)
	e.World.updateRegistry(e)
}

func (e *Entity) GetComponent(componentType uint16) (interface{}, bool) {
	component, has := e.Components[componentType]
	if !has {
		return nil, has
	}
	return reflect.ValueOf(component).Interface(), has
}

func (e *Entity) HasComponent(componentType uint16) bool {
	_, has := e.Components[componentType]
	return has
}
