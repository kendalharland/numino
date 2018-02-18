package numino

import (
	"github.com/faiface/pixel/pixelgl"
)

// ViewFunc runs an game view, such as the main menu or gameplay screen.
//
// A view is expected to signal when it is done by emitting a GoToCmd on the
// given done channel.
type ViewFunc func(win *pixelgl.Window, done chan GoToCmd)

// GoToCmd instructs the numino application to navigate to another view.
// These values are emitted when a Viewer is done controlling the application.
type GoToCmd int

const (
	// GoToExit instructs numino to exit.
	GoToExit GoToCmd = iota
	// GoToNewGame instructs numino to start a new game.
	GoToNewGame
	// GoToMenu instructs numino to go to the main menu.
	GoToMenu
	// GoToControls instructs numino to show the controls.
	GoToControls
)
