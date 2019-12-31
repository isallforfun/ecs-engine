package main

import "ecs_engine"

type Transform struct {
	ecs_engine.BaseComponent
	x float32
	y float32
}

func NewTransform() *Transform {
	return &Transform{
		BaseComponent: ecs_engine.BaseComponent{ComponentType: COMPONENT_TRANSFORM},
		x:             0,
		y:             0,
	}
}
