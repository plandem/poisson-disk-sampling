package poisson

import (
	"math"
)

type grid struct {
	cols     int
	rows     int
	cellSize float64
	points   []*Point
}

func newGrid(cols, rows int, cellSize float64) grid {
	return grid{
		cols:     cols,
		rows:     rows,
		cellSize: cellSize,
		points:   make([]*Point, cols*rows),
	}
}

func (grid grid) PointToCell(point *Point) (int, int) {
	col := (int)(float64(point.X) / grid.cellSize)
	row := (int)(float64(point.Y) / grid.cellSize)
	return col, row
}

func (grid grid) Cols() int {
	return grid.cols
}

func (grid grid) Rows() int {
	return grid.rows
}

func (grid grid) SetPoint(point *Point) {
	col, row := grid.PointToCell(point)
	offset := row*grid.cols + col
	grid.points[offset] = point
}

func (grid grid) IsNeighbourhood(point *Point, minDistance float64) bool {
	col, row := grid.PointToCell(point)

	//determine a neighborhood of cells around (x,y)
	colMin := int(math.Max(float64(col-2), 0))
	rowMin := int(math.Max(float64(row-2), 0))
	colMax := int(math.Min(float64(col+3), float64(grid.cols)))
	rowMax := int(math.Min(float64(row+3), float64(grid.rows)))

	//search around (x,y)
	for row := rowMin; row < rowMax; row++ {
		width := row * grid.cols

		for col := colMin; col < colMax; col++ {
			//check if the sample point exists on the grid and too close
			gridPoint := grid.points[width+col]
			if gridPoint != nil && gridPoint.Distance(point) < minDistance {
				return true
			}
		}
	}

	return false
}
