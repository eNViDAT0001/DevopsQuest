package order

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/paging"
	"github.com/eNViDAT0001/GolangAdventure/external/paging/paging_query"
	"github.com/eNViDAT0001/GolangAdventure/external/request"
	"github.com/eNViDAT0001/GolangAdventure/internal/order/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s *orderHandler) ListByProviderID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := paging_query.GetPagingParams(cc.Context, entities.Order{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		providerID, err := strconv.Atoi(cc.Param("provider_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		orders, total, err := s.orderUC.ListByProviderID(newCtx, uint(providerID), paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CursorPaging && len(orders) > 0 {
			paginator.Marker = int(orders[len(orders)-1].ID)
		}

		cc.OkPaging(paginator, orders)
	}
}
