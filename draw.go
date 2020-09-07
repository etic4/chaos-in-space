package main

import (
	"image"
	"image/color"

	eb "github.com/hajimehoshi/ebiten"
	ebu "github.com/hajimehoshi/ebiten/ebitenutil"
)

func drawBackground(screen *eb.Image) {
	op := &eb.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(backgroundImg, op)
}

func drawPlayers(screen *eb.Image) {
	for _, player := range players {
		if !player.destroyed {
			playerOp := &eb.DrawImageOptions{}
			playerOp.GeoM.Translate(player.x, player.y)
			screen.DrawImage(player.image, playerOp)
		}
	}
}

func drawBullets(screen *eb.Image) {
	for _, bullet := range bullets {
		bullOp := &eb.DrawImageOptions{}

		bullOp.GeoM.Scale(0.5, 0.5)
		bullet.width = bullet.width * 0.8
		bullet.height = bullet.height * 0.8

		bullOp.GeoM.Translate(bullet.x, bullet.y)
		screen.DrawImage(bullet.image, bullOp)
		drawRect(screen, bullet.x, bullet.y, bullet.width, bullet.height, color.White)
	}
}

func drawAsteroids(screen *eb.Image) {
	for _, asteroid := range asteroids {
		if asteroid.exploding {
			asteroid.nextExplosionStage()
		}
		op := &eb.DrawImageOptions{}
		op.GeoM.Scale(asteroid.scale, asteroid.scale)
		op.GeoM.Translate(asteroid.x, asteroid.y)

		screen.DrawImage(asteroid.image, op)
		drawRect(screen, asteroid.x, asteroid.y, asteroid.width, asteroid.height, color.White)
	}
}

// pour debugger
func drawRect(screen *eb.Image, x float64, y float64, width float64, height float64, clr color.Color) {
	lines := [4][2]coord{
		{coord{x, y}, coord{x + width, y}},
		{coord{x + width, y}, coord{x + width, y + height}},
		{coord{x + width, y + height}, coord{x, y + height}},
		{coord{x, y + height}, coord{x, y}},
	}

	for _, coords := range lines {
		c1, c2 := coords[0], coords[1]
		ebu.DrawLine(screen, c1.x, c1.y, c2.x, c2.y, color.White)
	}
}

func drawAll(screen *eb.Image) {
	drawBackground(screen)
	drawPlayers(screen)
	drawBullets(screen)
	drawAsteroids(screen)
}

func drawHome(screen *eb.Image) {
	drawBackground(screen)

	// Le titre
	op := &eb.DrawImageOptions{}
	op.GeoM.Scale(0.7, 0.7)
	op.GeoM.Translate(60, 60)
	screen.DrawImage(titreImg, op)

	// Les infos
	op = &eb.DrawImageOptions{}
	op.GeoM.Scale(0.7, 0.7)
	op.GeoM.Translate(80, 360)
	screen.DrawImage(infosImg, op)

	// Les créateurs
	op = &eb.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	w, h := créateurs.Size()
	op.GeoM.Translate(float64(screenWidth-(w/2)-20), float64(screenHeight-(h/2)-20))
	screen.DrawImage(créateurs, op)
}

func drawGameOver(screen *eb.Image) {
	op := &eb.DrawImageOptions{}
	op.GeoM.Translate(60, 20)
	screen.DrawImage(gameoverImg, op)

	w, h := infosImg.Size()
	op.GeoM.Scale(0.7, 0.7)
	op.GeoM.Translate(120, 500)

	screen.DrawImage(infosImg.SubImage(image.Rect(0, h/2, 0+w, h/2+h)).(*eb.Image), op)
}
