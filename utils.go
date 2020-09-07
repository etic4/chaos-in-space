package main

func deleteFromBulletsSlice(sl []*bullet, i int) []*bullet {
	if i < len(sl)-1 {
		copy(sl[i:], sl[i+1:])
	}
	sl[len(sl)-1] = nil
	sl = sl[:len(sl)-1]

	return sl
}

func deleteFromAsteroidsSlice(sl []*asteroid, i int) []*asteroid {
	if i < len(sl)-1 {
		copy(sl[i:], sl[i+1:])
	}
	sl[len(sl)-1] = nil
	sl = sl[:len(sl)-1]

	return sl
}

func getSign(speed float64) float64 {
	if speed > 0 {
		return 1.0
	}
	return -1.0
}
