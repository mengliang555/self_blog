package boot

import (
	"github.com/gin-gonic/gin"
)

type innerEngine struct {
	ginEngine *gin.Engine
}

var engine *innerEngine

func init() {
	engine = &innerEngine{
		ginEngine: gin.Default(),
	}
}

func RefEngine() *innerEngine {
	return engine
}

func (engine *innerEngine) GetGinEngine() *gin.Engine {
	return engine.ginEngine
}
