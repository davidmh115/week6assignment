# Overview

This project implements an image processing pipeline in Go, demonstrating both concurrent and sequential processing modes. It uses goroutines and channels.

This idea and structure are based on [Amrit Singh’s goroutines pipeline project](https://github.com/code-heim/go_21_goroutines_pipeline)

There are several enhancements and modifications for benchmarking and testing I have added. 

The pipeline performs the following steps:
- Load input images from a directory
- Resize each image to a fixed resolution
- Convert the image to grayscale
- Save the processed image to an output folder

Go 1.20 or later is required

# How To Run

1. Place your input images in the `images/` folder. Supported formats are `.jpeg`, `.jpg`, and `.png`.

2. Run the program with:

bash 
go run main.go

3. Run with goroutines enabled (concurrent processing):

bash
go run main.go -goroutines=true

4. Run with goroutines disabled (sequential processing)

bash 
run main.go -goroutines=false

# Testing and Benchmarks 

To run unit tests: 

bash
go test 

To run benchmarks and see performance metrics:

bash
go test -bench=.

# Error Handling
Error checking was added throughout the pipeline to handle missing or unreadable image files, unsupported file formats, and failures in image decoding or saving. 

The program logs appropriate error messages and safely continues processing other images instead of crashing. This makes the pipeline more reliable and production friendly.

# Performance Results

go run main.go:
Success!
Success!
Success!
Success!
Pipeline completed in 148.9307ms

go run main.go -goroutines=true:
Success!
Success!
Success!
Success!
Pipeline completed in 200.8188ms

go run main.go -goroutines=false:
Success!
Success!
Success!
Success!
Pipeline completed in 201.0631ms

The pipeline was executed successfully for concurrent and sequential modes.The default run (no flag) performed slightly faster than these. The concurrent and sequential modes had comparable runtimes for this small batch of images.

For small workloads, goroutines may not offer a noticeable performance gain due to the overhead of concurrency.

However, for larger batches or heavier processing, goroutines would likely show greater benefits so the next step would be to test with a larger amount of images. 

# Test Results 
PASS
ok      goroutines_pipeline/image_processing    0.481s

# Benchmark Results 
ok      goroutines_pipeline/image_processing    0.571s

# GenAI Tools

I used Copilot to generate my own Star Wars themed images to replace the old ones. ChatGPT helped me determine an efficient way to compare processing times by adding the following code to main():

// Start measuring CPU time
	start := time.Now()

// End measuring CPU time
elapsed := time.Since(start)
	fmt.Printf("Pipeline completed in %s\n", elapsed)