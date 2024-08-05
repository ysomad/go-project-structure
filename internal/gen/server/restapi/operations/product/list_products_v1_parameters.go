// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewListProductsV1Params creates a new ListProductsV1Params object
// with the default values initialized.
func NewListProductsV1Params() ListProductsV1Params {

	var (
		// initialize parameters with default values

		pageSizeDefault = int32(60)
	)

	return ListProductsV1Params{
		PageSize: &pageSizeDefault,
	}
}

// ListProductsV1Params contains all the bound params for the list products v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters list_products_v1
type ListProductsV1Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Brands []string
	/*
	  In: query
	*/
	Categories []string
	/*
	  In: query
	*/
	Colors []string
	/*
	  Required: true
	  In: query
	*/
	Gender string
	/*
	  In: query
	*/
	Labelsin []string
	/*
	  In: query
	*/
	MaxPrice *int32
	/*
	  In: query
	*/
	MinPrice *int32
	/*
	  Maximum: 100
	  In: query
	  Default: 60
	*/
	PageSize *int32
	/*
	  In: query
	*/
	PageToken *string
	/*
	  In: query
	*/
	Sizes []string
	/*
	  In: query
	*/
	Sort *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListProductsV1Params() beforehand.
func (o *ListProductsV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qBrands, qhkBrands, _ := qs.GetOK("brands")
	if err := o.bindBrands(qBrands, qhkBrands, route.Formats); err != nil {
		res = append(res, err)
	}

	qCategories, qhkCategories, _ := qs.GetOK("categories")
	if err := o.bindCategories(qCategories, qhkCategories, route.Formats); err != nil {
		res = append(res, err)
	}

	qColors, qhkColors, _ := qs.GetOK("colors")
	if err := o.bindColors(qColors, qhkColors, route.Formats); err != nil {
		res = append(res, err)
	}

	qGender, qhkGender, _ := qs.GetOK("gender")
	if err := o.bindGender(qGender, qhkGender, route.Formats); err != nil {
		res = append(res, err)
	}

	qLabelsin, qhkLabelsin, _ := qs.GetOK("labelsin")
	if err := o.bindLabelsin(qLabelsin, qhkLabelsin, route.Formats); err != nil {
		res = append(res, err)
	}

	qMaxPrice, qhkMaxPrice, _ := qs.GetOK("max_price")
	if err := o.bindMaxPrice(qMaxPrice, qhkMaxPrice, route.Formats); err != nil {
		res = append(res, err)
	}

	qMinPrice, qhkMinPrice, _ := qs.GetOK("min_price")
	if err := o.bindMinPrice(qMinPrice, qhkMinPrice, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("page_size")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageToken, qhkPageToken, _ := qs.GetOK("page_token")
	if err := o.bindPageToken(qPageToken, qhkPageToken, route.Formats); err != nil {
		res = append(res, err)
	}

	qSizes, qhkSizes, _ := qs.GetOK("sizes")
	if err := o.bindSizes(qSizes, qhkSizes, route.Formats); err != nil {
		res = append(res, err)
	}

	qSort, qhkSort, _ := qs.GetOK("sort")
	if err := o.bindSort(qSort, qhkSort, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBrands binds and validates array parameter Brands from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *ListProductsV1Params) bindBrands(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvBrands string
	if len(rawData) > 0 {
		qvBrands = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	brandsIC := swag.SplitByFormat(qvBrands, "")
	if len(brandsIC) == 0 {
		return nil
	}

	var brandsIR []string
	for _, brandsIV := range brandsIC {
		brandsI := brandsIV

		brandsIR = append(brandsIR, brandsI)
	}

	o.Brands = brandsIR

	return nil
}

// bindCategories binds and validates array parameter Categories from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *ListProductsV1Params) bindCategories(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvCategories string
	if len(rawData) > 0 {
		qvCategories = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	categoriesIC := swag.SplitByFormat(qvCategories, "")
	if len(categoriesIC) == 0 {
		return nil
	}

	var categoriesIR []string
	for _, categoriesIV := range categoriesIC {
		categoriesI := categoriesIV

		categoriesIR = append(categoriesIR, categoriesI)
	}

	o.Categories = categoriesIR

	return nil
}

// bindColors binds and validates array parameter Colors from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *ListProductsV1Params) bindColors(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvColors string
	if len(rawData) > 0 {
		qvColors = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	colorsIC := swag.SplitByFormat(qvColors, "")
	if len(colorsIC) == 0 {
		return nil
	}

	var colorsIR []string
	for _, colorsIV := range colorsIC {
		colorsI := colorsIV

		colorsIR = append(colorsIR, colorsI)
	}

	o.Colors = colorsIR

	return nil
}

// bindGender binds and validates parameter Gender from query.
func (o *ListProductsV1Params) bindGender(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("gender", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("gender", "query", raw); err != nil {
		return err
	}
	o.Gender = raw

	if err := o.validateGender(formats); err != nil {
		return err
	}

	return nil
}

// validateGender carries on validations for parameter Gender
func (o *ListProductsV1Params) validateGender(formats strfmt.Registry) error {

	if err := validate.EnumCase("gender", "query", o.Gender, []interface{}{"men", "women"}, true); err != nil {
		return err
	}

	return nil
}

// bindLabelsin binds and validates array parameter Labelsin from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *ListProductsV1Params) bindLabelsin(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvLabelsin string
	if len(rawData) > 0 {
		qvLabelsin = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	labelsinIC := swag.SplitByFormat(qvLabelsin, "")
	if len(labelsinIC) == 0 {
		return nil
	}

	var labelsinIR []string
	for _, labelsinIV := range labelsinIC {
		labelsinI := labelsinIV

		labelsinIR = append(labelsinIR, labelsinI)
	}

	o.Labelsin = labelsinIR

	return nil
}

// bindMaxPrice binds and validates parameter MaxPrice from query.
func (o *ListProductsV1Params) bindMaxPrice(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("max_price", "query", "int32", raw)
	}
	o.MaxPrice = &value

	return nil
}

// bindMinPrice binds and validates parameter MinPrice from query.
func (o *ListProductsV1Params) bindMinPrice(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("min_price", "query", "int32", raw)
	}
	o.MinPrice = &value

	return nil
}

// bindPageSize binds and validates parameter PageSize from query.
func (o *ListProductsV1Params) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewListProductsV1Params()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("page_size", "query", "int32", raw)
	}
	o.PageSize = &value

	if err := o.validatePageSize(formats); err != nil {
		return err
	}

	return nil
}

