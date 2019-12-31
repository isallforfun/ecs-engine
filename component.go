package ecs_engine

type BaseComponent struct {
	EntityId      uint32
	ComponentType uint16
}

func (b *BaseComponent) SetEntityId(id uint32) {
	b.EntityId = id
}

func (b *BaseComponent) GetEntityId() uint32 {
	return b.EntityId
}

func (b *BaseComponent) GetComponentType() uint16 {
	return b.ComponentType
}

type Component interface {
	SetEntityId(id uint32)
	GetEntityId() uint32
	GetComponentType() uint16
}
