package robot

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgb/xtest"
	"image"
	"image/color"
	"time"
)

type Robot interface {
	// Creates a NewRobot connection
	NewRobot() Robot

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

type XRobot struct {
	X      *xgb.Conn
	Screen *xproto.ScreenInfo
}

func (robot *XRobot) NewRobot() {
	X, err := xgb.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	setup := xproto.Setup(X)

	robot.Screen = setup.DefaultScreen(X)
	robot.X = X
}

func (robot XRobot) String() string {
	return "XRobot Session"
}

func (robot XRobot) MouseMove(x, y int16) {
	root := robot.Screen.Root
	none := xproto.Window(0)
	cookie := xproto.WarpPointerChecked(robot.X, none, root, 0, 0, 0, 0, x, y)
	err := cookie.Check()
	if err != nil {
		fmt.Printf("Error:%v\n", err)
	} else {
		fmt.Printf("Unchecked\n")
	}

	time.Sleep(2000)
}

func (robot XRobot) KeyPress(keycode int) {
	if err := xtest.Init(robot.X); err != nil {
		fmt.Printf("Error:%v\n", err)
		return
	}
	typ := byte(xproto.KeyPress)
	detail := byte(keycode)
	time := uint32(0)
	id := byte(0)
	cookie := xtest.FakeInputChecked(robot.X, typ, detail, time,
		robot.Screen.Root, 0, 0, id)
	err := cookie.Check()
	if err != nil {
		fmt.Printf("Error:%v\n", err)
	} else {
		fmt.Printf("Unchecked\n")
	}

}

func (robot XRobot) Version() {
	cookie := xtest.GetVersionUnchecked(robot.X, 2, 7)
	fmt.Printf("Cookie:%v\n", cookie)
	reply, err := cookie.Reply()
	if err != nil {
		fmt.Printf("Error:%v\n", cookie)
	}
	fmt.Printf("reply:%v\n", reply)

}

/*
	// Might come in handy to keep this around
	ev, xerr := robot.X.WaitForEvent()
	if ev == nil && xerr == nil {
		fmt.Println("Both null")
	}
	if ev != nil {
		fmt.Printf("Event:%v\n", ev)
	}
	if xerr != nil {
		fmt.Printf("Event:%v\n", xerr)
*/
