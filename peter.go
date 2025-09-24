package peter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math"
)

var (
	img    *ebiten.Image
	turtle *Turtle
)

type Turtle struct {
	X, Y     float64
	Angle    float64
	PenDown  bool
	PenColor color.Color
	Message  string
}

// Initialisation de la fenêtre et du turtle
func Init(width, height int) {
	img = ebiten.NewImage(width, height)
	turtle = &Turtle{
		X:        float64(width) / 2,
		Y:        float64(height) / 2,
		Angle:    0,
		PenDown:  true,
		PenColor: color.White,
	}
}

// Fonctions du turtle
func Down()   { turtle.PenDown = true }
func Up()     { turtle.PenDown = false }
func Right()  { turtle.Angle += math.Pi / 2 }
func Left()   { turtle.Angle -= math.Pi / 2 }
func Pivote(deg float64) { turtle.Angle += deg * math.Pi / 180 }
func Forward(nbPas float64) {
	x2 := turtle.X + nbPas*math.Cos(turtle.Angle)
	y2 := turtle.Y + nbPas*math.Sin(turtle.Angle)
	if turtle.PenDown {
		ebitenutil.DrawLine(img, turtle.X, turtle.Y, x2, y2, turtle.PenColor)
	}
	turtle.X, turtle.Y = x2, y2
}
func Color(c color.Color) { turtle.PenColor = c }
func Say(msg string)      { turtle.Message = msg }

// Lancer la fenêtre Ebiten et exécuter un script
func Run() {
	game := &Game{}
	ebiten.RunGame(game)
}

// Structure pour Ebiten
type Game struct{}
func (g *Game) Update() error { return nil }
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
	if turtle.Message != "" {
		ebitenutil.DebugPrint(screen, turtle.Message)
	}
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return img.Bounds().Dx(), img.Bounds().Dy()
}
