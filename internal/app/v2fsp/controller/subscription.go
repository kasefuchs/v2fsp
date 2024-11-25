package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/kasefuchs/v2fsp/internal/pkg/repository"
)

func SubscriptionGet(ctx *fiber.Ctx) error {
	items := repository.Items()

	buf := bytes.NewBuffer(nil)
	for _, item := range items {
		marshal, _ := json.Marshal(item)
		_, _ = fmt.Fprintln(buf, string(marshal))
	}

	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	return ctx.SendString(str)
}
