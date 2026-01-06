package handler

import (
	"rapnews/internal/core/service"

	"github.com/gofiber/fiber/v2"
)

type ContentHandler interface {
	GetContents(c *fiber.Ctx) error
	GetContentByID(c *fiber.Ctx) error
	CreateContent(c *fiber.Ctx) error
	EditContentByID(c *fiber.Ctx) error
	DeleteContentByID(c *fiber.Ctx) error
	UploadImageR2(c *fiber.Ctx) error
}

type contentHandler struct {
	contentService service.ContentService
}

// CreateContent implements ContentHandler.
func (*contentHandler) CreateContent(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteContentByID implements ContentHandler.
func (*contentHandler) DeleteContentByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// EditContentByID implements ContentHandler.
func (*contentHandler) EditContentByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetContentByID implements ContentHandler.
func (*contentHandler) GetContentByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetContents implements ContentHandler.
func (*contentHandler) GetContents(c *fiber.Ctx) error {
	
	panic("unimplemented")
}

// UploadImageR2 implements ContentHandler.
func (*contentHandler) UploadImageR2(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewContentHandler(contentService service.ContentService) ContentHandler {
	return &contentHandler{contentService: contentService}
}
