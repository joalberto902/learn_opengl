package main

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.4/glfw"
	"log"
	"runtime"
	"strings"
)

const (
	vertexShaderSource = `
	#version 460 core
	layout (location = 0) in vec3 aPos;

	void main() {
		gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
	}
	` + "\x00"

	fragmentShaderSource = `
	#version 460 core
	out vec4 FragColor;

	void main() {
		FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
	}
	` + "\x00"
)

var (
	vertices = []float32{
		0.5, 1.0, 0.0,
		1.0, 1.0, 0.0,
		1.0, 0.5, 0.0,

		-0.5, 1.0, 0.0,
		-1.0, 1.0, 0.0,
		-1.0, 0.5, 0.0,

		0.5, -1.0, 0.0,
		1.0, -1.0, 0.0,
		1.0, -0.5, 0.0,

		-0.5, -1.0, 0.0,
		-1.0, -1.0, 0.0,
		-1.0, -0.5, 0.0,
	}
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
	window, err := glfw.CreateWindow(800, 600, "LearnOpenGL", nil, nil)
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

	//Generating Vertex Buffer Object
	var vbo uint32
	gl.GenBuffers(1, &vbo)                                                            // Generates Buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)                                               //Binding Buffer to context
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW) //copy user-defined data into the currently bound buffer

	//Generating vertex shader
	var vertexShader uint32
	vertexShader = gl.CreateShader(gl.VERTEX_SHADER)
	//Creating pointer to *uint8 of string
	vertexSourcePtr, free := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, vertexSourcePtr, nil)
	free() //Freeing the c counterpart of the pointer
	gl.CompileShader(vertexShader)

	//Verify if it compiled successfuly
	var success int32
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(vertexShader, gl.INFO_LOG_LENGTH, &logLength)

		infoLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(vertexShader, logLength, nil, gl.Str(infoLog))
		log.Printf("ERROR::SHADER::VERTEX::COMPILATION_FAILED\n%v", infoLog)
	}

	//Generating fragment Shader
	var fragmentShader uint32
	fragmentShader = gl.CreateShader(gl.FRAGMENT_SHADER)
	fragmentShaderPtr, free := gl.Strs(fragmentShaderSource)
	gl.ShaderSource(fragmentShader, 1, fragmentShaderPtr, nil)
	free()

	gl.CompileShader(fragmentShader)

	//See if it is success
	gl.GetShaderiv(fragmentShader, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(fragmentShader, gl.INFO_LOG_LENGTH, &logLength)

		infoLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(fragmentShader, logLength, nil, gl.Str(infoLog))
		log.Printf("ERROR::SHADER::FRAGMENT::COMPILATION_FAILED\n%v", infoLog)
	}

	//Creating Shader Program
	var shaderProgram uint32
	shaderProgram = gl.CreateProgram()
	//Attach shaders
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	//Link program
	gl.LinkProgram(shaderProgram)

	//See if it succeeds
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		infoLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(infoLog))
		log.Printf("ERROR::SHADER::PROGRAM::LINK_FAILED\n%v", infoLog)
	}

	//Cleaning up resources
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	//Setting up vao
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vao)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	//Render loop
	for !window.ShouldClose() {
		processInput(window) //processes inputs for the window

		//Render commands
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))

		window.SwapBuffers() //Swaps the buffers on the screen
		glfw.PollEvents()    // Registers the events that happenend on the window
	}
}

func framebufferSizeCallback(w *glfw.Window, width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func processInput(w *glfw.Window) {
	if w.GetKey(glfw.KeyEscape) == glfw.Press {
		w.SetShouldClose(true)
	}
}
