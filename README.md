# Boids

Yet another boids simulation — or a try at it at least.

A 3D flocking simulation built with Go and [raylib-go](https://github.com/gen2brain/raylib-go), implementing the three classic steering behaviours:

- **Separation** — steer away from nearby boids, harder as they get closer
- **Alignment** — steer toward the average direction of neighbours
- **Cohesion** — steer toward the centre of mass of neighbours

## Requirements

- Go 1.22+
- raylib-go dependencies (see [raylib-go build instructions](https://github.com/gen2brain/raylib-go#requirements))

## Running

```sh
make run
```

## License

GNU General Public License v3.0 — see [LICENSE.md](LICENSE.md).
