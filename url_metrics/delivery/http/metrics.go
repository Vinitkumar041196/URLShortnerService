package http

import (
	"net/http"
	"strconv"
	"url-shortener/domain"

	"github.com/gin-gonic/gin"
)

type DomainMetricsHttpHandler struct {
	service domain.DomainMetricsService
}

// returns a new http handler for domain metrics api
func NewDomainMetricsHttpHandler(srvc domain.DomainMetricsService) *DomainMetricsHttpHandler {
	return &DomainMetricsHttpHandler{service: srvc}
}

// GetTopDomains godoc
// @Summary List Top Shortened Domains
// @Description Returns a list of domains shortened maximum number of times
// @Tags Metrics
// @Produce json
// @Param limit query int false "optional query param to specify how many domains to get"
// @Success 200 {object} GetTopDomainsSuccessResponse
// @Failure 500 {object} GetTopDomainsErrorResponse
// @Router /metrics/domains/top [get]
func (h *DomainMetricsHttpHandler) GetTopDomains(c *gin.Context) {
	limit := 3
	//use limit from query param if exists
	if limitStr := c.Query("limit"); limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil {
			limit = l
		}
	}

	topNDomains, err := h.service.GetTopDomains(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GetTopDomainsErrorResponse{Error: err.Error(), Message: "FAILED"})
		return
	}

	c.JSON(http.StatusOK, GetTopDomainsSuccessResponse{Data: topNDomains, Message: "SUCCESS"})
}