// validatePageSize carries on validations for parameter PageSize
func (o *ListProductsV1Params) validatePageSize(formats strfmt.Registry) error {

	if err := validate.MaximumInt("page_size", "query", int64(*o.PageSize), 100, false); err != nil {
		return err
	}

	return nil
}

// bindPageToken binds and validates parameter PageToken from query.
func (o *ListProductsV1Params) bindPageToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.PageToken = &raw

	return nil
}

// bindSizes binds and validates array parameter Sizes from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *ListProductsV1Params) bindSizes(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvSizes string
	if len(rawData) > 0 {
		qvSizes = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	sizesIC := swag.SplitByFormat(qvSizes, "")
	if len(sizesIC) == 0 {
		return nil
	}

	var sizesIR []string
	for _, sizesIV := range sizesIC {
		sizesI := sizesIV

		sizesIR = append(sizesIR, sizesI)
	}

	o.Sizes = sizesIR

	return nil
}

// bindSort binds and validates parameter Sort from query.
func (o *ListProductsV1Params) bindSort(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Sort = &raw

	if err := o.validateSort(formats); err != nil {
		return err
	}

	return nil
}

// validateSort carries on validations for parameter Sort
func (o *ListProductsV1Params) validateSort(formats strfmt.Registry) error {

	if err := validate.EnumCase("sort", "query", *o.Sort, []interface{}{"price_asc", "price_desc"}, true); err != nil {
		return err
	}

	return nil
}
