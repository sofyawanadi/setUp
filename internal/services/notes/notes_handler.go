package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"setUp/internal/domain"
	"setUp/internal/utils"
)

type NoteHandler struct {
	uc  NoteUsecase
	log *zap.Logger
}

func NewNoteHandler(uc NoteUsecase, log *zap.Logger) *NoteHandler {
	return &NoteHandler{uc: uc, log: log}
}

var _ = domain.GenericResponse{}

// GetNotes godoc
// @Security BearerAuth
//
//	@Summary		Get a Notes
//	@Description	Retrieve a Notes
//	@Tags			Notes
//
// @Param id path string true "ID catatan yang ingin dicari"
//
//	@Accept			json
//	@Produce		json
//	@Success		200	{object} utils.SuccessResponse
//	@Failure		500	{object} utils.BaseResponse
//	@Router			/notes/{id} [get]
func (h *NoteHandler) GetNote(c *gin.Context) {
	id := c.Param("id")
	note, err := h.uc.GetByID(c, id)
	if err != nil {
		utils.ErrorResp(c, http.StatusNotFound, err.Error())
		return
	}
	utils.SuccessResp(c, "note fetched successfully", note)
}

// GetNotes godoc
// @Security BearerAuth
//
//	@Summary		Get list of Notes
//	@Description	Retrieve a list of all Notes
//	@Tags			Notes
//	@Accept			json
//	@Produce		json
//	@Success		200	{object} utils.SuccessWithPaginationResponse
//	@Failure		500	{object} utils.BaseResponse
//	@Router			/notes [get]
func (h *NoteHandler) GetNotes(c *gin.Context) {
	filters := utils.GetFilter(c)
	params := utils.QueryParams{
		Filters:   filters,
		SortBy:    c.DefaultQuery("sort_by", "created_at"),
		SortOrder: c.DefaultQuery("sort_order", "desc"),
		Page:      int64(utils.ParseInt(c.DefaultQuery("page", "1"))),
		PageSize:  int64(utils.ParseInt(c.DefaultQuery("page_size", "10"))),
	}
	notes, count, err := h.uc.GetAll(c, params)
	if err != nil {
		utils.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessWithPaginationResp(c, "notes fetched successfully", notes, params.Page, params.PageSize, count)
}

type NoteRequest struct {
	Title   string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
}

// CreateNote godoc
// @Security BearerAuth
// @Summary		Create a new notes
// @Description	Add a new notes to the system
// @Tags			Notes
// @Accept			json
// @Produce		json
// @Param			notes	body	NoteRequest	true	"notes data"
// @Success		200		{object} utils.SuccessResponse
// @Failure		500		{object} utils.BaseResponse
// @Failure		400		{object} utils.ValidationErrorResponse
// @Router			/notes [post]
func (h *NoteHandler) CreateNote(c *gin.Context) {
	var req NoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResp(c, http.StatusBadRequest, err.Error())
		return
	}
	// Panggil helper untuk validasi
	if !utils.ValidateRequest(&req, c, h.log) {
		return
	}

	note, err := h.uc.Create(c, &req)
	if err != nil {
		utils.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResp(c, "req created successfully", map[string]interface{}{
		"data": note,
	})
}

type NoteUpdateRequest struct {
	Id      uuid.UUID `json:"id" form:"id" validate:"required"`
	Title   string    `json:"title" form:"title" validate:"required"`
	Content string    `json:"content" form:"content" validate:"required"`
}

// UpdateNote godoc
// @Security BearerAuth
// @Summary		update a new notes
// @Description	Add a new notes to the system
// @Tags			Notes
// @Accept			json
// @Produce		json
// @Param			notes	body	NoteUpdateRequest	true	"notes data"
// @Success		200		{object} utils.SuccessResponse
// @Failure		500		{object} utils.BaseResponse
// @Failure		400		{object} utils.ValidationErrorResponse
// @Router			/notes [put]
func (h *NoteHandler) UpdateNote(c *gin.Context) {
	var req NoteUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResp(c, http.StatusBadRequest, err.Error())
		return
	}
	// Panggil helper untuk validasi
	if !utils.ValidateRequest(&req, c, h.log) {
		return
	}

	note, err := h.uc.Update(c, &req)
	if err != nil {
		utils.ErrorResp(c, http.StatusInternalServerError, "failed to update note")
		return
	}
	utils.SuccessResp(c, "req updated successfully", note)
}

//	DeleteNote godoc
//	@Security BearerAuth
//	@Summary		delete a new notes
//	@Description	delete a new notes to the system
//	@Tags			Notes
//
// @Param id path string true "ID catatan yang akan dihapus"
//
//	@Accept			json
//	@Produce		json
//	@Success		200		{object} utils.SuccessResponse
//	@Failure		500		{object} utils.BaseResponse
//	@Router			/notes/{id} [delete]
func (h *NoteHandler) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	if err := h.uc.Delete(c, id); err != nil {
		utils.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResp(c, "note deleted successfully", nil)
}
