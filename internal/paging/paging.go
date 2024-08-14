package paging

type TokenParams struct {
	PageToken string
	PageSize  int32
}

type OffsetParams struct {
	Offset int32
	Limit  int32
}

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type CursorParams[T integer] struct {
	LastID   T
	PageSize int32
}
