package categoryhandlers

import (
	"net/http"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/repository"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		helpers.ErrorResponse(
			w,
			http.StatusMethodNotAllowed,
			"Method not allowed",
		)

		return
	}

	categories, err := repository.GetCategories()
	if err != nil {

		helpers.ErrorResponse(
			w,
			http.StatusInternalServerError,
			"Failed to get categories",
		)

		return
	}

	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Categories retrieved successfully",
		categories,
	)
}
