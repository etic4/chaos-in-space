package main

import (
	"log"

	eb "github.com/hajimehoshi/ebiten"
)

// Constantes
const (
	screenWidth, screenHeight = 960, 720
	defaultShipSpeed          = 4.0 // pixels (?)
	defaultBulletSpeed        = 8.0
	defaultAsteroidSpeed      = -2.0
	bulletsTime               = 200 // millisecondes
)

// Variables globales
var (
	err                  error
	gameStarted          bool
	gameEnded            bool
	titreImg             *eb.Image
	infosImg             *eb.Image
	gameoverImg          *eb.Image
	backgroundImg        *eb.Image
	fuséeImg             *eb.Image
	fusée2Img            *eb.Image
	bulletImg            *eb.Image
	asteroidImg          *eb.Image
	asteroid90Img        *eb.Image
	asteroidpetitImg     *eb.Image
	asteroidtoutpetitImg *eb.Image
	asteroidExplosion1   *eb.Image
	asteroidExplosion2   *eb.Image
	asteroidExplosion3   *eb.Image
	asteroidExplosion4   *eb.Image
	créateurs            *eb.Image
	asteroidsChoice      []*eb.Image
	asteroidMaxDamages   []int
	bullets              []*bullet
	asteroids            []*asteroid
	playerOne            *player
	playerTwo            *player
	players              []*player
)

// ebiten execute cette fonction une seule fois au lancement
func init() {
	loadImages()

	// images (et donc tailles) possibles pour les asteroïdes
	asteroidsChoice = []*eb.Image{asteroidImg, asteroid90Img, asteroidpetitImg, asteroidtoutpetitImg}
	// nombre de tirs qu'un asteroïde peut supporter, (// asteroidCoice)
	asteroidMaxDamages = []int{3, 3, 2, 1}

	gameStarted = false
	gameEnded = false
}

func update(screen *eb.Image) error {
	if !gameStarted {
		drawHome(screen)
		initGame()
		return nil
	}

	// tire les missiles
	fireBullets()

	// envoie des asteroides
	spawnAsteroid()

	// détermine le prochain mouvement
	nextMovement()

	//enlève objets détruits ou sortis de la map
	removeDestroyeds()

	// détermine les collisions et les modifications de mouvement qu'elle entraînent
	resolvCollisions()

	// mets à jour le positions en fonction de la résolution des collisions
	updatePositions()

	// TODO: mettre mise à jour des explosions ici plutôt que dans drawAsteroids

	if eb.IsDrawingSkipped() {
		return nil
	}

	drawAll(screen)

	if gameEnded {
		drawGameOver(screen)
		initGame()
		return nil
	}

	return nil
}

func main() {
	eb.SetMaxTPS(60)
	if err := eb.Run(update, screenWidth, screenHeight, 1, "Chaos dans l'espace"); err != nil {
		log.Fatal(err)
	}
}
