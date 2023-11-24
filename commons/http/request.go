package rest

import "errors"

// IDRequest
type IDRequest struct {
	ID int64 `json:"id"`
}

func (req *IDRequest) Validate() error {
	if req.ID == 0 {
		return errors.New("must have id")
	}
	return nil
}

type IDRequestUint struct {
	ID uint64 `json:"id"`
}

func (req *IDRequestUint) Validate() error {
	if req.ID == 0 {
		return errors.New("must have id")
	}
	return nil
}

type IDRequestString struct {
	ID string `json:"id"`
}

func (req *IDRequestString) Validate() error {
	if req.ID == "" {
		return errors.New("must have id")
	}
	return nil
}
