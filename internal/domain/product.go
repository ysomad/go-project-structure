package domain

type Product struct {
	ID            string
	Title         string
	Description   string
	Color         string
	Label         string
	Brand         Brand
	Price         Price
	VariantSizes  []string
	ImageURLs     []string
	MerchantIcons []string
}

type Brand struct {
	ID   string
	Name string
}

type (
	Price struct {
		CurrentPrice       int32
		OldPrice           int32
		CurrencyCode       string
		DiscountPercentage int32
	}

	PriceWithShipment struct {
		Price
		ShipmentPrice string
	}
)

type (
	DetailedProduct struct {
		ID          string
		Title       string
		Description string
		Color       string
		Label       string
		Brand       Brand
		Price       Price
		Variants    []Variant
		Categories  []Category
		ImageURLs   []string
	}

	Category struct {
		ID       string
		ParentID string
		Name     string
	}

	Variant struct {
		Size   string
		Price  PriceWithShipment
		Offers []Offer
	}

	Merchant struct {
		Name     string
		IconURL  string
		LogoURL  string
		LogoText string
		Official bool
	}

	Offer struct {
		Merchant Merchant
		InStock  bool
		LastLeft bool
	}
)

type ProductList struct {
	Items         []*Product
	TotalPages    int32
	NextPageToken string
}

type ProductFilters struct {
	Gender     string
	Categories []string
	Brands     []string
	Colors     []string
	Sizes      []string
	Labels     []string
	MinPrice   int32
	MaxPrice   int32
}

type ProductSort int8

const (
	ProductSortDESC ProductSort = iota
	ProductSortASC
)
