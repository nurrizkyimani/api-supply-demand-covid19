package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/api-supply-demand-covid19/model"
	"github.com/ngavinsir/api-supply-demand-covid19/models"
)

// DonationResource struct
type DonationResource struct {
	*model.DonationDataStore
	*model.UserDatastore
}

func (store *DonationResource) router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(AuthMiddleware)
	r.Use(UserCtx(store.UserDatastore))
	r.Post("/", CreateDonation(store.DonationDataStore))

	return r
}

// CreateDonation return donations
func CreateDonation(repo interface {
	model.HasCreateDonation
}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &CreateDonationRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		user, _ := r.Context().Value(UserCtxKey).(*models.User)
		if user.Role != model.RoleDonator && user.Role != model.RoleAdmin {
			render.Render(w, r, ErrUnauthorized(ErrInvalidRole))
			return
		}

		request, err := repo.CreateDonation(r.Context(), data.DonationItems, user.ID)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, request)
	}
}

// CreateDonationRequest struct
type CreateDonationRequest struct {
	DonationItems []*models.DonationItem `json:"donationItems"`
}

// Bind RegisterRequest ([]RequestItem) [Required]
func (req *CreateDonationRequest) Bind(r *http.Request) error {
	if req.DonationItems == nil || len(req.DonationItems) == 0 {
		return ErrMissingReqFields
	}

	return nil
}