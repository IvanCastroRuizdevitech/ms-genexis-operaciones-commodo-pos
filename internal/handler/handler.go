package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"ms-genexis-pos-operaciones/internal/printer"
	"ms-genexis-pos-operaciones/internal/types"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Printer *printer.Client
	Logger  Logger
}

// Logger is the minimal logger interface we need (log.Logger compatible).
type Logger interface {
	Printf(format string, v ...any)
}

func New(printerClient *printer.Client, logger Logger) *Handler {
	return &Handler{
		Printer: printerClient,
		Logger:  logger,
	}
}

// Register registers HTTP routes on the given gin engine.
func (h *Handler) Register(router *gin.Engine) {
	router.POST("/printer", h.print)
}

func (h *Handler) print(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	req := types.PrintRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.PrintResponse{
			State:   "error",
			Message: "invalid payload: " + err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	if strings.TrimSpace(req.Host) == "" {
		ctx.JSON(http.StatusBadRequest, types.PrintResponse{
			State:   "error",
			Message: "host is required",
			Status:  http.StatusBadRequest,
		})
		return
	}
	if req.Port <= 0 {
		ctx.JSON(http.StatusBadRequest, types.PrintResponse{
			State:   "error",
			Message: "port must be greater than 0",
			Status:  http.StatusBadRequest,
		})
		return
	}
	if len(req.Template) == 0 {
		ctx.JSON(http.StatusBadRequest, types.PrintResponse{
			State:   "error",
			Message: "template is required and cannot be empty",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// TODO: handle extraData.imageBase64 when required.

	var builder strings.Builder
	for _, line := range req.Template {
		builder.WriteString(line)
	}
	payload := []byte(builder.String())

	start := time.Now()
	h.Logger.Printf("print start host=%s port=%d bytes=%d", req.Host, req.Port, len(payload))
	sent, err := h.Printer.Send(requestContext(ctx), req.Host, req.Port, payload)
	duration := time.Since(start)

	if err != nil {
		h.Logger.Printf("print error host=%s port=%d sent=%d duration=%s err=%v", req.Host, req.Port, sent, duration, err)
		ctx.JSON(http.StatusInternalServerError, types.PrintResponse{
			State:   "error",
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}

	h.Logger.Printf("print success host=%s port=%d sent=%d duration=%s", req.Host, req.Port, sent, duration)
	ctx.JSON(http.StatusOK, types.PrintResponse{
		State:   "success",
		Message: "printed",
		Status:  http.StatusOK,
	})
}

func requestContext(ctx *gin.Context) context.Context {
	if ctx.Request != nil && ctx.Request.Context() != nil {
		return ctx.Request.Context()
	}
	return context.Background()
}
