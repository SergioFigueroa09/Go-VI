package multimedia

import "fmt"

type Multimedia interface{
	Mostrar()
}

type ContenidoWeb struct{
	SliceMulti []Multimedia
}

func (mm *ContenidoWeb) Mostrar() {
	for _,m := range mm.SliceMulti{
		m.Mostrar()
	}
}


type Imagen struct{
	Titulo string
	Formato string
	Canales int
}

func (i *Imagen) Mostrar() {
	fmt.Println("IMAGEN--------------")
	fmt.Println("Titulo: ",i.Titulo)
	fmt.Println("Formato: ",i.Formato)
	fmt.Println("Canales: ",i.Canales)
}



type Audio struct{
	Titulo string
	Formato string
	Duracion int
}

func (a *Audio) Mostrar(){
	fmt.Println("AUDIO---------------")
	fmt.Println("Titulo: ",a.Titulo)
	fmt.Println("Formato: ",a.Formato)
	fmt.Println("Duración (seg): ",a.Duracion)
}



type Video struct{
	Titulo string
	Formato string
	Frames int
}

func (v *Video) Mostrar(){
	fmt.Println("VIDEO---------------")
	fmt.Println("Titulo: ",v.Titulo)
	fmt.Println("Formato: ",v.Formato)
	fmt.Println("Frames: ",v.Frames)
}

