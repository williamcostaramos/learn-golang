package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
)

type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

type machine struct {
	Cpu      string
	Core     int64
	Velocity float64
}

var albums = []album{
	{ID: "1", Title: "blue try", Artist: "William", Price: 59.99},
	{ID: "2", Title: "Emoções", Artist: "Aviões", Price: 59.99},
}

func main() {
	log := log.Default()
	log.Println("INICIANDO SERVIDOR DE APLICAÇÃO")
	albums = []album{}
	router := gin.Default()
	router.GET("/test", index)
	router.GET("/machine", handleMachine)
	router.Run("localhost:8081")
	log.Println("Aguardando conexões")
}

func index(c *gin.Context) {
	if len(albums) > 0 {
		c.IndentedJSON(http.StatusOK, albums)
	}
	c.IndentedJSON(http.StatusNotFound, albums)
}

func handleMachine(c *gin.Context) {
	var maquinas = []machine{}
	cpu, err := cpu.Info()

	if err != nil {
		return
	}

	for _, info := range cpu {
		maquina := machine{
			Cpu:      info.ModelName,
			Core:     int64(info.Cores),
			Velocity: info.Mhz,
		}
		maquinas = append(maquinas, maquina)

	}
	c.IndentedJSON(http.StatusOK, maquinas)

}
