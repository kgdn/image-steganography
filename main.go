/*
 * Author: Kieran Gordon, Heriot-Watt University
 * Date:   09/06/2022
 *
 * Description: This program is a simple program that encrypts and decrypts text using a PPM image.
 * This program is based on the Image Stenography coursework for F28HS - Hardware-Software Interfaces
 * at Heriot-Watt University, Edinburgh. The program is written in Go, and the program which this program
 * is based on was implemented in C by Kieran Gordon. This code is not available due to the university's
 * plagiarism policy.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"os"
)

/* The RGB values of a pixel. */
type Pixel struct {
	R, G, B int
}

/* An image loaded from a PPM file. */
type PPM struct {
	Width, Height, Max int
	Header             string
	Pixels             []Pixel
}

/* Reads an image from an open PPM file.
 * Returns a new struct PPM, or NULL if the image cannot be read.
 */
func getPPM(file *os.File) PPM {
	// Read the header
	var ppm PPM
	// check if the header is valid
	if _, err := fmt.Fscanf(file, "%s", &ppm.Header); err != nil {
		return ppm
	}
	if ppm.Header != "P3" {
		// Not a PPM file so return NULL
		fmt.Println("Not a PPM file")
	}
	// Read the width, height
	if _, err := fmt.Fscanf(file, "%d %d", &ppm.Width, &ppm.Height); err != nil {
		return ppm
	}
	// Read the max value
	if _, err := fmt.Fscanf(file, "%d", &ppm.Max); err != nil {
		return ppm
	}
	// Read the pixels
	ppm.Pixels = make([]Pixel, ppm.Width*ppm.Height)
	for i := 0; i < ppm.Width*ppm.Height; i++ {
		if _, err := fmt.Fscanf(file, "%d %d %d", &ppm.Pixels[i].R, &ppm.Pixels[i].G, &ppm.Pixels[i].B); err != nil {
			return ppm
		}
	}
	return ppm
}

/* Write img to stdout in PPM format. */
func showPPM(img PPM) {
	fmt.Println(img.Header)
	fmt.Println(img.Width, img.Height)
	fmt.Println(img.Max)
	for _, pixel := range img.Pixels {
		fmt.Printf("%d %d %d\n", pixel.R, pixel.G, pixel.B)
	}
}

/* Opens and reads a PPM file, returning a pointer to a new struct PPM.
 * On error, prints an error message and returns NULL.
 */
func readPPM(filename string) PPM {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	ppm := getPPM(file)
	file.Close()
	return ppm
}

/* Encode the string text into the red channel of image img.
 * Returns a new struct PPM, or NULL on error.
 */
func encode(img PPM, text string) PPM {
	// Create a new PPM with the same dimensions as img
	newPPM := PPM{Width: img.Width, Height: img.Height, Max: img.Max, Header: img.Header}
	newPPM.Pixels = make([]Pixel, img.Width*img.Height)
	// Encode the text into the red channel
	for i := 0; i < len(text); i++ {
		newPPM.Pixels[i].R = int(text[i])
	}
	// Copy the green and blue channels from img
	for i := len(text); i < img.Width*img.Height; i++ {
		newPPM.Pixels[i].G = img.Pixels[i].G
		newPPM.Pixels[i].B = img.Pixels[i].B
	}
	return newPPM
}

/* Extract the string encoded in the red channel of newimg, by comparing it
 * with oldimg. The two images must have the same size.
 * Returns a new C string, or NULL on error.
 */
func decode(newimg PPM, oldimg PPM) string {
	// check if the images are the same size
	if newimg.Width != oldimg.Width || newimg.Height != oldimg.Height {
		fmt.Println("Images are not the same size.")
		return ""
	}
	// create a new string to store the message
	message := ""
	// loop through the pixels and extract the message and stop when the message is complete
	for i := 0; i < newimg.Width*newimg.Height; i++ {
		if newimg.Pixels[i].R != oldimg.Pixels[i].R {
			message += string(rune(newimg.Pixels[i].R))
		} else {
			break
		}
	}
	return message
}

func main() {
	if os.Args[1] == "t" {
		// print the input image
		ppm := readPPM(os.Args[2])
		showPPM(ppm)
	} else if os.Args[1] == "e" {
		// Get user input
		fmt.Fprintln(os.Stderr, "Enter a message:")
		var message string
		fmt.Scanln(&message)
		// Read the old image
		oldimg := readPPM(os.Args[2])
		// Encode the message
		newimg := encode(oldimg, message)
		// Write the new image
		showPPM(newimg)
	} else if os.Args[1] == "d" {
		// Read the old image
		oldimg := readPPM(os.Args[2])
		// Read the new image
		newimg := readPPM(os.Args[3])
		// Decode the message
		message := decode(newimg, oldimg)
		// Print the message
		fmt.Println("Decoded message:", message)
	} else {
		fmt.Println("Usage: ./steg [e|d|t] <input file> ><output file>")
		os.Exit(1)
	}
}
