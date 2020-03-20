package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
)

const screenW int = 320
const screenH int = 240
const screenS float64 = 4
const tileSize int = 16
const imgSpeed int = 1

var roomArray [screenW / tileSize][screenH / tileSize]int

var tileSet *ebiten.Image

var player *ebiten.Image
var playerX float64 = 64
var playerY float64 = 64
var playerSpeed float64 = 2
var playerCount int = 0
var playerImg = 0
var isMoving bool = false

func update(screen *ebiten.Image) error {
	isMoving = false

	if player == nil {
		// Create an 16x16 image
		player, _ = ebiten.NewImage(tileSize, tileSize, ebiten.FilterNearest)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && playerY > 0 {
		playerY -= playerSpeed
		if !isMoving {
			playerCount++
			isMoving = true
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && playerY < float64(screenH-tileSize) {
		playerY += playerSpeed
		if !isMoving {
			playerCount++
			isMoving = true
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && playerX > 0 {
		playerX -= playerSpeed
		if !isMoving {
			playerCount++
			isMoving = true
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && playerX < float64(screenW-tileSize) {
		playerX += playerSpeed
		if !isMoving {
			playerCount++
			isMoving = true
		}

	}

	if playerCount > imgSpeed {
		playerImg++
		if playerImg > 7 {
			playerImg = 0
		}
		playerCount = 0
	}
	if !isMoving {
		playerImg = 0
	}
	player, _ = getWhiteGuyImg(playerImg)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(playerX, playerY)

	screen.Fill(color.NRGBA{227, 49, 87, 1})
	//screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	screen.DrawImage(player, opts)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()), 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %d", int(playerX)/tileSize), 0, 8)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Y: %d", int(playerY)/tileSize), 0, 16)

	return nil
}

func getWhiteGuyImg(imgNr int) (*ebiten.Image, error) {

	switch imgNr {
	case 0:
		return ebiten.NewImageFromImage(tileSet.SubImage(image.Rect(16*23, 80, (16*23)+16, 80+16)).(*ebiten.Image), ebiten.FilterDefault)
	case 1:
		return ebiten.NewImageFromImage(tileSet.SubImage(image.Rect(16*24, 80, (16*24)+16, 80+16)).(*ebiten.Image), ebiten.FilterDefault)
	case 2:
		return ebiten.NewImageFromImage(tileSet.SubImage(image.Rect(16*25, 80, (16*25)+16, 80+16)).(*ebiten.Image), ebiten.FilterDefault)
	case 3:
		return ebiten.NewImageFromImage(tileSet.SubImage(image.Rect(16*26, 80, (16*26)+16, 80+16)).(*ebiten.Image), ebiten.FilterDefault)
	case 4:
		return ebiten.NewImageFromImage(tileSet.SubImage(image.Rect(16*27, 80, (16*27)+16, 80+16)).(*ebiten.Image), ebiten.FilterDefault)
	case 5:
		return ebiten.NewImageFromImage(tileSet.SubImage(image.Rect(16*28, 80, (16*28)+16, 80+16)).(*ebiten.Image), ebiten.FilterDefault)
	case 6:
		return ebiten.NewImageFromImage(tileSet.SubImage(image.Rect(16*29, 80, (16*29)+16, 80+16)).(*ebiten.Image), ebiten.FilterDefault)
	case 7:
		return ebiten.NewImageFromImage(tileSet.SubImage(image.Rect(16*30, 80, (16*30)+16, 80+16)).(*ebiten.Image), ebiten.FilterDefault)
	}

	return nil, nil
}

func init() {
	var err error
	tileSet, _, err = ebitenutil.NewImageFromFile("tiles.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := ebiten.Run(update, screenW, screenH, screenS, "Hello world!"); err != nil {
		panic(err)
	}
}
