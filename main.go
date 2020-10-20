package main

import (
	"fmt"

	 "./Multimedias"
)

func main() {
	opcion := 0
	var aux1 string
	var aux2 string
	var aux3 int

	mm := multimedia.ContenidoWeb{
		SliceMulti: []multimedia.Multimedia{},
	}

	for opcion != 9 {
		fmt.Println("--CAPTURAR NUEVA MULTIMEDIA--")
		fmt.Println("1.---------IMAGEN")
		fmt.Println("2.---------AUDIO")
		fmt.Println("3.---------VIDEO")
		fmt.Println("4.---------MOSTRAR TODO")
		fmt.Println("9.---------SALIR")
		fmt.Println("Elija una opción: ")
		fmt.Scan(&opcion)
		switch opcion {
		case 1:
			fmt.Println("Ingrese el Titulo (sin espacios): ")
			fmt.Scan(&aux1)
			fmt.Println("Ingrese el Formato: ")
			fmt.Scan(&aux2)
			fmt.Println("Ingrese el numero de Canales: ")
			fmt.Scan(&aux3)

			mm.SliceMulti = append(mm.SliceMulti, &multimedia.Imagen{aux1, aux2, aux3})
			break
		case 2:
			fmt.Println("Ingrese el Titulo (sin espacios): ")
			fmt.Scan(&aux1)
			fmt.Println("Ingrese el Formato: ")
			fmt.Scan(&aux2)
			fmt.Println("Ingrese la Duración en segundos: ")
			fmt.Scan(&aux3)

			mm.SliceMulti = append(mm.SliceMulti, &multimedia.Audio{aux1, aux2, aux3})
			break
		case 3:
			fmt.Println("Ingrese el Titulo (sin espacios): ")
			fmt.Scan(&aux1)
			fmt.Println("Ingrese el Formato: ")
			fmt.Scan(&aux2)
			fmt.Println("Ingrese el numero de Frames: ")
			fmt.Scan(&aux3)

			mm.SliceMulti = append(mm.SliceMulti, &multimedia.Video{aux1, aux2, aux3})
			break
		case 4:
			mm.Mostrar()
			break
		case 9:
			return
		default:
			fmt.Println("Escoja una opción valida.")
		}
	}
}
