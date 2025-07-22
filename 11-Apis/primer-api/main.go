package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// datos
var albums = []album{
	{ID: "1", Title: "Blue TRain", Artist: "John Coltrade", Price: 56.99},
	{ID: "2", Title: "Kind of Blue", Artist: "Miles Davis", Price: 49.99},
	{ID: "3", Title: "A Love Supreme", Artist: "John Coltrane", Price: 52.50},
	{ID: "4", Title: "Time Out", Artist: "The Dave Brubeck Quartet", Price: 45.00},
	{ID: "5", Title: "Getz/Gilberto", Artist: "Stan Getz & João Gilberto", Price: 47.75},
	{ID: "6", Title: "Mingus Ah Um", Artist: "Charles Mingus", Price: 48.25},
}

// Manejador GET para tomar todos los álbumes
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) // Devuelve los álbumes con formato inde
}

// Manejador POST para agregar un nuevo álbum
func postAlbums(c *gin.Context) {
	var newAlbum album
	// Manejamos el error.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)          // Agrega el nuevo álbum al slice
	c.IndentedJSON(http.StatusCreated, albums) // Agrega el nuevo álbum al slice
}

// Manejador GET para retornar un álbum específico por ID
func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

func main() {
	router := gin.Default()

	// Ruta base que retorna "Holamundo"
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Holamundo",
		})
	})

	//definimos las rutas
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.Run("localhost:8081")
}
