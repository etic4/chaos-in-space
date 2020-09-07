package main

import (
	"math/rand"

	eb "github.com/hajimehoshi/ebiten"
)

// fait feu s'il y a lieu
func fireBullets() {
	for _, player := range players {
		if !player.destroyed && eb.IsKeyPressed(player.keyFire) {
			player.fireBullet(&bullets)
		}
	}
}

// Détermine pour chaque objets la quantité du prochain déplacement
func nextMovement() {
	// joueur
	for _, player := range players {
		if !player.destroyed {
			player.getDxDy()
		}
	}

	// Pour l'instant c'est pas nécessaire de mettre à jour dx avec speedX et dy avec speedY
	//  parce que si deux objets se rencontrent (seule situation ou dx et dy sont modifiés) ils sont détruits et sortis du jeux.
	// Mais si j'implémente un mouvement non linéaire de vaisseaux ennemis, par exemple, j'aurai besoin d'en tenir compte
	for _, bullet := range bullets {
		bullet.dx = bullet.speedX
		bullet.dy = bullet.speedY
	}

	for _, asteroid := range asteroids {
		asteroid.dx = asteroid.speedX
		asteroid.dy = asteroid.speedY
	}
}

// Enlève objets qui se sont rencontrés au frame précédent
// met gameEnded à true si vaisseau détruit
func removeDestroyeds() {
	bulletsToBeRemoved := []int{}
	asteroidsToBeRemoved := []int{}

	for _, player := range players {
		if player.destroyed {
			player.destroy()
			gameEnded = true
		}
	}

	if gameEnded {
		players = players[:0]
	}

	for i, asteroid := range asteroids {
		if asteroid.destroyed {
			asteroidsToBeRemoved = append(asteroidsToBeRemoved, i)
		}
		coll := isExiting(asteroid.object)
		if coll.collide && !coll.isVisible {
			asteroidsToBeRemoved = append(asteroidsToBeRemoved, i)
		}
	}

	for i, bullet := range bullets {
		if bullet.destroyed {
			bulletsToBeRemoved = append(bulletsToBeRemoved, i)
		}
		coll := isExiting(bullet.object)
		if coll.collide && !coll.isVisible {
			bulletsToBeRemoved = append(bulletsToBeRemoved, i)
		}
	}

	// supprime les objets
	for _, i := range bulletsToBeRemoved {
		bullets = deleteFromBulletsSlice(bullets, i)
	}

	for _, i := range asteroidsToBeRemoved {
		asteroids = deleteFromAsteroidsSlice(asteroids, i)
	}
}

// résoud les collisions, met à jour les quantité de déplacement si nécessaire
func resolvCollisions() {
	// collision avec d'autres objets
	// asteroids et lasers
	for _, asteroid := range asteroids {
		if !asteroid.exploding {
			for _, bullet := range bullets {
				coll := checkCollision(asteroid.object, bullet.object)
				if coll.collide {
					asteroid.updateDxDy(coll.dx, coll.dy)
					asteroid.isColliding = true
					asteroid.setDamages(1)

					if asteroid.damages >= asteroid.maxDamages {
						asteroid.initExplosion()
					}

					bullet.updateDxDy(coll.dx, coll.dy)
					bullet.isColliding = true
					bullet.destroyed = true
				}
			}

			// asteroids et player
			for _, player := range players {
				coll := checkCollision(asteroid.object, player.object)
				if coll.collide {
					asteroid.updateDxDy(coll.dx, coll.dy)
					asteroid.isColliding = true
					player.updateDxDy(coll.dx, coll.dy)
					player.isColliding = true
					player.destroyed = true
				}
			}
		}

	}

	// players et bords de la map
	for _, player := range players {
		coll := isExiting(player.object)
		if coll.collide {
			player.updateDxDy(coll.dx, coll.dy)
		}
	}
}

func updatePositions() {
	for _, player := range players {
		player.updatePos()
	}

	for _, bullet := range bullets {
		bullet.updatePos()
	}

	for _, asteroid := range asteroids {
		asteroid.updatePos()
	}
}

func spawnAsteroid() {
	if len(asteroids) < 4 {
		// La vitesse
		speed := float64(100-(rand.Intn(100)-30)) / 100

		// choix de l'image et donc de la résistance
		choice := rand.Intn(len(asteroidsChoice))
		asteroid := newAsteroid(asteroidsChoice[choice], asteroidMaxDamages[choice])

		// positionnement
		y := rand.Intn(screenHeight - int(asteroid.height))
		asteroid.y = float64(y)

		// direction
		asteroid.speedY = float64(rand.Intn(3)) - 1
		asteroid.speedX = asteroid.speedX * speed
		asteroid.speedY = asteroid.speedY * speed
		asteroids = append(asteroids, asteroid)
	}
}

func initGame() {
	if eb.IsKeyPressed(eb.Key1) || eb.IsKeyPressed(eb.KeyKP1) || eb.IsKeyPressed(eb.Key2) || eb.IsKeyPressed(eb.KeyKP2) {
		gameStarted = true
		playerOne = newPlayer(fuséeImg)
		playerOne.keyUp = eb.KeyUp
		playerOne.keyDown = eb.KeyDown
		playerOne.keyLeft = eb.KeyLeft
		playerOne.keyRight = eb.KeyRight
		playerOne.keyFire = eb.KeySpace
		players = append(players, playerOne)

		if eb.IsKeyPressed(eb.Key2) || eb.IsKeyPressed(eb.KeyKP2) {
			playerTwo = newPlayer(fusée2Img)
			playerTwo.keyUp = eb.KeyT
			playerTwo.keyDown = eb.KeyG
			playerTwo.keyLeft = eb.KeyF
			playerTwo.keyRight = eb.KeyH
			playerTwo.keyFire = eb.KeyA
			players = append(players, playerTwo)

			// changer les clés de joueur 1
			playerOne.keyUp = eb.KeyKP5
			playerOne.keyDown = eb.KeyKP2
			playerOne.keyLeft = eb.KeyKP1
			playerOne.keyRight = eb.KeyKP3
			playerOne.keyFire = eb.KeyUp

			// positionne les vaisseaux
			playerOne.y -= 140
			playerTwo.y += 140

		}

		gameEnded = false
		asteroids = asteroids[:0]
		bullets = bullets[:0]
	}
}
