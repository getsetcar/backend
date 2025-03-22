package models

type AllBrandsResponse struct {
	Brand string `json:"brand"`
	Logo  string `json:"logo"`
}

type BrandSearchResponse struct {
	VariantName string `json:"variantName"`
	Image       string `json:"image"`
	Price       string `json:"price"`
}

// CarData holds the entire car database
type CarData struct {
	Brands map[string]Brand
}

// Brand represents a car manufacturer
type Brand struct {
	Models map[string]Model
}

// Model represents a specific car model from a brand
type Model struct {
	Variants map[string]Variant `json:",omitempty"`
	Images   []Image            `json:"images,omitempty"`
}

// Variant represents a specific variant of a car model
type Variant struct {
	Specifications        map[string]map[string]string `json:"specifications,omitempty"`
	BasicPrice            string                       `json:"basic_price,omitempty"`
	AlternativeBasicPrice string                       `json:"   ,omitempty"`
	CityPrices            []CityPrice                  `json:"city_price,omitempty"`
	Colors                []Color                      `json:"colours,omitempty"`
}

type CityPrice struct {
	OnRoadPrice            string `json:"onRoadPrice,omitempty"`
	MaskingName            string `json:"maskingName,omitempty"`
	UnformattedOnRoadPrice string `json:"unformattedOnRoadPrice,omitempty"`
	ID                     int    `json:"id,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Color struct {
	VersionID    int    `json:"versionId,omitempty"`
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	ImagePath    string `json:"imagePath,omitempty"`
	HexCode      string `json:"hexCode,omitempty"`
	AltImageName string `json:"altImageName,omitempty"`
}

type Image struct {
	Path              string `json:"path,omitempty"`
	ID                int    `json:"id,omitempty"`
	CategoryID        int    `json:"categoryId,omitempty"`
	ImageTitle        string `json:"imageTitle,omitempty"`
	USPText           string `json:"uspText,omitempty"`
	TagID             int    `json:"tagId,omitempty"`
	AltImageName      string `json:"altImageName,omitempty"`
	Name              string `json:"name,omitempty"`
	ModelTour         string `json:"modelTour,omitempty"`
	IsModelTourActive bool   `json:"isModelTourActive,omitempty"`
}
