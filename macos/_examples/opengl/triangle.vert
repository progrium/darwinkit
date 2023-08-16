#version 150

uniform vec2 p;
in vec4 position;
in vec4 colour;
out vec4 colourV;

void main (void)
{
    colourV = colour;
    gl_Position = vec4(p, 0.0, 0.0) + position;
}