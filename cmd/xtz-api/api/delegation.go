package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ducvan-hpy/xtz-api/internal/domain/model"
)

func (a *Api) GetXtzDelegations(c *gin.Context, params GetXtzDelegationsParams) {
	domainDelegations := a.repository.Delegation.List(c, params.Year)
	c.JSON(http.StatusOK, NewDelegationsResponse(domainDelegations))
}

func NewDelegationsResponse(domainDelegations []model.Delegation) DelegationsResponse {
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
