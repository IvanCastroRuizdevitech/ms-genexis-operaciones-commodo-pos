package routes_reports

import (
	handler_reports "ms-genexis-pos-operaciones/context/reports/presentation/handler_reports"

	"github.com/gin-gonic/gin"
)

func LoadReportsRoutes(router *gin.RouterGroup) {
	reportsGroup := router.Group("/reports")
	{
		reportsGroup.GET(
			"/day-closing-report/:fecha",
			handler_reports.GetDayClosingReportHandler,
		)
		reportsGroup.GET(
			"/fuel-report/:fecha",
			handler_reports.GetFuelReportHandler,
		)
		reportsGroup.POST(
			"/closing-novelties",
			handler_reports.GetDailyNoveltiesHandler,
		)
		reportsGroup.POST(
			"/tanks/print-event",
			handler_reports.CreateTankPrintEventHandler,
		)
		reportsGroup.GET(
			"/movement-types",
			handler_reports.GetMovementTypesHandler,
		)
		reportsGroup.GET(
			"/tanks",
			handler_reports.GetTanksHandler,
		)
		reportsGroup.POST(
			"/shift-summary",
			handler_reports.GetShiftSummaryHandler,
		)
		reportsGroup.POST(
			"/shift-consolidated",
			handler_reports.GetShiftConsolidatedHandler,
		)
	}
}
