package maze

type Cell struct {
	val       int
	neighbors *Set
	walls     *Set
}

func NewCell() *Cell {
	this := &Cell{val: 0, neighbors: NewSet(), walls: NewSet()}
	return this
}

func Connect(c1, c2 *Cell) bool {
	if c1.CanAddNeighbor() && c2.CanAddNeighbor() {
		c1.AddNeighbor(c2)
		c2.AddNeighbor(c1)
		return true
	}
	return false
}

func WallExists(c1, c2 *Cell) bool {
	return c1.HasWall(c2) && c2.HasWall(c1)
}

func BuildWall(c1, c2 *Cell) bool {
	isConnected := true
	if !c1.HasNeighbor(c2) || !c2.HasNeighbor(c1) {
		isConnected = Connect(c1, c2)
	}

	if isConnected && c1.CanAddWall() && c2.CanAddWall() {
		c1.SetWall(c2)
		c2.SetWall(c1)
		return true
	}
	return false
}

func BreakWall(c1, c2 *Cell) bool {
	if c1.HasWall(c2) && c2.HasWall(c1) {
		c1.RemoveWall(c2)
		c2.RemoveWall(c1)
		return true
	}
	return false
}

func (this *Cell) CanAddNeighbor() bool {
	return this.neighbors.Size() < 4
}

func (this *Cell) CanAddWall() bool {
	return this.walls.Size() < 4
}

func (this *Cell) GetValue() int {
	return this.val
}

func (this *Cell) SetValue(val int) *Cell {
	this.val = val
	return this
}

func (this *Cell) GetNeighbors() []*Cell {
	nbors := make([]*Cell, this.neighbors.Size())
	for i, item := range this.neighbors.Values() {
		nbor, ok := item.(*Cell)
		if !ok {
			nbor = nil
		}
		nbors[i] = nbor
	}
	return nbors
}

func (this *Cell) GetWalls() []*Cell {
	walls := make([]*Cell, this.walls.Size())
	for i, item := range this.walls.Values() {
		wall, ok := item.(*Cell)
		if !ok {
			wall = nil
		}
		walls[i] = wall
	}
	return walls
}

func (this *Cell) AddNeighbor(nbor *Cell) bool {
	if this.CanAddNeighbor() {
		this.neighbors.Add(nbor)
		return true
	}
	return false
}

func (this *Cell) RemoveNeighbor(cell *Cell) *Cell {
	var nbor *Cell = nil
	if this.HasNeighbor(cell) {
		nbor = cell
		this.neighbors.Remove(cell)
	}

	if nbor != nil && this.HasWall(nbor) {
		this.walls.Remove(nbor)
	}
	return nbor
}

func (this *Cell) HasNeighbor(cell *Cell) bool {
	return this.neighbors.Contains(cell)
}

func (this *Cell) SetWall(nbor *Cell) bool {
	if !this.neighbors.Contains(nbor) || !this.CanAddWall() {
		return false
	}
	this.walls.Add(nbor)
	return true
}

func (this *Cell) AddWall(wall *Cell) bool {
	if !this.CanAddWall() || !this.AddNeighbor(wall) {
		return false
	}
	return this.SetWall(wall)
}

func (this *Cell) RemoveWall(cell *Cell) *Cell {
	var wall *Cell = nil
	if this.HasWall(cell) {
		wall = cell
		this.walls.Remove(cell)
	}
	return wall
}

func (this *Cell) HasWall(cell *Cell) bool {
	return this.walls.Contains(cell)
}

func (this *Cell) String() string {
	return string(this.val)
}
