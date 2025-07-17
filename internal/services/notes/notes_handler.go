package services

import (
	"net/http"
	"setUp/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type NoteHandler struct {
	uc  NoteUsecase
	log *zap.Logger
}

func NewNoteHandler(uc NoteUsecase, log *zap.Logger) *NoteHandler {
	return &NoteHandler{uc: uc, log: log}
}

func (h *NoteHandler) GetNote(c *gin.Context) {
	id := c.Param("id")
	note, err := h.uc.GetByID(c, id)
	if err != nil {
		utils.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResp(c, "note fetched successfully", note)
}

func (h *NoteHandler) GetNotes(c *gin.Context) {
	filters := utils.GetFilter(c)
	params := utils.QueryParams{
		Filters:   filters,
		SortBy:    c.DefaultQuery("sort_by", "created_at"),
		SortOrder: c.DefaultQuery("sort_order", "desc"),
		Page:      int64(utils.ParseInt(c.DefaultQuery("page", "1"))),
		PageSize:  int64(utils.ParseInt(c.DefaultQuery("page_size", "10"))),
	}
	notes,count, err := h.uc.GetAll(c,params)
	if err != nil {
		utils.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessWithPaginationResp(c, "notes fetched successfully", notes,params.Page, params.PageSize, count)
}
type NoteRequest struct {
	Title    string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
}


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

	note, err := h.uc.Create(c, &req);
	if err != nil {
		utils.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResp(c, "req created successfully", map[string]interface{}{
		"data":note,
	})
}
type NoteUpdateRequest struct {
	Id    uuid.UUID `json:"id" form:"id" validate:"required"`
	Title    string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
}

func (h *NoteHandler) UpdateNote(c *gin.Context) {
	var req NoteUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResp(c, http.StatusBadRequest, err.Error())
		return
	}
	note, err := h.uc.Update(c, &req);
	if  err != nil {
		utils.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResp(c, "req updated successfully", map[string]interface{}{
		"data":note,
	})
}

func (h *NoteHandler) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	if err := h.uc.Delete(c, id); err != nil {
		utils.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResp(c, "note deleted successfully", nil)
}
