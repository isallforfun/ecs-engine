package ecs_engine_test

import (
	"ecs_engine"
	"testing"
	"time"
)

func TestWorld_GetEntity(t *testing.T) {
	world := ecs_engine.NewWorld()
	entity1 := world.GetEntity()
	entity2 := world.GetEntity()

	if entity1.Id != 1 {
		t.Error("entity1 with wrong id")
	}

	if entity2.Id != 2 {
		t.Error("entity2 with wrong id")
	}
}

func TestNewWorld(t *testing.T) {
	world := ecs_engine.NewWorld()
	if world == nil {
		t.Error("world not created")
	}
}

type SimpleSystem struct {
	count    int
	entities []*ecs_engine.Entity
}

func (s *SimpleSystem) Init() {
	s.count = 0
}

func (s *SimpleSystem) Update(duration time.Duration) {
	s.count++
}

func (s SimpleSystem) RegisterWorld(world *ecs_engine.World) {
}

func (s *SimpleSystem) RegisterEntity(entity *ecs_engine.Entity) {
	s.entities = append(s.entities, entity)
}

func (s *SimpleSystem) DeRegisterEntity(entity *ecs_engine.Entity) {
	entities := make([]*ecs_engine.Entity, 0, 0)
	for _, e := range s.entities {
		if e.Id != entity.Id {
			entities = append(entities, e)
		}
	}
	s.entities = entities
}

func (s SimpleSystem) HasRequirements(entity *ecs_engine.Entity) bool {
	return entity.HasComponent(COMPONENT_TYPE)
}

func TestWorld_AddSystem(t *testing.T) {
	world := ecs_engine.NewWorld()
	system := &SimpleSystem{}
	world.AddSystem(system)
	world.Init()

	if system.count != 0 {
		t.Error("system counter is not initialized")
	}

	world.Update()

	if system.count != 1 {
		t.Error("system update not work")
	}

	entity := world.GetEntity()
	entity.AddComponent(&SimpleComponent{BaseComponent: ecs_engine.BaseComponent{
		ComponentType: COMPONENT_TYPE,
	}})

	component, has := entity.GetComponent(COMPONENT_TYPE)
	if !has {
		t.Error("Entity not register the component")
	} else {
		if component.(*SimpleComponent).GetEntityId() != entity.Id {
			t.Error("Component don't have the entity id")
		}
	}

	if len(system.entities) != 1 {
		t.Error("system not get the entity")
	}

	entity.RemoveComponent(COMPONENT_TYPE)

	if len(system.entities) != 0 {
		t.Error("system not remove the entity")
	}

	world.RemoveEntity(entity)
	if len(world.AllEntities()) != 0 {
		t.Error("world not remove the entity")
	}
}

func TestWorld_RemoveEntity(t *testing.T) {
	world := ecs_engine.NewWorld()
	entity := world.GetEntity()

	if len(world.AllEntities()) != 1 {
		t.Error("entity not create on world")
	}

	world.RemoveEntity(entity)

	if len(world.AllEntities()) != 0 {
		t.Error("entity not removed from world")
	}
}

type SimpleComponent struct {
	ecs_engine.BaseComponent
}

const (
	COMPONENT_TYPE = 1
)

func TestWorld_GetEntitiesWithComponent(t *testing.T) {

	world := ecs_engine.NewWorld()
	entity := world.GetEntity()
	entity.AddComponent(&SimpleComponent{
		BaseComponent: ecs_engine.BaseComponent{
			ComponentType: COMPONENT_TYPE,
		},
	})

	entities := world.GetEntitiesWithComponent(COMPONENT_TYPE)

	if len(entities) != 1 {
		t.Error("Component not found")
	}

	entity.RemoveComponent(COMPONENT_TYPE)

	entities = world.GetEntitiesWithComponent(COMPONENT_TYPE)

	if len(entities) != 0 {
		t.Error("Entity component not Removed")
	}

	world.RemoveEntity(entity)

	entities = world.GetEntitiesWithComponent(COMPONENT_TYPE)

	if len(entities) != 0 {
		t.Error("Entity not Removed")
	}

}
