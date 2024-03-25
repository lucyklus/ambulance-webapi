package ambulance_wl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Nasledujúci kód je kópiou vygenerovaného a zakomentovaného kódu zo súboru api_ambulance_conditions.go
func (this *implAmbulanceConditionsAPI) GetConditions(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		result := ambulance.PredefinedConditions
		if result == nil {
			result = []Condition{}
		}
		return nil, result, http.StatusOK
	})
}

func (this *implAmbulanceConditionsAPI) CreateCondition(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		condition := Condition{}
		err := ctx.BindJSON(&condition)
		if err != nil {
			return nil, gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}
		ambulance.PredefinedConditions = append(ambulance.PredefinedConditions, condition)
		return ambulance, condition, http.StatusCreated
	})
}

func (this *implAmbulanceConditionsAPI) UpdateCondition(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		condition := Condition{}
		err := ctx.BindJSON(&condition)
		if err != nil {
			return nil, gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}
		conditionCode := ctx.Param("conditionCode")
		for i, c := range ambulance.PredefinedConditions {
			if c.Code == conditionCode {
				if condition.Code != "" {
					ambulance.PredefinedConditions[i].Code = condition.Code
				}
				if condition.Reference != "" {
					ambulance.PredefinedConditions[i].Reference = condition.Reference
				}
				if condition.Value != "" {
					ambulance.PredefinedConditions[i].Value = condition.Value
				}
				if condition.TypicalDurationMinutes != 0 {
					ambulance.PredefinedConditions[i].TypicalDurationMinutes = condition.TypicalDurationMinutes
				}
				return ambulance, condition, http.StatusOK
			}
		}
		return nil, gin.H{
			"status":  "Not Found",
			"message": "Condition not found",
			"error":   "Condition not found",
		}, http.StatusNotFound
	})
}

func (this *implAmbulanceConditionsAPI) GetCondition(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		conditionCode := ctx.Param("conditionCode")
		for _, c := range ambulance.PredefinedConditions {
			if c.Code == conditionCode {
				return nil, c, http.StatusOK
			}
		}
		return nil, gin.H{
			"status":  "Not Found",
			"message": "Condition not found",
			"error":   "Condition not found",
		}, http.StatusNotFound
	})
}

func (this *implAmbulanceConditionsAPI) DeleteCondition(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		conditionCode := ctx.Param("conditionCode")
		for i, c := range ambulance.PredefinedConditions {
			if c.Code == conditionCode {
				ambulance.PredefinedConditions = append(ambulance.PredefinedConditions[:i], ambulance.PredefinedConditions[i+1:]...)
				return ambulance, nil, http.StatusNoContent
			}
		}
		return nil, gin.H{
			"status":  "Not Found",
			"message": "Condition not found",
			"error":   "Condition not found",
		}, http.StatusNotFound
	})
}
