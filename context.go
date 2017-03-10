//+build !js

package main

import "github.com/go-gl/glfw/v3.1/glfw"

type Context struct {
	window *glfw.Window
}

func getContext(width, height int) (*Context, error) {
	err := glfw.Init()
	if err != nil {
		return nil, err
	}
	window, err := glfw.CreateWindow(width, height, "Test", nil, nil)
	if err != nil {
		return nil, err
	}
	return &Context{
		window: window,
	}, nil
}

func (c *Context) Close() error {
	glfw.Terminate()
	return nil
}
