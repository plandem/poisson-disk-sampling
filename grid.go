package poisson

import (
	"math"
)

//Grid is a structure that holds information about points inside of grid
type Grid struct {
	cols     int
	rows     int
	cellSize float64
	points   []*Point
}

//NewGrid returns Grid with cols x rows dimension and size of cell equals to cellSize
func NewGrid(cols, rows int, cellSize float64) *Grid {
	return &Grid{
		cols:     cols,
		rows:     rows,
		cellSize: cellSize,
		points:   make([]*Point, cols*rows),
	}
}

//PointToCell returns col and row of point
func (grid *Grid) PointToCell(point *Point) (int, int) {
	col := (int)(float64(point.X) / grid.cellSize)
	row := (int)(float64(point.Y) / grid.cellSize)
	return col, row
}

//Cols returns total number of cols in the grid
func (grid *Grid) Cols() int {
	return grid.cols
}

//Rows returns total number of rows in the grid
func (grid *Grid) Rows() int {
	return grid.rows
}

//SetPoint stores point information in the grid. Point must fit to dimension of grid
func (grid *Grid) SetPoint(point *Point) {
	col, row := grid.PointToCell(point)
	offset := row*grid.cols + col
	grid.points[offset] = point
}

//GetPoint returns point information from cell with col and row in the grid
func (grid *Grid) GetPoint(col, row int) *Point {
	offset := row*grid.cols + col
	return grid.points[offset]
}

//IsNeighbourhood checks if point has any neighbourhood point with distance less than minDistance
func (grid *Grid) IsNeighbourhood(point *Point, minDistance float64) bool {
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
