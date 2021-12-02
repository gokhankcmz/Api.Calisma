package ErrorModels

import (
	"Api.Calisma/src/CustomerService/Constants"
	"fmt"
	"net/http"
)

type Error struct {
	PublicError PublicError
	StatusCode  int
	DetailCode  int
	Args        interface{}
}

type PublicError struct {
	From        string
	Description string
	Detail      string
}

func New(from string, statusCode int, detailCode int, description string, args interface{}) *Error {
	return &Error{
		PublicError: PublicError{
			From:        from,
			Description: description,
		},
		StatusCode: statusCode,
		DetailCode: detailCode,
		Args:       args,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("From: %s, Description: %s, DetailCode: %d,  Args: %v", e.PublicError.From, e.PublicError.Description, e.DetailCode, e.Args)
}

func (e *Error) SetArgs(args interface{}) *Error {
	return &Error{
		PublicError: PublicError{
			From:        e.PublicError.From,
			Description: e.PublicError.Description,
			Detail:      e.PublicError.Detail,
		},
		StatusCode: e.StatusCode,
		DetailCode: e.DetailCode,
		Args:       args,
	}
}

func (e *Error) SetDetailCode(detailCode int) *Error {
	return &Error{
		PublicError: PublicError{
			From:        e.PublicError.From,
			Description: e.PublicError.Description,
			Detail:      e.PublicError.Detail,
		},
		StatusCode: e.StatusCode,
		DetailCode: detailCode,
		Args:       e.Args,
	}
}
func (e *Error) SetPublicDetail(detail string) *Error {
	return &Error{
		PublicError: PublicError{
			From:        e.PublicError.From,
			Description: e.PublicError.Description,
			Detail:      detail,
		},
		StatusCode: e.StatusCode,
		DetailCode: e.DetailCode,
		Args:       e.Args,
	}
}

var (
	UnknownError        = New(Constants.ErrorFrom, http.StatusInternalServerError, 0, "An unknown error occured. ", nil)
	EntityNotFound      = New(Constants.ErrorFrom, http.StatusNoContent, 0, "Entity not found. ", nil)
	InvalidModel        = New(Constants.ErrorFrom, http.StatusBadRequest, 0, "Invalid model. ", nil)
	InvalidToken        = New(Constants.ErrorFrom, http.StatusUnauthorized, 0, "Invalid Token.", nil)
	InvalidCredentials  = New(Constants.ErrorFrom, http.StatusUnauthorized, 0, "Invalid Token.", nil)
	UnauthorizedRequest = New(Constants.ErrorFrom, http.StatusForbidden, 0, "You are not authorized to perform this query.", nil)
	CustomerNotFound    = New(Constants.ErrorFrom, http.StatusNotFound, 0, "Customer not found.", nil)
	RedirectionFailed   = New(Constants.ErrorFrom, http.StatusServiceUnavailable, 0, "Redirection failed", nil)
	/*DeleteManyFailed      = New(repoOp, "Delete many failed", 3001, http.StatusInternalServerError)
	DeleteOneFailed       = New(repoOp, "Delete one failed", 3002, http.StatusInternalServerError)
	FromBytesFailed       = New(repoOp, "From byte failed !!!", 3003, http.StatusInternalServerError)
	InsertManyFailed      = New(repoOp, "Insert many failed !!!", 3004, http.StatusInternalServerError)
	InsertOneFailed       = New(repoOp, "Insert one failed !!!", 3005, http.StatusInternalServerError)
	UpdateOneFailed       = New(Constants.ErrorFrom, "Update one Failed !!!", 3006, http.StatusInternalServerError)*/
	/*FindFailed            = New(repoOp, "FindByMerchantId Failed !!!", 3007, http.StatusInternalServerError)
	MongoCursorFailed     = New(repoOp, "Mongo cursor Failed !!!", 3008, http.StatusInternalServerError)
	MongoDecodeFailed     = New(repoOp, "Mongo decode Failed !!!", 3009, http.StatusInternalServerError)
	MongoBulkWriteFailed  = New(repoOp, "Mongo bulk write Failed !!!", 3010, http.StatusInternalServerError)
	OffsetGreaterThanZero = New(repoOp, "Offset value must be greater than zero", 3011, http.StatusBadRequest)
	LimitGreaterThanZero  = New(repoOp, "Limit value must be greater than zero", 3012, http.StatusBadRequest)
	CountDocumentsFailed  = New(repoOp, "Count Documents Failed !!!", 3013, http.StatusInternalServerError)
	UpdateManyFailed      = New(repoOp, "Update many Failed !!!", 3014, http.StatusInternalServerError)
	InvalidTime           = New(repoOp, "Given Date is in the future", 3015, http.StatusBadRequest)
	FindOneFailed         = New(repoOp, "FindOne failed !!!", 3016, http.StatusInternalServerError)*/
)
