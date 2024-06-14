package main

import ( 
    "fmt"
    "time"
    "os/exec"
    "os"
)

type GameOfLife struct { 
    generation int
    x          int
    y          int 
    cells      [][]bool
}

func NewGameOfLife(x, y int) GameOfLife { 
    cells := make([][]bool, y)
    for row := range cells { 
        cells[row] = make([]bool, x)

        for col := range cells[row] {
            cells[row][col] = false
        }
    }

    return GameOfLife{ 
        generation: 0,
        x: x,
        y: y,
        cells: cells,
    }
}

var dir = [][]int{ 
    {1, 0}, 
    {1, 1}, 
    {0, 1},
    {-1, 1},
    {-1, 0},
    {-1, -1},
    {0, -1},
    {1, -1},
}

func GetNeighbors(cells [][]bool, x, y int) int { 
    count := 0
    for _, d := range dir { 
        row := d[0] + y
        col := d[1] + x

        if row < 0 || row >= len(cells) { 
            continue;
        }

        if col < 0 || col >= len(cells[0]) { 
            continue;
        }

        if cells[row][col] { 
            count++
        }
    }

    return count
}

func (l *GameOfLife) PlayRound() { 
    newCells := make([][]bool, len(l.cells))
    for i := range newCells { 
        newCells[i] = make([]bool, len(l.cells[i]))

    }

    for y := range l.cells { 
        for x := range l.cells[y] { 
            nbs := GetNeighbors(l.cells, x, y)

            if l.cells[y][x] {
                if nbs < 2 || nbs > 3 {
                    newCells[y][x] = false
                } else {
                    newCells[y][x] = true
                }
            } else {
                if nbs == 3 {
                    newCells[y][x] = true
                } else {
                    newCells[y][x] = false
                }
            }
        }
    }

    l.cells = newCells
    l.generation++
}

func (l *GameOfLife) Display() { 
    cells := l.cells

    for _, row := range cells { 
        for _, cell := range row { 
            if cell { 
                fmt.Print("X")
            }else { 
                fmt.Print(" ")
            }
        }
        fmt.Print("\n")
    }
}

func main(){ 
    gof := NewGameOfLife(60, 40)

    gof.cells[5 + 15][1 + 50] = true
    gof.cells[6 + 15][3 + 50] = true
    gof.cells[7 + 15][0 + 50] = true
    gof.cells[7 + 15][1 + 50] = true
    gof.cells[7 + 15][4 + 50] = true
    gof.cells[7 + 15][5 + 50] = true
    gof.cells[7 + 15][6 + 50] = true
    
    for { 
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()

        fmt.Println("current generation: ", gof.generation)
        gof.Display()
        gof.PlayRound()

        time.Sleep(100 * time.Millisecond)
    }
}
