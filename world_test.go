package ecs_engine_test

import (
	"github.com/isallforfun/ecs_engine"
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
	world    *ecs_engine.World
}

func (s *SimpleSystem) OnEntityRemove(entity *ecs_engine.Entity) {
}

func (s *SimpleSystem) Init() {
	s.count = 0
}

func (s *SimpleSystem) Update(duration time.Duration) {
	s.count++
}

func (s *SimpleSystem) RegisterWorld(world *ecs_engine.World) {
	s.world = world
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

	_, has = entity.GetComponent(Not_EXISTS_COMPONENT_TYPE)
	if has {
		t.Error("Entity has not existed component")
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
	COMPONENT_TYPE            = 1
	Not_EXISTS_COMPONENT_TYPE = 2
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

	if !entity.HasComponent(COMPONENT_TYPE) {
		t.Error("component not saved")
	}

	components := world.GetComponents(COMPONENT_TYPE)

	if len(components) != 1 {
		t.Error("not found the components")
	}

	component, has := world.GetComponentFromEntity(COMPONENT_TYPE, entity.Id)
	if !has {
		t.Error("not found components")
	} else {
		if component == nil {
			t.Error("found the components")
		}

		if _, instanceof := component.(*SimpleComponent); !instanceof {
			t.Error("component not is instance of SimpleComponent")
		}
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

	components = world.GetComponents(COMPONENT_TYPE)
	if len(components) != 0 {
		t.Error("found the components")
	}
}

func TestWorld_RemoveEntity2(t *testing.T) {

	world := ecs_engine.NewWorld()
	entity := world.GetEntity()
	entity.AddComponent(&SimpleComponent{
		BaseComponent: ecs_engine.BaseComponent{
			ComponentType: COMPONENT_TYPE,
		},
	})

	world.RemoveEntity(entity)
	if len(world.AllEntities()) != 0 {
		t.Error("not remove the entity")
	}
}

func TestWorld_GetComponents_Empty(t *testing.T) {
	world := ecs_engine.NewWorld()

	if len(world.GetComponents(COMPONENT_TYPE)) != 0 {
		t.Error("found mystery component")
	}
}
