package main

import eb "github.com/hajimehoshi/ebiten"

type coord struct {
	x, y float64
}

// Une objet simple, qui a une image associée, des coordonnées, une vitesse (signée),
//  et qui peut entrer en collision avec d'autres objets
type object struct {
	image *eb.Image
	coord
	width, height float64
	speedX        float64
	speedY        float64
	dx            float64
	dy            float64

	isColliding bool
}

// retourne les coordonées du coin supérieur gauche et inférieur droit de l'image
func (so *object) bounds() (coord, coord) {
	return coord{so.x, so.y}, coord{so.x + so.width, so.y + so.height}
}

// met à jour le déplacement prévu en fonction de la résolution de collision
func (so *object) updateDxDy(dx float64, dy float64) {
	so.dx -= getSign(so.dx) * dx
	so.dy -= getSign(so.dy) * dy
}

// met à jour la position de l'objet avec le déplacement sur les deux axes
//  doit être appelé après la résolution de collisions
func (so *object) updatePos() {
	so.x += so.dx
	so.y += so.dy
}

type resolution struct {
	obj          *object
	collide      bool
	collideWorld bool
	isVisible    bool
	collideWith  *object
	x            float64 // dernières coordonnées avant collision
	y            float64
	dx           float64 // de combien il faut diminuer/augmenter le déplacement pour arrêter les objets au point de collision
	dy           float64
	otherDx      float64
	otherDy      float64
}

func abs(n float64) float64 {
	if n >= 0 {
		return n
	}
	return -n
}

func newResolution(one *object) resolution {
	res := resolution{}
	res.obj = one
	res.collide = false
	res.collideWorld = false
	res.isVisible = true
	res.collideWith = nil
	res.x = 0
	res.y = 0
	res.dx = 0
	res.dy = 0

	return res
}

/*detection collision avec murs
retourne resolution
*/
func isExiting(obj *object) resolution {
	topLeft, bottomRight := obj.bounds()
	topLeft.x += obj.dx
	topLeft.y += obj.dy
	bottomRight.x += obj.dx
	bottomRight.y += obj.dy

	res := newResolution(obj)
	// res.x = topLeft.x
	// res.y = topLeft.y

	//limite gauche
	if topLeft.x <= 0 {
		res.collide = true
		res.collideWorld = true
		res.dx = abs(topLeft.x) + 1
		if bottomRight.x <= 0 {
			res.isVisible = false
		}
	}

	//limite haut
	if topLeft.y <= 0 {
		res.collide = true
		res.collideWorld = true
		res.dy = abs(topLeft.y) + 1
		if bottomRight.y <= 0 {
			res.isVisible = false
		}
	}

	//limite droite
	if bottomRight.x >= screenWidth {
		res.collide = true
		res.collideWorld = true
		res.dx = bottomRight.x - screenWidth - 1
		if topLeft.x >= screenWidth {
			res.isVisible = false
		}
	}

	//limite bas
	if bottomRight.y >= screenHeight {
		res.collide = true
		res.collideWorld = true
		res.dy = bottomRight.y - screenHeight - 1
		if topLeft.y >= screenHeight {
			res.isVisible = false
		}
	}
	return res
}

// Détection collision entre deux objets
func checkCollision(one *object, two *object) resolution {
	res := newResolution(one)

	oneTopLeft, oneBottomRight := one.bounds()
	oneTopLeft.x += one.dx
	oneTopLeft.y += one.dy
	oneBottomRight.x += one.dx
	oneBottomRight.y += one.dy

	twoTopLeft, twoBottomRight := two.bounds()
	twoTopLeft.x += two.dx
	twoTopLeft.y += two.dy
	twoBottomRight.x += two.dx
	twoBottomRight.y += two.dy

	//collision axe des x
	axeX := oneBottomRight.x >= twoTopLeft.x && twoBottomRight.x >= oneTopLeft.x
	//collision axe des y
	axeY := oneBottomRight.y >= twoTopLeft.y && twoBottomRight.y >= oneTopLeft.y

	//collistion 2 axes
	if axeX && axeY {
		res.collide = true
		res.collideWith = two

		// si oneTopLeft < twoTopLeft (alors one à gauche de two)
		var dx float64
		if oneTopLeft.x < twoTopLeft.x {
			dx = oneBottomRight.x - twoTopLeft.x
		} else {
			dx = twoBottomRight.x - oneTopLeft.x
		}

		// pondérer rectif dx selon déplacement effectif des 2 objets
		oneProp := one.dx / (one.dx + two.dx)
		twoProp := two.dx / (one.dx + two.dx)

		res.dx = dx * oneProp
		res.dy = one.dy
		res.otherDx = dx * twoProp
		res.otherDy = two.dy
	}

	return res
}
