package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unsafe"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/progrium/macdriver/cocoa"
)

// OpenGL calls based on https://stackoverflow.com/a/22502999/762175

//go:embed triangle.vert
var vertData string

//go:embed triangle.frag
var fragData string

var program, vao uint32
var uPosition, aColor, aPosition int32

var vertexData = [][2][3]float32{
	{{+0.0, +0.5, +0.0}, {0, 1, 1}},
	{{+0.5, -0.5, +0.0}, {1, 0, 1}},
	{{-0.5, -0.5, +0.0}, {1, 1, 0}},
}

func setupGL() error {
	view.Send("openGLContext").Send("makeCurrentContext")
	err := gl.Init()
	if err != nil {
		return err
	}

	var m, n int32
	gl.GetIntegerv(gl.MAJOR_VERSION, &m)
	gl.GetIntegerv(gl.MINOR_VERSION, &n)
	if m != 4 || n != 1 {
		return fmt.Errorf("expected GL version 4.1, got %d.%d", m, n)
	}

	vs, err := compileShader(gl.VERTEX_SHADER, vertData)
	if err != nil {
		return err
	}
	defer gl.DeleteShader(vs)

	fs, err := compileShader(gl.FRAGMENT_SHADER, fragData)
	if err != nil {
		return err
	}
	defer gl.DeleteShader(fs)

	program, err = linkProgram(vs, fs)
	if err != nil {
		return err
	}

	uPosition = gl.GetUniformLocation(program, gl.Str("p\x00"))
	aColor = gl.GetAttribLocation(program, gl.Str("colour\x00"))
	aPosition = gl.GetAttribLocation(program, gl.Str("position\x00"))

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	const vecColorWidth = 3                        // [3]float32
	const f32Size = int(unsafe.Sizeof(float32(0))) // float32
	const vecSize = f32Size * vecColorWidth        // [3]float32
	const vecColorSize = vecSize * 2               // [2]float32

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexData)*vecColorSize, gl.Ptr(vertexData), gl.STATIC_DRAW)

	gl.VertexAttribPointer(uint32(aPosition), vecColorWidth, gl.FLOAT, false, int32(vecColorSize), nil)
	gl.VertexAttribPointer(uint32(aColor), vecColorWidth, gl.FLOAT, false, int32(vecColorSize), gl.PtrOffset(vecSize))

	gl.EnableVertexAttribArray(uint32(aPosition))
	gl.EnableVertexAttribArray(uint32(aColor))

	setGLRenderFunc(view, renderGL)
	return nil
}

func renderGL(view cocoa.NSView) bool {
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.UseProgram(program)
	gl.Uniform2f(uPosition, 0, 0)
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(vertexData)))
	return true
}

func compileShader(kind uint32, source string) (uint32, error) {
	shader := gl.CreateShader(kind)

	nullTerminate(&source)
	cstr, free := gl.Strs(source)
	defer free()

	gl.ShaderSource(shader, 1, cstr, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.TRUE {
		return shader, nil
	}

	var n int32
	gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &n)
	log := strings.Repeat("\x00", int(n+1))
	gl.GetShaderInfoLog(shader, n, nil, gl.Str(log))
	return 0, fmt.Errorf("failed to compile shader: %s", log)
}

func linkProgram(shaders ...uint32) (uint32, error) {
	program := gl.CreateProgram()

	for _, s := range shaders {
		gl.AttachShader(program, s)
	}
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.TRUE {
		return program, nil
	}

	var n int32
	gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &n)
	log := strings.Repeat("\x00", int(n+1))
	gl.GetProgramInfoLog(program, n, nil, gl.Str(log))
	return 0, fmt.Errorf("failed to link program: %s", log)
}

func nullTerminate(s *string) {
	if strings.HasSuffix(*s, "\x00") {
		return
	}

	*s += "\x00"
}
