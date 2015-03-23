# GoPio
A Go package to interact with the GPIO pins on the [Aria G25][aria].

## Usage

    import "github.com/orangetux/gopio"

    pin = gopio.Pin(gopio.N22)
    // Value is either 0 or 1.
    pin.Read()

[aria]:http://www.acmesystems.it/aria
