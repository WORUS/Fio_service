package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ClientFilter struct {
	Name         []string
	Surname      []string
	Patronymic   []string
	Age          []int
	AgeIntervals []int
	Gender       []string
	CountryId    []string
}

func (h *Handler) GetClients(c *gin.Context) {
	var filter ClientFilter
	params := c.Request.URL.Query()
	if val := strings.Split(params.Get("name"), ","); val != nil {
		filter.Name = val
	}
	if val := strings.Split(params.Get("surname"), ","); val != nil {
		filter.Surname = val
	}
	if val := strings.Split(params.Get("patronymic"), ","); val != nil {
		filter.Patronymic = val
	}
	if val := strings.Split(params.Get("age"), ","); len(val) != 0 {
		for i := range val {
			intervals := strings.Split(val[i], "-")
			if len(intervals) == 2 {
				firstValue, err := strconv.Atoi(intervals[0])
				if err != nil {
					newErrorResponse(c, http.StatusBadRequest, err.Error())
					return
				}
				secondValue, err := strconv.Atoi(intervals[1])
				if err != nil {
					newErrorResponse(c, http.StatusBadRequest, err.Error())
					return
				}
				if firstValue > secondValue {
					temp := firstValue
					firstValue = secondValue
					secondValue = temp
				}
				for ; firstValue <= secondValue; firstValue++ {
					filter.Age = append(filter.Age, firstValue)
				}
			} else {
				res, err := strconv.Atoi(val[i])
				//filter.Age[i], err = strconv.Atoi(val[i])
				if err != nil {
					newErrorResponse(c, http.StatusBadRequest, err.Error())
					return
				}
				filter.Age = append(filter.Age, res)
			}
		}
	}
	if val := strings.Split(params.Get("gender"), ","); val != nil {
		filter.Gender = val
	}
	if val := strings.Split(params.Get("country_id"), ","); val != nil {
		filter.CountryId = val
	}

	c.JSON(http.StatusOK, filter)

}
