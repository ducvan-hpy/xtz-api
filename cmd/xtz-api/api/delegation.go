package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ducvan-hpy/xtz-api/internal/domain/models"
)

func (a *Api) GetXtzDelegations(c *gin.Context) {
	domainDelegations, err := a.repository.Delegation.List(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, NewDelegationsResponse(domainDelegations))
}

func NewDelegationsResponse(domainDelegations []models.Delegation) DelegationsResponse {
	data := make([]Delegation, 0, len(domainDelegations))
	for _, dd := range domainDelegations {
		delegation := Delegation{
			Amount:    dd.Amount,
			Block:     dd.Block,
			Delegator: dd.Delegator,
			Timestamp: dd.Timestamp,
		}
		data = append(data, delegation)
	}
	return DelegationsResponse{Data: data}
}
