package main

type Entity struct {
	On    bool
	Ammo  int
	Power int
}

func (e *Entity) Shoot() bool {
	if e.On && e.Ammo > 0 {
		e.Ammo--
		return true
	}
	return false
}

func (e *Entity) RideBike() bool {
	if e.On && e.Power > 0 {
		e.Power--
		return true
	}
	return false
}

func main() {
	testStruct := new(Entity)
}
