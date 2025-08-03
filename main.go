package main

import (
	"flag"
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"strings"
	"time"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

var useGoroutines = flag.Bool("goroutines", true, "Enable pipeline with goroutines")

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		for _, p := range paths {
			img := imageprocessing.ReadImage(p)
			if img == nil {
				fmt.Printf("Skipping image %s due to read error.\n", p)
				continue
			}
			job := Job{
				InputPath: p,
				OutPath:   strings.Replace(p, "images/", "images/output/", 1),
				Image:     img,
			}
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input {
			imageprocessing.WriteImage(job.OutPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

// --------- Sequential version (no goroutines) ----------

func runSequential(imagePaths []string) {
	for _, p := range imagePaths {
		img := imageprocessing.ReadImage(p)
		if img == nil {
			fmt.Printf("Skipping image %s due to read error.\n", p)
			continue
		}
		img = imageprocessing.Resize(img)
		img = imageprocessing.Grayscale(img)
		outPath := strings.Replace(p, "images/", "images/output/", 1)
		imageprocessing.WriteImage(outPath, img)
		fmt.Println("Success!")
	}
}

// -------------------- MAIN -----------------------------

func main() {
	// Start measuring CPU time
	start := time.Now()
	flag.Parse()

	imagePaths := []string{
		"images/myimage1.jpeg",
		"images/myimage2.jpeg",
		"images/myimage3.jpeg",
		"images/myimage4.jpeg",
	}

	if *useGoroutines {
		// Parallel pipeline version
		channel1 := loadImage(imagePaths)
		channel2 := resize(channel1)
		channel3 := convertToGrayscale(channel2)
		writeResults := saveImage(channel3)

		for success := range writeResults {
			if success {
				fmt.Println("Success!")
			} else {
				fmt.Println("Failed!")
			}
		}
	} else {
		// Sequential version
		runSequential(imagePaths)
	}
	// End measuring CPU time
	elapsed := time.Since(start)
	fmt.Printf("Pipeline completed in %s\n", elapsed)
}
