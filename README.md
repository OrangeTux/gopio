# GoPio
A Go package to interact with the GPIO pins on the [Aria G25][aria].

## Usage

    import "github.com/orangetux/gopio"

    pin = gopio.Pin("N22")
    // State is either 0 or 1.
    state := pin.Read()
    pin.Write(0)

[aria]:http://www.acmesystems.it/aria
