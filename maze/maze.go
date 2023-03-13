package maze

type Maze struct {
	cells [][]*Cell
	size  int
}

func NewMaze(size int) *Maze {
	cells := make([][]*Cell, size)
	for i := range cells {
		cells[i] = make([]*Cell, size)
	}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			cells[y][x] = NewCell()
		}
	}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			current := cells[y][x]

			if y > 0 {
				north := cells[y-1][x]
				BuildWall(current, north)
			}

			if y < size-1 {
				south := cells[y+1][x]
				BuildWall(current, south)
			}

			if x > 0 {
				west := cells[y][x-1]
				BuildWall(current, west)
			}

			if x < size-1 {
				east := cells[y][x+1]
				BuildWall(current, east)
			}
		}
	}

	this := &Maze{cells: cells, size: size}
	return this
}

func (this *Maze) String() string {
	str := "start\n+ - +"
	for i := 1; i < this.size; i++ {
		str += " - +"
	}
	str += "\n"

	for _, row := range this.cells {
		str += "| "
		for x, current := range row {
			str += cellString(current) + " "

			if x < this.size-1 {
				east := row[x+1]
				if WallExists(current, east) {
					str += "| "
				} else {
					str += "  "
				}
			} else {
				str += "|\n"
			}
		}
	}
	return str
}

func cellString(cell *Cell) string {
	switch cell.GetValue() {
	case -1:
		return "#"
	case 0:
		return " "
	default:
		return string(cell.GetValue())
	}
}
