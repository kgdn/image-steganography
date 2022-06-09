## Image Steganography in Go

This is a simple implementation of image stenography in Go. The program was originally written in C for F28HS (Hardware-Software Interfaces) at Heriot-Watt University, but has been converted to Go so I could learn the language more in depth. The original C code is not available on GitHub or any other website due to the university's plagiarism policy, which can be read [here](https://www.hw.ac.uk/uk/students/doc/plagiarismguide.pdf). *Please note that this is a very simple implementation of image stenography. It is not meant to be a robust implementation, but it is meant to be a simple implementation that can be used to learn the language.*

The program is designed to encode text into a PPM image and then decode the text back into the image. The program can be invoked using the following command:
`./steg <mode> <input_image> <output_image>`
The program can be run in three modes: `e(ncode)`, `d(ecode)` and `t(est)`. The `e` mode encodes text into the image, the `d` mode decodes the text from the image and the `t` mode simply prints the image to stdout.

## Building the program

To build this program, clone the repository and run the following command:
`go build`
