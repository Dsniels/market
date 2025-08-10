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

type Categoria struct {
	categoriaSvc services.ICategoria
}

// GetCategoriasHandler godoc
// @Summary Get categorias
// @Tags categorias
// @Accept json
// @Produce json
// @Param id path int true "Categoria ID"
// @Success 200 {object} pkg.Body[types.Categoria]
// @Router /Categoria/GetByID/{id} [get]
func (c *Categoria) GetCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	id := pkg.GetIDFromUrl[uint](r)
	categoria, err := c.categoriaSvc.GetCategoria(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pkg.NotFound()
		}
		pkg.BadRequest(err.Error())
	}

	pkg.Response(w, http.StatusOK, categoria)
}

// GetCategoriasHandler godoc
// @Summary Get categorias
// @Tags categorias
// @Accept json
// @Produce json
// @Success 200 {object} pkg.Body[[]types.Categoria]
// @Router /Categoria/GetAll [get]
func (c *Categoria) GetCategoriasHandler(w http.ResponseWriter, r *http.Request) {
	categorias, err := c.categoriaSvc.GetCategorias(r.Context())
	if err != nil {
		pkg.BadRequest(err.Error())
	}

	pkg.Response(w, http.StatusOK, categorias)
}

// CreateCategoriaHandler godoc
// @Summary Create a categoria
// @Tags categorias
// @Accept json
// @Produce json
// @Param categoria body types.Categoria true "Categoria data"
// @Success 200 {object} pkg.Body[types.Categoria]
// @Router /Categoria/Create [post]
func (c *Categoria) CreateCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	categoria := new(types.Categoria)
	err := json.NewDecoder(r.Body).Decode(categoria)
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	err = c.categoriaSvc.CreateCategoria(r.Context(), categoria)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			pkg.DuplicatedKey()
		}
		pkg.BadRequest(err.Error())
	}
	pkg.Response(w, http.StatusOK, categoria)
}

// DeleteCategoriaHandler godoc
// @Summary Delete a categoria by ID
// @Tags categorias
// @Accept json
// @Produce json
// @Param id path int true "categoria ID"
// @Success 200 {object} pkg.Body[string]
// @Router /Categoria/DeleteByID/{id} [post]
func (c *Categoria) DeleteCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	id := pkg.GetIDFromUrl[uint](r)
	_, err := c.categoriaSvc.GetCategoria(r.Context(), id)
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	err = c.categoriaSvc.DeleteCategoria(r.Context(), id)
	if err != nil {
		pkg.BadRequest(err.Error())
	}

	pkg.Response(w, http.StatusOK, "ok")
}

func NewCategoria(categoriaSvc services.ICategoria) *Categoria {
	return &Categoria{
		categoriaSvc: categoriaSvc,
	}
}
