## Image Stenography in Go

This is a simple implementation of image steganography in Go. *Please note that this is a very simple implementation of image steganography. It is not meant to be a robust implementation, but instead a simple implementation that was made to learn the basics of Go.*

The program is designed to encode text into a PPM image and then decode the text back from the image. The program can be invoked using the following command:
`./steg <mode> <input_image> <output_image>`
The program can be run in three modes: `e(ncode)`, `d(ecode)` and `t(est)`. The `e` mode encodes text into the image, the `d` mode decodes the text from the image and the `t` mode simply prints the image to stdout.

## Building the program

To build this program, clone the repository using: 

`git clone https://github.com/kgdn/image-steganography.git`. 

Then, in the `image-steganography` directory, run the following command:

```go build```
