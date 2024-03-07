package models

type Propiedad struct {
	IDpropiedad string // UUID
	Domicilio   string
	Favorito    bool
	Localidad   string   // EJ: El Jaguel, Esteban Echaverría, GBA sur
	Ambientes   [5]int   // Ambientes --- {mts, ambientes, dormitorios, baños, cochera}
	Imgs        []string // URL de imagenes de la propiedad alojadas en S3
}
