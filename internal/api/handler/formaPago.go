package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/services"
	"github.com/dsniels/market/pkg"
	"gorm.io/gorm"
)

type FormaPago struct {
	FormaPagovc services.IFormaPago
}

// GetFormaPagosHandler godoc
// @Summary Get FormaPago
// @Tags FormaPago
// @Accept json
// @Produce json
// @Param id path int true "FormaPago ID"
// @Success 200 {object} pkg.Body[types.FormaPago]
// @Router /api/FormaPago/GetByID/{id} [get]
func (c *FormaPago) GetFormaPagoHandler(w http.ResponseWriter, r *http.Request) {
	id := pkg.GetIDFromUrl[uint](r)
	FormaPago, err := c.FormaPagovc.GetFormaPago(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pkg.NotFound()
		}
		pkg.BadRequest(err.Error())
	}

	pkg.Response(w, http.StatusOK, FormaPago)
}

// GetFormaPagosHandler godoc
// @Summary Get FormaPago
// @Tags FormaPago
// @Accept json
// @Produce json
// @Success 200 {object} pkg.Body[[]types.FormaPago]
// @Router /api/FormaPago/GetAll [get]
func (c *FormaPago) GetFormasPagoHandler(w http.ResponseWriter, r *http.Request) {
	FormasPago, err := c.FormaPagovc.GetFormaPagos(r.Context())
	if err != nil {
		pkg.BadRequest(err.Error())
	}

	pkg.Response(w, http.StatusOK, FormasPago)
}

// CreateFormaPagoHandler godoc
// @Summary Create a FormaPago
// @Tags FormaPago
// @Accept json
// @Produce json
// @Param FormaPago body types.FormaPago true "FormaPago data"
// @Success 200 {object} pkg.Body[types.FormaPago]
// @Router /api/FormaPago/Create [post]
func (c *FormaPago) CreateFormaPagoHandler(w http.ResponseWriter, r *http.Request) {
	FormaPago := new(types.FormaPago)
	err := json.NewDecoder(r.Body).Decode(FormaPago)
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	err = c.FormaPagovc.CreateFormaPago(r.Context(), FormaPago)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			pkg.DuplicatedKey()
		}
		pkg.BadRequest(err.Error())
	}
	pkg.Response(w, http.StatusOK, FormaPago)
}

// DeleteFormaPagoHandler godoc
// @Summary Delete a FormaPago by ID
// @Tags FormaPago
// @Accept json
// @Produce json
// @Param id path int true "FormaPago ID"
// @Success 200 {object} pkg.Body[string]
// @Router /api/FormaPago/DeleteByID/{id} [post]
func (c *FormaPago) DeleteFormaPagoHandler(w http.ResponseWriter, r *http.Request) {
	id := pkg.GetIDFromUrl[uint](r)
	_, err := c.FormaPagovc.GetFormaPago(r.Context(), id)
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	err = c.FormaPagovc.DeleteFormaPago(r.Context(), id)
	if err != nil {
		pkg.BadRequest(err.Error())
	}

	pkg.Response(w, http.StatusOK, "ok")
}

func NewFormaPago(FormaPagovc services.IFormaPago) *FormaPago {
	return &FormaPago{
		FormaPagovc: FormaPagovc,
	}
}
