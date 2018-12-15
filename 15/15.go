package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type entityType byte

const (
	entityWall   entityType = '#'
	entityElf    entityType = 'E'
	entityGoblin entityType = 'G'
)

type entity struct {
	typ    entityType
	row    int
	col    int
	health int
	power  int
	dead   bool
}

func newEntity(typ entityType, row, col int) *entity {
	switch typ {
	case entityWall:
	case entityElf:
	case entityGoblin:
	default:
		return nil
	}
	e := &entity{
		typ: typ,
		row: row,
		col: col,
	}
	if typ == entityElf || typ == entityGoblin {
		e.health = 6
		e.power = 3
	}
	return e
}

func (e *entity) String() string {
	if e == nil {
		return "."
	}
	return string(e.typ)
}

func (e *entity) inRangeOf(target *entity) bool {
	if e.row == target.row {
		return e.col == target.col-1 || e.col == target.col+1
	}
	if e.col == target.col {
		return e.row == target.row-1 || e.row == target.row+1
	}
	return false
}

type world struct {
	grid [][]*entity
}

func (w *world) String() string {
	var sb strings.Builder
	for _, row := range w.grid {
		for _, e := range row {
			sb.WriteString(e.String())
		}
		sb.WriteByte('\n')
	}
	for _, row := range w.grid {
		for _, e := range row {
			if e == nil {
				continue
			}
			if e.typ == entityElf || e.typ == entityGoblin {
				sb.WriteString(fmt.Sprintf("%s: %d\n", e.String(), e.health))
			}
		}
	}
	return sb.String()
}

func readWorld(r io.Reader) (*world, error) {
	var w world
	scanner := bufio.NewScanner(r)
	row := 0
	for scanner.Scan() {
		var line []*entity
		for col, r := range scanner.Text() {
			line = append(line, newEntity(entityType(r), row, col))
		}
		w.grid = append(w.grid, line)
		row++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &w, nil
}

// doRound performs one round of combat and reports whether it ran to
// completion.
func (w *world) doRound() (completed bool) {
	priority := w.determinePriority()
	for _, e := range priority {
		if e.dead {
			continue
		}
		if !w.doTurn(e) {
			return false
		}
	}
	return true
}

func (w *world) determinePriority() []*entity {
	var priority []*entity
	for _, row := range w.grid {
		for _, e := range row {
			if e == nil {
				continue
			}
			if e.typ == entityElf || e.typ == entityGoblin {
				priority = append(priority, e)
			}
		}
	}
	return priority
}

func (w *world) doTurn(e *entity) (foundTargets bool) {
	targets := w.findTargets(e)
	if len(targets) == 0 {
		return false
	}

	// Check for in-range targets to attack.
	for _, target := range targets {
		if !e.inRangeOf(target) {
			continue
		}
		target.health -= e.power
		if target.health <= 0 {
			w.killEntity(target)
		}
		return true
	}

	// Find all open squares in range of a target.
	emptySquares := make(map[position]int)
	for _, target := range targets {
		empty := w.emptyAdjacentPositions(target)
		for _, pos := range empty {
			emptySquares[pos] = 0
		}
	}

	return true
}

func (w *world) findTargets(e *entity) []*entity {
	var targets []*entity
	var enemyType entityType
	switch e.typ {
	case 'E':
		enemyType = entityGoblin
	case 'G':
		enemyType = entityElf
	default:
		panic(fmt.Sprintf("invalid entity type: %c", e.typ))
	}
	for _, row := range w.grid {
		for _, ent := range row {
			if ent == nil {
				continue
			}
			if ent.typ == enemyType {
				targets = append(targets, ent)
			}
		}
	}
	return targets
}

// position is (row, col).
type position [2]int

func (w *world) emptyAdjacentPositions(e *entity) []position {
	var empty []position
	for _, pos := range []position{
		{e.row - 1, e.col},
		{e.row, e.col - 1},
		{e.row, e.col + 1},
		{e.row + 1, e.col},
	} {
		if w.grid[pos[0]][pos[1]] == nil {
			empty = append(empty, pos)
		}
	}
	return empty
}

func (w *world) killEntity(e *entity) {
	e.dead = true
	for r, row := range w.grid {
		for c, ent := range row {
			if e == ent {
				w.grid[r][c] = nil
				return
			}
		}
	}
}

func main() {
	w, err := readWorld(strings.NewReader(world0))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(w)
	fmt.Scanln()
	for w.doRound() {
		fmt.Println(w)
		fmt.Scanln()
	}
}

var world0 = `#####
#.G.#
#...#
#E..#
#####`
