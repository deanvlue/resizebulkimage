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

	// pasas el nombre del archivo
	//define los tamaños en los que lo va a generar
	// guarda los 17 archivos en el folder newImages

	flag.Parse()
	originalArt := flag.Arg(0)

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
	/*artes, err := ioutil.ReadDir(root)
		if err != nil {
	y
		}*/

	artes, err := os.Open(originalArt)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer artes.Close()

	fmt.Printf("Abriendo Archivo %s\n", originalArt)

	//rompemos en nombre del archivo en sus componenetes mínimos:
	splitPWD := strings.Split(artes.Name(), "/")
	fileName := strings.Split(splitPWD[2], ".")
	fmt.Println(fileName)

	//defer imageReader.Close()

	im, _, err := image.Decode(artes)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for i := 0; i < len(medidas.Medidas); i++ {

		newImageName := fileName[0] + "_" + medidas.Medidas[i].ImageType + "." + fileName[1]

		fmt.Println("Resizing image to W:", strconv.FormatUint(uint64(medidas.Medidas[i].Width), 10))
		//fmt.Println("Resizing image to H:", strconv.FormatUint(medidas.Medidas[i].Height, 10))
		fmt.Printf("Imagen ./newimages/%s_%s.%s Guardada\n", fileName[0], medidas.Medidas[i].ImageType, fileName[1])
		//fmt.Println("Tipo de Imagen: " + medidas.Medidas[i].ImageType)
		newImage := resize.Resize(medidas.Medidas[i].Width, 0, im, resize.Lanczos3)
		if fileName[1] == "jpg" {
			out, err := os.Create("./newimages/" + newImageName)
			err = jpeg.Encode(out, newImage, nil)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		}

	}

}
