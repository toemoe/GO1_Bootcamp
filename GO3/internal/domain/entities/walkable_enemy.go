package entities

import (
	"math/rand"
	"s21_rogue/internal/domain/value_objects"
)

type WalkableEnemyBehaviour interface {
	GetPosition() *value_objects.Position
	SetPosition(pos *value_objects.Position)
	CanAttack(position *value_objects.Position) bool
	IsChasing() bool
	SetChasing(bool)
	Walk(topLeft, botRight *value_objects.Position, busyPos []*value_objects.Position)
	IsVisible() bool
}

type walkableEnemy struct {
	Position  *value_objects.Position
	isChasing bool
	walkFn    func(topLeft, botRight *value_objects.Position, busyPos []*value_objects.Position)
}

func NewSimpleWalkableEnemy(position *value_objects.Position) WalkableEnemyBehaviour {
	we := &walkableEnemy{Position: position, isChasing: false}
	we.walkFn = we.defaultWalk
	return we
}

func NewGhostWalkableEnemy(position *value_objects.Position) WalkableEnemyBehaviour {
	we := &ghostWalkableEnemy{walkableEnemy: &walkableEnemy{Position: position, isChasing: false}, isVisible: true}
	we.walkFn = we.ghostWalk
	return we
}

func NewOrgeWalkableEnemy(position *value_objects.Position) WalkableEnemyBehaviour {
	we := &walkableEnemy{Position: position, isChasing: false}
	we.walkFn = we.orgeWalk
	return we
}

func NewSnakeWalkableEnemy(position *value_objects.Position) WalkableEnemyBehaviour {
	we := &snakeWalkableEnemy{walkableEnemy: &walkableEnemy{Position: position, isChasing: false}, dirX: 1, dirY: 1}
	we.walkFn = we.snakeWalk
	return we
}

func (we *walkableEnemy) IsChasing() bool {
	return we.isChasing
}

func (we *walkableEnemy) SetChasing(chasing bool) {
	we.isChasing = chasing
}

func (we *walkableEnemy) GetPosition() *value_objects.Position {
	return we.Position
}

func (we *walkableEnemy) SetPosition(pos *value_objects.Position) {
	we.Position = pos
}

func (we *walkableEnemy) IsVisible() bool {
	return true
}

func (we *walkableEnemy) CanAttack(position *value_objects.Position) bool {
	for x := we.Position.X - 1; x <= we.Position.X+1; x++ {
		for y := we.Position.Y - 1; y <= we.Position.Y+1; y++ {
			if position.IsEqual(value_objects.NewPosition(x, y)) {
				return true
			}
		}
	}
	return false
}
func (we *walkableEnemy) Walk(topLeft, botRight *value_objects.Position, busyPos []*value_objects.Position) {
	we.walkFn(topLeft, botRight, busyPos)
}

func (we *walkableEnemy) defaultWalk(topLeft, botRight *value_objects.Position, busyPos []*value_objects.Position) {
	isHorizontalPos := true
	firstPos := we.Position.X
	secondPos := we.Position.Y
	startPos := topLeft.X
	endPos := botRight.X
	if rand.Intn(2) == 0 {
		isHorizontalPos = false
		firstPos = we.Position.Y
		secondPos = we.Position.X
		startPos = botRight.Y
		endPos = topLeft.Y
	}

	isBusy := true
	var nextPos *value_objects.Position
	tryCount := 0

	for isBusy && tryCount < 3 {
		isBusy = false
		if startPos == firstPos {
			firstPos++
		} else if endPos == firstPos {
			firstPos--
		} else {
			if rand.Intn(2) == 0 {
				firstPos++
			} else {
				firstPos--
			}
		}
		if isHorizontalPos {
			nextPos = value_objects.NewPosition(firstPos, secondPos)
		} else {
			nextPos = value_objects.NewPosition(secondPos, firstPos)
		}

		for _, p := range busyPos {
			if nextPos.IsEqual(p) {
				isBusy = true
			}
		}
		tryCount++
	}

	if !isBusy {
		we.Position = nextPos
	}
}

func (we *walkableEnemy) orgeWalk(topLeft, botRight *value_objects.Position, busyPos []*value_objects.Position) {
	isHorizontalPos := true
	firstPos := we.Position.X
	secondPos := we.Position.Y
	if rand.Intn(2) == 0 {
		isHorizontalPos = false
		firstPos = we.Position.Y
		secondPos = we.Position.X
	}

	isBusy := true
	var nextPos *value_objects.Position
	tryCount := 0

	for isBusy && tryCount < 3 {
		isBusy = false
		if rand.Intn(2) == 0 {
			firstPos += 2
		} else {
			firstPos -= 2
		}
		if isHorizontalPos {
			nextPos = value_objects.NewPosition(firstPos, secondPos)
		} else {
			nextPos = value_objects.NewPosition(secondPos, firstPos)
		}

		if !(topLeft.X <= nextPos.X &&
			nextPos.X <= botRight.X &&
			botRight.Y <= nextPos.Y &&
			nextPos.Y <= topLeft.Y) {
			isBusy = true
		} else {
			for _, p := range busyPos {
				if nextPos.IsEqual(p) {
					isBusy = true
				}
			}
		}

		tryCount++
	}

	if !isBusy {
		we.Position = nextPos
	}
}

type ghostWalkableEnemy struct {
	*walkableEnemy
	isVisible bool
}

func (we *ghostWalkableEnemy) ghostWalk(topLeft, botRight *value_objects.Position, busyPos []*value_objects.Position) {
	isBusy := true
	var nextPos *value_objects.Position
	tryCount := 0

	for isBusy && tryCount < 3 {
		isBusy = false
		targetX := rand.Intn(botRight.X-topLeft.X) + topLeft.X
		targetY := rand.Intn(topLeft.Y-botRight.Y) + botRight.Y
		nextPos = value_objects.NewPosition(targetX, targetY)

		for _, p := range busyPos {
			if nextPos.IsEqual(p) {
				isBusy = true
			}
		}
		tryCount++
	}

	if !isBusy {
		we.Position = nextPos
		if rand.Intn(4) == 0 {
			we.isVisible = false
		} else {
			we.isVisible = true
		}
	}
}

func (we *ghostWalkableEnemy) IsVisible() bool {
	return we.isVisible
}

func (we *ghostWalkableEnemy) SetChasing(chasing bool) {
	we.isChasing = chasing
	if we.isChasing {
		we.isVisible = chasing
	}
}

type snakeWalkableEnemy struct {
	*walkableEnemy
	dirX, dirY int
}

func (we *snakeWalkableEnemy) snakeWalk(topLeft, botRight *value_objects.Position, busyPos []*value_objects.Position) {
	isBusy := true
	var nextPos *value_objects.Position
	tryCount := 0

	for isBusy && tryCount < 6 {
		isBusy = false

		newDirX := 1
		if rand.Intn(2) == 0 {
			newDirX = -1
		}

		newDirY := 1
		if rand.Intn(2) == 0 {
			newDirY = -1
		}
		nextPos = value_objects.NewPosition(we.Position.X+newDirX, we.Position.Y+newDirY)

		if we.dirX == newDirX && we.dirY == newDirY {
			isBusy = true
		} else {
			if !(topLeft.X <= nextPos.X &&
				nextPos.X <= botRight.X &&
				botRight.Y <= nextPos.Y &&
				nextPos.Y <= topLeft.Y) {
				isBusy = true
			} else {
				for _, p := range busyPos {
					if nextPos.IsEqual(p) {
						isBusy = true
					}
				}
			}
		}

		tryCount++
	}

	if !isBusy {
		we.Position = nextPos
	}
}
