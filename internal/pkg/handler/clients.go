package handler

import (
	. "fio"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetClients(c *gin.Context) {
	var filter ClientFilter
	params := c.Request.URL.Query()
	if val := strings.Split(params.Get("name"), ","); val[0] != "" {
		filter.Name = val
	}
	if val := strings.Split(params.Get("surname"), ","); val[0] != "" {
		filter.Surname = val
	}
	if val := strings.Split(params.Get("patronymic"), ","); val[0] != "" {
		filter.Patronymic = val
	}
	if val := strings.Split(params.Get("age"), "-"); val[0] != "" {
		if len(val) == 2 {
			firstValue, err := strconv.Atoi(val[0])
			if err != nil {
				newErrorResponse(c, http.StatusBadRequest, err.Error())
				return
			}
			secondValue, err := strconv.Atoi(val[1])
			if err != nil {
				newErrorResponse(c, http.StatusBadRequest, err.Error())
				return
			}
			if firstValue > secondValue {
				temp := firstValue
				firstValue = secondValue
				secondValue = temp
			}
			filter.Age = append(filter.Age, firstValue, secondValue)
		} else {
			newErrorResponse(c, http.StatusBadRequest, "Error occured on invalid format: age")
			return
		}

	}
	if val := strings.Split(params.Get("gender"), ","); val[0] != "" {
		filter.Gender = val
	}
	if val := strings.Split(params.Get("country_id"), ","); val[0] != "" {
		filter.CountryId = val
	}
	clients, err := h.services.GetClientsByFilter(filter)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, clients)

}
