package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"getsetcar/models"
	"os"
)

func ReadCompressedJSON(filePath string) (*models.CarData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open compressed file: %w", err)
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	// Read all the data into memory
	var rawData map[string]json.RawMessage
	decoder := json.NewDecoder(gzipReader)
	if err := decoder.Decode(&rawData); err != nil {
		return nil, fmt.Errorf("failed to decode JSON into raw data: %w", err)
	}

	// Create and initialize CarData structure
	carData := &models.CarData{
		Brands: make(map[string]models.Brand),
	}

	// Process each brand
	for brandName, brandData := range rawData {
		// Initialize brand with models map
		brand := models.Brand{
			Models: make(map[string]models.Model),
		}

		// Decode brand data into temporary map
		var modelMap map[string]json.RawMessage
		if err := json.Unmarshal(brandData, &modelMap); err != nil {
			return nil, fmt.Errorf("failed to decode brand %s: %w", brandName, err)
		}

		// Process each model
		for modelName, modelData := range modelMap {
			// Initialize model
			model := models.Model{
				Variants: make(map[string]models.Variant),
			}

			// Handle images array or variant map
			if modelName == "images" {
				// This is the images array for the model
				if err := json.Unmarshal(modelData, &model.Images); err != nil {
					return nil, fmt.Errorf("failed to decode images for brand %s: %w", brandName, err)
				}
				// Skip to next iteration as we've processed this key
				continue
			}

			// Check if this is actually a variant (typical case)
			var variantMap map[string]json.RawMessage
			if err := json.Unmarshal(modelData, &variantMap); err != nil {
				// If this fails, it might be an images array or something else
				// Try to unmarshal as model directly
				if err := json.Unmarshal(modelData, &model); err != nil {
					return nil, fmt.Errorf("failed to decode model %s for brand %s: %w", modelName, brandName, err)
				}
			} else {
				// Process each variant
				for variantName, variantData := range variantMap {
					// Skip "images" as it's not a variant
					if variantName == "images" {
						if err := json.Unmarshal(variantData, &model.Images); err != nil {
							return nil, fmt.Errorf("failed to decode images for model %s: %w", modelName, err)
						}
						continue
					}

					// Decode variant data
					var variant models.Variant
					if err := json.Unmarshal(variantData, &variant); err != nil {
						return nil, fmt.Errorf("failed to decode variant %s for model %s: %w", variantName, modelName, err)
					}

					// Add variant to model
					model.Variants[variantName] = variant
				}
			}

			// Add model to brand
			brand.Models[modelName] = model
		}

		// Add brand to car data
		carData.Brands[brandName] = brand
	}

	return carData, nil
}
