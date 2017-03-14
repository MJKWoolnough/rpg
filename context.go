//+build !js

package main

import (
	"image/color"
	"runtime"

	"github.com/go-gl/gl/v3.1/gles2"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Context struct {
	window *glfw.Window
}

func getContext(width, height int) (*Context, error) {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		return nil, err
	}
	glfw.WindowHint(glfw.ClientAPI, glfw.OpenGLESAPI)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)
	window, err := glfw.CreateWindow(width, height, "Test", nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()
	err = gles2.Init()
	if err != nil {
		return nil, err
	}
	window.SwapBuffers()
	return &Context{
		window: window,
	}, nil
}

func (c *Context) Close() error {
	c.window.Destroy()
	glfw.Terminate()
	return nil
}

func (c *Context) SetColour(cl color.Color) {
	r, g, b, a := cl.RGBA()
	gles2.ClearColor(float32(r)/255, float32(g)/255, float32(b)/255, float32(a)/255)
	gles2.Clear(gles2.COLOR_BUFFER_BIT)
	c.window.SwapBuffers()
}
