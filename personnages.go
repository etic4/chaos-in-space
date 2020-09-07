package main

import (
	"time"

	"github.com/hajimehoshi/ebiten"
	eb "github.com/hajimehoshi/ebiten"
)

// enum modèles de munitions
const (
	BASIC int = iota
	BLASER
	BMISSILE
)

type player struct {
	*object
	prevBulletTime time.Time
	keyUp          eb.Key
	keyDown        eb.Key
	keyLeft        eb.Key
	keyRight       eb.Key
	keyFire        eb.Key
	destroyed      bool
}

// Détermine quelle clé de direction a été pressée et les déplacements résultants sur les deux axes
func (p *player) getDxDy() {
	p.dx = 0
	p.dy = 0

	if eb.IsKeyPressed(p.keyUp) { //up
		p.dy = -p.speedY
	}
	if eb.IsKeyPressed(p.keyDown) { //down
		p.dy = p.speedY
	}
	if eb.IsKeyPressed(p.keyLeft) {
		p.dx = -p.speedX
	}
	if eb.IsKeyPressed(p.keyRight) {
		p.dx = +p.speedX
	}
}

func (p *player) fireBullet(bullets *[]*bullet) {
	now := time.Now()
	if now.Sub(p.prevBulletTime) > bulletsTime*time.Millisecond {
		bullet := newBullet(p)
		p.prevBulletTime = now
		*bullets = append(*bullets, bullet)
	}
}

func (p *player) destroy() {
	p.destroyed = true
	p.x = 0
	p.y = 0
	p.width = 0
	p.height = 0
}

// Crée un nouveau joueur
func newPlayer(image *eb.Image) *player {
	player := &player{}
	player.object = &object{} // nécessaire sinon 'nil pointer dereference' sur player.image qui est un pointeur

	player.image = image

	w, h := image.Size()
	player.width = float64(w)
	player.height = float64(h)

	player.x = 60
	player.y = screenHeight/2.0 - float64(h)/2

	player.speedX = defaultShipSpeed
	player.speedY = defaultShipSpeed

	player.dx = 0
	player.dy = 0

	player.prevBulletTime = time.Now()

	return player
}

type bullet struct {
	*object
	model     int // modèle de balle = enum
	destroyed bool
}

// Crée une nouvelle balle
func newBullet(player *player) *bullet {
	bull := &bullet{}
	bull.object = &object{}

	bull.image = bulletImg
	w, h := bulletImg.Size()
	bull.width = float64(w)
	bull.height = float64(h)

	bull.x = player.x + player.width
	bull.y = player.y + player.height/2 - float64(h)/2

	bull.speedX = defaultBulletSpeed
	bull.speedY = 0
	bull.dx = bull.speedX
	bull.dy = bull.speedY

	return bull
}

type asteroid struct {
	*object
	damages         int
	maxDamages      int
	destroyed       bool
	scale           float64
	exploding       bool
	explosionStage  int
	explosionStages []*eb.Image
}

// encaisse les dégâts càd diminue la vie et réduit la taille de l'image (et de la hit box)
func (ast *asteroid) setDamages(dam int) {
	ast.damages += dam
	ast.scale *= .9

	newW := ast.width * ast.scale
	newH := ast.height * ast.scale

	ast.x += (ast.width - newW) / 2
	ast.y += (ast.height - newH) / 2

	ast.width = newW
	ast.height = newH
}

// initie l'animation d'explosion
func (ast *asteroid) initExplosion() {
	ast.exploding = true
	ast.explosionStages = []*eb.Image{asteroidExplosion1, asteroidExplosion2, asteroidExplosion3, asteroidExplosion4}
	ast.explosionStage = 0

	// reset de l'échelle de l'image
	ast.scale /= ast.scale
}

// passe à l'étape suivante de l'explosion
// met destroyed à true une fois l'animation terminée
func (ast *asteroid) nextExplosionStage() {
	ast.explosionStage++

	nextStage := ast.explosionStage / (int(ebiten.CurrentTPS()/2) / len(ast.explosionStages))

	if nextStage < len(ast.explosionStages) {
		// ancien centre
		cX, cY := ast.getImgCEnter()

		// changement d'image
		ast.setNewImg(ast.explosionStages[nextStage])

		// centrage de la nouvelle image sur l'ancien centre
		ast.centerImgOn(cX, cY)

	} else {
		ast.exploding = false
		ast.destroyed = true
	}
}

// change l'image et les dimensions
func (ast *asteroid) setNewImg(image *eb.Image) {
	ast.image = image
	w, h := ast.image.Size()
	ast.width = float64(w)
	ast.height = float64(h)
}

// retourne le centre de l'image
func (ast *asteroid) getImgCEnter() (x float64, y float64) {
	return ast.x + ast.width/2, ast.y + ast.height/2
}

// modifier les coordonées de l'image de sorte à la centrer sur les coordonées passées
func (ast *asteroid) centerImgOn(x float64, y float64) {
	ast.x = x - ast.width/2
	ast.y = y - ast.height/2
}

// Nouvel astéroïde
func newAsteroid(image *eb.Image, maxDamages int) *asteroid {
	ast := &asteroid{}
	ast.object = &object{}

	ast.image = image
	w, h := ast.image.Size()
	ast.width = float64(w)
	ast.height = float64(h)

	ast.x = float64(screenWidth) - 1
	ast.y = float64(screenHeight) / 2

	ast.speedX = defaultAsteroidSpeed
	ast.speedY = 0
	ast.dx = ast.speedX
	ast.dy = ast.speedY

	ast.maxDamages = maxDamages
	ast.scale = float64(1)

	return ast
}
