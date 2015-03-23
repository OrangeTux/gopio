# GoPio
A Go package to interact with the GPIO pins on the [Aria G25][aria].

## Usage

    import "github.com/orangetux/gopio"

    pin = gopio.Pin{gopio.N22}
    pin.export()
    // Value is either 0 or 1.
    v := pin.Read()
    pin.Write(v)

## Todo
    
    - [x] ~~Read and write.~~
    - [ ] #1 Write tests
    - [ ] #3 Support callback on edges.
    - [ ] #4 Use memory map to interacti with GPIO pins.
    - [ ] #2 Write benchmarks

[aria]:http://www.acmesystems.it/aria
