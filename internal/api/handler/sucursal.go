package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/services"
	"github.com/dsniels/market/pkg"
	"gorm.io/gorm"
)

type Sucursal struct {
	sucursalSvc services.ISucursal
}

// GetSucursalsHandler godoc
// @Summary Get all Sucursales
// @Tags Sucursal
// @Accept json
// @Produce json
// @Success 200 {object} pkg.Body[[]types.Sucursal]
// @Router /api/Sucursal/GetAll [get]
func (p *Sucursal) GetSucursalsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := p.sucursalSvc.GetAll(r.Context())
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	pkg.Response(w, http.StatusOK, products)
}

// GetSucursalHandler godoc
// @Summary Get a sucursal by ID
// @Tags Sucursal
// @Accept json
// @Produce json
// @Param id path int true "Sucursal ID"
// @Success 200 {object} pkg.Body[types.Sucursal]
// @Router /api/Sucursal/GetByID/{id} [get]
func (p *Sucursal) GetSucursalHandler(w http.ResponseWriter, r *http.Request) {
	id := pkg.GetIDFromUrl[uint](r)

	product, err := p.sucursalSvc.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pkg.NotFound()
		}
		pkg.BadRequest(err.Error())
	}
	pkg.Response(w, http.StatusOK, product)
}

// UpdateSucursalHandler godoc
// @Summary Update a sucursal
// @Tags Sucursal
// @Accept json
// @Produce json
// @Param id path int true "Sucursal ID"
// @Param product body types.Sucursal true "Sucursal data"
// @Success 200 {object} pkg.Body[types.Sucursal]
// @Router /api/Sucursal/Update/{id} [post]
func (s *Sucursal) UpdateSucursalHandler(w http.ResponseWriter, r *http.Request) {
	sucursal := new(types.Sucursal)
	err := json.NewDecoder(r.Body).Decode(sucursal)
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	id := pkg.GetIDFromUrl[uint](r)
	err = s.sucursalSvc.Update(r.Context(), id, sucursal)
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	pkg.Response(w, http.StatusOK, sucursal)
}

// CreateSucursalHandler godoc
// @Summary Create a new sucursal
// @Tags Sucursal
// @Accept json
// @Produce json
// @Param product body types.Sucursal true "Sucursal data"
// @Success 200 {object} pkg.Body[types.Sucursal]
// @Router /api/Sucursal/Create [post]
func (p *Sucursal) CreateSucursalHandler(w http.ResponseWriter, r *http.Request) {
	product := new(types.Sucursal)
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	err = p.sucursalSvc.Create(r.Context(), product)
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	pkg.Response(w, http.StatusOK, product)
}

// DeleteSucursalHandler godoc
// @Summary Delete a sucursal by ID
// @Tags Sucursal
// @Accept json
// @Produce json
// @Param id path int true "Sucursal ID"
// @Success 200 {object} pkg.Body[string]
// @Router /api/Sucursal/DeleteByID/{id} [post]
func (p *Sucursal) DeleteSucursalHandler(w http.ResponseWriter, r *http.Request) {

	id := pkg.GetIDFromUrl[uint](r)
	err := p.sucursalSvc.DeleteByID(r.Context(), id)
	if err != nil {
		pkg.PanicException(400, err.Error())
	}

	pkg.Response(w, http.StatusOK, "ok")
}

func NewSucursal(sucursalSvc services.ISucursal) *Sucursal {
	return &Sucursal{
		sucursalSvc: sucursalSvc,
	}
}
