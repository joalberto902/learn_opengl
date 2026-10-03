package main

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.4/glfw"
	"log"
	"runtime"
)

func main() {
	//Locking main function to the main os thread
	runtime.LockOSThread()

	//Initiating the glfw packager
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	//Window hints
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	//Creating window
	window, err := glfw.CreateWindow(800, 6000, "LearnOpenGL", nil, nil)
	if err != nil {
		log.Fatalf("Failed to create GLFW Window: %v\n", err)
	}
	window.MakeContextCurrent()

	//Initiating gl package
	if err := gl.Init(); err != nil {
		panic(err)
	}

	//Setting up resizing of window
	window.SetFramebufferSizeCallback(framebufferSizeCallback)

	//Render loop
	for !window.ShouldClose() {
		window.SwapBuffers() //Swaps the buffers on the screen
		glfw.PollEvents()    // Registers the events that happenend on the window
	}
}

func framebufferSizeCallback(w *glfw.Window, width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
