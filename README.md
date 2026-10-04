# learn_opengl

My journey through [LearnOpenGL](https://learnopengl.com/), with every example written in **Go** instead of C++.

LearnOpenGL is a free online book by Joey de Vries that teaches modern OpenGL step by step. This repository holds my own code as I work through it, translating the C++ examples to Go and recording what I learn along the way. All credit for the course material goes to the original author.

## Why Go?

The book's examples use C++, but the OpenGL API is the same in any language, so the concepts carry over directly. Writing them in Go forces me to understand each call instead of copying it, and it surfaces Go-specific details such as thread handling, C string conversion, and byte sizes when uploading data to the GPU.

## Progress

Following the book's "Getting started" section:

- [x] Hello Window
- [x] Hello Triangle
- [ ] Shaders
- [ ] Textures
- [ ] Transformations
- [ ] Coordinate Systems
- [ ] Camera

## Requirements

- **Go**, using the version declared in `go.mod`
- **A C compiler**, because the GLFW and OpenGL bindings use cgo
- **A GPU and driver with OpenGL support**

Platform setup for cgo:

| OS      | What to install                                                                 |
| ------- | ------------------------------------------------------------------------------- |
| Linux   | `gcc`, plus OpenGL and X11 headers (on Debian/Ubuntu: `libgl1-mesa-dev`, `xorg-dev`) |
| Windows | A GCC toolchain such as MinGW-w64 (for example via MSYS2)                       |
| macOS   | Xcode Command Line Tools (`xcode-select --install`). Note that macOS supports OpenGL only up to version 4.1 |

## Running

```sh
git clone https://github.com/joalberto902/learn_opengl.git
cd learn_opengl
go mod tidy
go run .
```

Press **Esc** to close the window.

## Project structure

```
.
├── main.go     # current chapter's example
├── go.mod      # module definition and dependencies
├── go.sum      # dependency checksums
└── .gitignore
```

## Dependencies

- [go-gl/glfw](https://github.com/go-gl/glfw) for window creation, input, and events
- [go-gl/gl](https://github.com/go-gl/gl) for OpenGL bindings

## Notes on translating C++ to Go

Things that differ from the book and are worth remembering:

- The main goroutine must be locked to its OS thread, because an OpenGL context only works on the thread where it was made current.
- The `gl` package must be initialized with `gl.Init()` after the context is made current, and before any other `gl.*` call.
- Buffer sizes are in **bytes**, so a slice of `float32` needs its length multiplied by 4.
- Shader sources must be null-terminated and converted with `gl.Strs`, and the C copy has to be freed afterwards.
- OpenGL errors are silent by default. Checking `gl.GetError()` and shader compile/link logs is essential when something renders as a black screen.

## Resources

- [LearnOpenGL](https://learnopengl.com/), the book this repository follows
- [go-gl documentation](https://github.com/go-gl)
- [GLFW documentation](https://www.glfw.org/documentation.html)

## License

No license has been chosen yet. Add a `LICENSE` file to state how others may use this code.
