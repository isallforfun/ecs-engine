package ecs_engine

import "sync/atomic"

type BaseComponent struct {
	EntityId      uint32
	ComponentType uint16
	Version       int64
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

func (b *BaseComponent) GetVersion() int64 {
	return b.Version
}

func (b *BaseComponent) UpdateVersion() {
	atomic.AddInt64(&b.Version, 1)
}

type Component interface {
	SetEntityId(id uint32)
	GetEntityId() uint32
	GetComponentType() uint16
	GetVersion() int64
	UpdateVersion()
}
