package controllers

import (
	"net/http"

	"../structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetSlot(c *gin.Context) {
	var (
		slot   structs.Slot
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&slot).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": slot,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) AssignSlot(c *gin.Context) {
	var (
		slot structs.Slot
		// newSlot structs.Slot
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&slot, id).Error
	if err != nil {
		result = gin.H{
			"message": "not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	if slot.Availability == 0 {
		result = gin.H{
			"message": "already used",
		}
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	slot.Availability = 0
	err = idb.DB.Save(&slot).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetSlots(c *gin.Context) {
	var (
		slots  []structs.Slot
		result gin.H
	)

	idb.DB.Find(&slots)
	if len(slots) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": slots,
			"count":  len(slots),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateSlot(c *gin.Context) {
	var (
		slot   structs.Slot
		result gin.H
	)
	name := c.PostForm("name")
	slot.Availability = 0
	slot.Name = name
	idb.DB.Create(&slot)
	result = gin.H{
		"result": slot,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateSlot(c *gin.Context) {
	id := c.Query("id")
	var (
		slot    structs.Slot
		newSlot structs.Slot
		result  gin.H
	)

	err := idb.DB.First(&slot, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newSlot.Availability = 1
	newSlot.Name = slot.Name
	err = idb.DB.Model(&slot).Updates(newSlot).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteSlot(c *gin.Context) {
	var (
		slot   structs.Slot
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&slot, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&slot).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
