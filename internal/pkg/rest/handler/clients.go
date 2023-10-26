package handler

import (
	. "fio"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	defaultPage = 1
)

func (h *Handler) CreateClient(c *gin.Context) {
	var client Client
	if err := c.BindJSON(&client); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Record.CreateClient(client)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) GetClients(c *gin.Context) {
	var filter ClientFilter
	var page int
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

	p := params.Get("p")
	if p != "" {
		parse, err := strconv.Atoi(p)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		page = parse
	} else {
		page = defaultPage
	}
	if page < defaultPage {
		page = defaultPage
	}
	clients, err := h.services.GetClientsByFilter(filter, page)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, clients)

}

func (h *Handler) UpdateClientRecord(c *gin.Context) {
	var client ClientUpdate

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.BindJSON(&client); err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	if err := h.services.UpdateClientRecord(id, client); err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("successfully updated client with id = %d", id),
	})
}

func (h *Handler) DeleteClientById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.DeleteClientById(id); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": fmt.Sprintf("successfully deleted client with id = %d", id),
	})

}
