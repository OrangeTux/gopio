# GoPio
A Go package to interact with the GPIO pins on the [Aria G25][aria].

## Usage

```go
    import "github.com/orangetux/gopio"

    pin = gopio.Pin{gopio.N22}
    pin.export()
    // Value is either 0 or 1.
    v := pin.Read()
    pin.Write(v)
```

## Todo
    
- [x] ~~Read and write from GPIO pin.~~
- [ ] [#1][1] Write tests
- [ ] [#3][3] Support callback on edges.
- [ ] [#4][4] Use memory map to interacti with GPIO pins.
- [ ] [#2][2] Write benchmarks

[aria]:http://www.acmesystems.it/aria
[1]:https://github.com/OrangeTux/gopio/issues/1
[2]:https://github.com/OrangeTux/gopio/issues/2
[3]:https://github.com/OrangeTux/gopio/issues/3
[4]:https://github.com/OrangeTux/gopio/issues/4
