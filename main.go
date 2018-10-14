package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

// Medidas: Tipo de medidas guarda un arreglo de medidas
type Medidas struct {
	Medidas []Medida `json:medidas`
}

// Medida: Tipo de medida guarda lso datos de la medida
type Medida struct {
	ImageType string `json:imageType`
	Width     uint   `json:width`
	Height    uint   `json:height`
}

func main() {

	flag.Parse()
	root := flag.Arg(0)

	medidasFile, err := os.Open("./sizes.json")
	//handlle error
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Cargando medidas...")
	defer medidasFile.Close()

	// abrimos nuestro archivo json
	byteValue, _ := ioutil.ReadAll(medidasFile)

	//inicializamos nuestras medidas
	var medidas Medidas

	//se le aplica el unmarshal al bytearray
	json.Unmarshal(byteValue, &medidas)

	//abriedno directorio
	artes, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range artes {
		fmt.Printf("Abriendo Archivo %s\n", f.Name())
		//iteramos en cada uno de nuestro arregoo de medidas
		// e imprimimos la información que contiene
		//abrimos imagen
		imageReader, err := os.Open(root + f.Name())
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		//defer imageReader.Close()

		im, _, err := image.Decode(imageReader)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		for i := 0; i < len(medidas.Medidas); i++ {
			nombre := strings.Split(f.Name(), ".")

			newImageName := nombre[0] + "_" + medidas.Medidas[i].ImageType + "." + nombre[1]

			fmt.Println("Resizing image to W:", strconv.FormatUint(uint64(medidas.Medidas[i].Width), 10))
			//fmt.Println("Resizing image to H:", strconv.FormatUint(medidas.Medidas[i].Height, 10))
			fmt.Printf("Imagen %s_%s.%s Guardada\n", nombre[0], medidas.Medidas[i].ImageType, nombre[1])
			//fmt.Println("Tipo de Imagen: " + medidas.Medidas[i].ImageType)
			newImage := resize.Resize(medidas.Medidas[i].Width, 0, im, resize.Lanczos3)
			if nombre[1] == "jpg" {
				out, err := os.Create("./newimages/" + newImageName)
				err = jpeg.Encode(out, newImage, nil)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			}

		}
	}

}

/*func main() {
V	medidas, err := os.Open("./sizes.json")
	//handlle error
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Cargando medidas...")
	defer medidas.Close()
}*/

/*
func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}*/
