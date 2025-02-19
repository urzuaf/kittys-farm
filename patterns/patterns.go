package patterns

import (
	"math"
)

func GetPattern(choice int, x, y, speedX, speedY, screenWidth, ScreenHeight float64, index, param, tick int) (float64, float64) {

	//0,1,2,4,5,6,7,8,9 works good
	//3, works wrong

	switch choice {
	case 0:
		x, y = Lineal(x, y, speedY, speedX)
	case 1:
		x, y = CurveSin(x, y, speedY, speedX, param)
	case 2:
		x, y = ZigZag(x, y, speedY, speedX, param)
	case 3:
		x, y = LMovement(x, y, speedY, speedX, index, tick, param)
	case 4:
		x, y = Spiral(x, y, speedY, speedX, param)
	case 5:
		x, y = Waves(x, y, speedY, speedX, index, param)
	case 6:
		x, y = VMovement(x, y, speedY, speedX, index, param)
	case 7:
		x, y = ProgresiveAccelaration(x, y, speedY, speedX, param)
	case 8:
		x, y = FalseRetreat(x, y, speedY, speedX, tick)
	case 9:
		x, y = Circle(x, y, speedY, speedX, screenWidth, param)
	default:
		x, y = Lineal(x, y, speedY, speedX)
	}

	if x > screenWidth-32 {
		x = screenWidth - 32
	}
	if x < 0 {
		x = 0
	}

	return x, y
}

func Lineal(x, y, speedY, speedX float64) (float64, float64) {
	y += speedY
	return x, y
}

func CurveSin(x, y, speedY, speedX float64, param int) (float64, float64) {
	y += speedY
	x += speedX * math.Sin(float64(param)*0.1)
	return x, y
}

func ZigZag(x, y, speedY, speedX float64, param int) (float64, float64) {
	direction := 1
	if param/30%2 == 0 {
		direction = -1
	}
	y += speedY
	x += float64(direction) * speedX
	return x, y
}

func LMovement(x, y, speedY, speedX float64, index, tick, param int) (float64, float64) {
	//10 frames every 2 seconds
	if tick%80 < 40 {
		x += speedX
		y += 0
	} else {
		x -= speedX
		y += speedY
	}
	return x, y
}

func Spiral(x, y, speedY, speedX float64, param int) (float64, float64) {
	angle := float64(param) * 0.1
	radius := 4
	y += speedY
	x += float64(radius) * math.Cos(angle)
	return x, y
}

func Waves(x, y, speedY, speedX float64, index, param int) (float64, float64) {
	waveAmplitude := 1.0
	waveSpeed := 0.05
	y += speedY
	x += waveAmplitude * math.Sin(float64(param)*waveSpeed+float64(index)*0.5)
	return x, y
}

func VMovement(x float64, y, speedY, speedX float64, index, param int) (float64, float64) {
	y += speedY
	if index%2 == 0 {
		x += speedX * 0.8
	} else {
		x -= speedX * 0.8
	}
	return x, y
}

func ProgresiveAccelaration(x, y, speedY, speedX float64, param int) (float64, float64) {
	speedFactor := float64(param) * 0.0005
	y += speedY * (1 + speedFactor)

	return x, y
}

func FalseRetreat(x float64, y, speedY, speedX float64, tick int) (float64, float64) {

	if tick > 20 && tick < 60 && y > 10 {
		y -= 0.4
	} else {
		y += speedY * 1.2
	}

	return x, y
}

func Circle(x, y, speedY, speedX, screenWidth float64, param int) (float64, float64) {
	//this ++ doesn't do anything but it makes the linter happy
	x++
	centerX := screenWidth/2 - 56
	radius := screenWidth/2 + 56
	angle := float64(param) * 0.05

	y += speedY
	x = centerX + radius*math.Cos(angle)

	return x, y

}
