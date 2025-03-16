package handlers

import (
	"encoding/json"
	"fmt"
	"getsetcar/models"
	"getsetcar/utils"
	"github.com/gofiber/fiber/v2"
)

type CarHandler struct {
	CarsData models.CarData
}

func NewCarHandler(carsData models.CarData) *CarHandler {
	return &CarHandler{
		CarsData: carsData,
	}
}

func (h *CarHandler) GetCarsForBrand(c *fiber.Ctx) error {
	param := c.Params("brand")
	param = param + "-cars"
	var response []models.BrandSearchResponse
	for brand, carModels := range h.CarsData.Brands {
		//fmt.Println(brand)
		if brand == param {
			PrettyPrint(carModels.Models)
			for variantName, car := range carModels.Models {
				mainImage, err := getMainImage(car.Images)
				if err != nil {
					fmt.Println(fmt.Sprintf("failed to get mainImage for - [%s]", variantName))
					continue
				}
				res := models.BrandSearchResponse{Image: mainImage.Path, VariantName: variantName}
				res.Price = getLowestBasicPrice(car.Variants)
				response = append(response, res)
			}
		}
	}
	c.Status(fiber.StatusOK)
	return c.JSON(response)
}

func (h *CarHandler) GetModel(c *fiber.Ctx) error {
	param := c.Params("brand")
	carBrand := param + "-cars"

	carModel := c.Params("model")
	brands := h.CarsData.Brands
	allModelsOfBrand := brands[carBrand]
	c.Status(fiber.StatusOK)
	return c.JSON(allModelsOfBrand.Models[carModel])
}

func getMainImage(images []models.Image) (models.Image, error) {
	for _, image := range images {
		if image.CategoryID == 0 {
			return image, nil
		}
	}
	return models.Image{}, fmt.Errorf("no main image found")
}

func getLowestBasicPrice(variants map[string]models.Variant) string {
	priceList := make([]string, 0)
	for _, variantData := range variants {
		priceList = append(priceList, variantData.BasicPrice)
	}
	return utils.GetLowestPrice(priceList)
}

func PrettyPrint(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}
