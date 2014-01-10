package robot

import (
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
	"image"
	"image/color"
)

type Robo interface {
	// Creates an image containing pixels read from the screen.
	CreateScreenCapture(screenRect image.Rectangle) image.Image

	// Sleeps for the specified time.
	Delay(ms int)

	// Returns the number of milliseconds this Robot sleeps after
	// generating an event.
	GetAutoDelay() int

	// Returns the color of a pixel at the given screen coordinates.
	GetPixelColor(x, y int) color.Color

	// Returns whether this Robot automatically invokes waitForIdle after
	// generating an event.
	IsAutoWaitForIdle() bool

	// Presses a given key.
	KeyPress(keycode int)

	// Releases a given key.
	KeyRelease(keycode int)

	// Moves mouse pointer to given screen coordinates.
	MouseMove(x, y int)

	// Presses one or more mouse buttons.
	MousePress(buttons int)

	// Releases one or more mouse buttons.
	MouseRelease(buttons int)

	// Rotates the scroll wheel on wheel-equipped mice.
	MouseWheel(wheelAmt int)

	// Sets the number of milliseconds this Robot sleeps
	// after generating an event.
	SetAutoDelay(ms int)

	// Sets whether this Robot automatically invokes waitForIdle
	// after generating an event.
	SetAutoWaitForIdle(isOn bool)

	// Returns a string representation of this Robot.
	String() string

	// Waits until all events currently on the event queue have been processed.
	WaitForIdle()
}

type Robot struct {
}

func (Robot) Window() {
	wde.Run()
	//wde.NewWindow(300, 100)
}
