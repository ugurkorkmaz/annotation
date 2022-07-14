package example

import (
	"github.com/gofiber/fiber/v2"
)

/**
 * @Route("/foo", name="foo", methods={"GET"})
 */
func Foo(*fiber.Ctx) error {
	return nil
}

/**
 * @Route("/bar", name=" bar", methods={"GET"})
 */
func Bar(*fiber.Ctx) error {
	return nil
}
