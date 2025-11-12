package routes_envelopes

import (
	"ms-genexis-pos-operaciones/context/envelopes/domain/entities"
	handler_enveloper "ms-genexis-pos-operaciones/context/envelopes/presentation/handler_envelopes"
	presentation_api_middlewares "ms-genexis-pos-operaciones/presentation/api/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadEnvelopesRoutes(router *gin.RouterGroup) {
	envelopesGroup := router.Group("/envelopes")
	{
		envelopesGroup.POST(
			"/total",
			presentation_api_middlewares.ValidateBodyStruct[entities.EnvelopesTotalRequest](),
			handler_enveloper.TotalEnvelopesByJournalPromoterHandler,
		)
	}
}
