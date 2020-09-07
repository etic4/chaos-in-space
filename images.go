package main

import (
	"log"

	eb "github.com/hajimehoshi/ebiten"
	ebu "github.com/hajimehoshi/ebiten/ebitenutil"
)

func loadImages() {

	gameoverImg, _, err = ebu.NewImageFromFile("./assets/game over.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	titreImg, _, err = ebu.NewImageFromFile("./assets/titre.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	infosImg, _, err = ebu.NewImageFromFile("./assets/infos.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	backgroundImg, _, err = ebu.NewImageFromFile("./assets/espace.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	fuséeImg, _, err = ebu.NewImageFromFile("./assets/fusée.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	fusée2Img, _, err = ebu.NewImageFromFile("./assets/fusée2.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	bulletImg, _, err = ebu.NewImageFromFile("./assets/missile.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	asteroidImg, _, err = ebu.NewImageFromFile("./assets/asteroid.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	asteroid90Img, _, err = ebu.NewImageFromFile("./assets/asteroid brun 90 degré.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	asteroidtoutpetitImg, _, err = ebu.NewImageFromFile("./assets/astéroid mini brun.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	asteroidpetitImg, _, err = ebu.NewImageFromFile("./assets/astéroid petit brun.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	asteroidExplosion1, _, err = ebu.NewImageFromFile("./assets/explosion1.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	asteroidExplosion2, _, err = ebu.NewImageFromFile("./assets/explosion2.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	asteroidExplosion3, _, err = ebu.NewImageFromFile("./assets/explosion3.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	asteroidExplosion4, _, err = ebu.NewImageFromFile("./assets/explosion4.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	créateurs, _, err = ebu.NewImageFromFile("./assets/créateurs.png", eb.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}
