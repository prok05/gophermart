package v1

import (
	"errors"
	"github.com/prok05/gophermart/internal/controller/http/request"
	"github.com/prok05/gophermart/internal/controller/http/response"
	"github.com/prok05/gophermart/internal/entity"
	"net/http"
)

// @Summary		Register new user
// @Description	Creates a new user with given login and password
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			payload	body		request.RegisterUser	true	"User credentials"
// @Success		200		{object}	nil
// @Failure		409		{object}	response.Error	"Resource already exists"
// @Failure		400		{object}	response.Error	"Invalid request body"
// @Failure		500		{object}	response.Error	"Server encountered a problem"
// @Router			/user/register [post]
func (h *V1) register(w http.ResponseWriter, r *http.Request) {
	var payload request.RegisterUser

	// parsing request
	if err := readJSON(w, r, &payload); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	// validating request
	if err := h.v.Struct(payload); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	// register use case
	if err := h.u.Register(r.Context(), payload); err != nil {
		switch {
		case errors.Is(err, entity.ErrDuplicateLogin):
			h.conflictResponse(w, r, err)
		default:
			h.internalServerError(w, r, err)
		}
		return
	}

	if err := h.jsonResponse(w, http.StatusOK, nil); err != nil {
		h.internalServerError(w, r, err)
	}
}

// @Summary		Login user
// @Description	Authenticates user with given login and password
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			payload	body		request.LoginUser	true	"User credentials"
// @Success		200		{object}	response.LoginResponse
// @Failure		400		{object}	response.Error	"Invalid request body"
// @Failure		401		{object}	response.Error	"Wrong login/password"
// @Failure		500		{object}	response.Error	"Server encountered a problem"
// @Router			/user/login [post]
func (h *V1) login(w http.ResponseWriter, r *http.Request) {
	var payload request.LoginUser

	// parsing request
	if err := readJSON(w, r, &payload); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	// validating request
	if err := h.v.Struct(payload); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	token, err := h.u.Login(r.Context(), payload)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			h.unauthorizedResponse(w, r, err)
		case errors.Is(err, entity.ErrWrongLoginOrPassword):
			h.unauthorizedResponse(w, r, err)
		default:
			h.internalServerError(w, r, err)
		}
		return
	}

	if err := writeJSON(w, http.StatusOK, response.LoginResponse{Token: token}); err != nil {
		h.internalServerError(w, r, err)
	}
}

// @Summary		Load order
// @Description	Load order for future processing
// @Tags			orders
// @Accept			plain
// @Produce		json
// @Param			orderNumber	body		string	true	"Order number"
// @Success		202			{object}	nil
// @Failure		200			{object}	response.Error	"Order was already loaded"
// @Failure		400			{object}	response.Error	"Invalid request body"
// @Failure		401			{object}	response.Error	"User is not authenticated"
// @Failure		409			{object}	response.Error	"Order was loaded by another user"
// @Failure		422			{object}	response.Error	"Wrong order number format"
// @Failure		500			{object}	response.Error	"Server encountered a problem"
// @Router			/user/orders [post]
// @Security		AuthToken
func (h *V1) loadOrder(w http.ResponseWriter, r *http.Request) {
	orderNumber, err := readPlain(r)
	if err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	if err := h.u.LoadOrder(r.Context(), orderNumber); err != nil {
		switch {
		case errors.Is(err, entity.ErrOrderAlreadyLoaded):
			h.jsonResponse(w, http.StatusOK, nil)
		case errors.Is(err, entity.ErrOrderLoadedByAnotherUser):
			h.conflictResponse(w, r, err)
		case errors.Is(err, entity.ErrInvalidOrderNumber):
			h.unprocessableEntityResponse(w, r, err)
		default:
			h.internalServerError(w, r, err)
		}
		return
	}

	if err := h.jsonResponse(w, http.StatusAccepted, nil); err != nil {
		h.internalServerError(w, r, err)
	}
}

// @Summary		Get user orders
// @Description	Get user orders
// @Tags			orders
// @Produce		json
// @Success		200	{object}	[]entity.UserOrder
// @Failure		204	{object}	response.Error	"Data is empty"
// @Failure		401	{object}	response.Error	"User is not authenticated"
// @Failure		500	{object}	response.Error	"Server encountered a problem"
// @Router			/user/orders [get]
// @Security		AuthToken
func (h *V1) getUserOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.u.GetOrders(r.Context())
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNoContent):
			h.noContentResponse(w, r, err)
		default:
			h.internalServerError(w, r, err)
		}
		return
	}

	if err := h.jsonResponse(w, http.StatusOK, orders); err != nil {
		h.internalServerError(w, r, err)
	}
}

// @Summary		User balance
// @Description	Retrieve user balance
// @Tags			user
// @Produce		json
// @Success		200	{object}	entity.UserBalance
// @Failure		401	{object}	response.Error	"User is not authenticated"
// @Failure		500	{object}	response.Error	"Server encountered a problem"
// @Router			/user/balance [get]
// @Security		AuthToken
func (h *V1) getUserBalance(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(entity.ContextUserID).(string)

	balance, err := h.u.GetBalance(r.Context(), userID)
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}

	if err := h.jsonResponse(w, http.StatusOK, balance); err != nil {
		h.internalServerError(w, r, err)
	}
}

// @Summary		Withdraw balance
// @Description	Request for withdrawing user balance
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			payload	body		request.WithdrawBalance	true	"Order number and amount"
// @Success		200		{object}	nil
// @Failure		401		{object}	response.Error	"User is not authenticated"
// @Failure		402		{object}	response.Error	"Balance is not enough"
// @Failure		422		{object}	response.Error	"Wrong order number format"
// @Failure		500		{object}	response.Error	"Server encountered a problem"
// @Router			/user/balance/withdraw [post]
// @Security		AuthToken
func (h *V1) withdrawUserBalance(w http.ResponseWriter, r *http.Request) {
	var payload request.WithdrawBalance

	// parsing request
	if err := readJSON(w, r, &payload); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	// validating request
	if err := h.v.Struct(payload); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	if err := h.u.WithdrawBalance(r.Context(), payload); err != nil {
		switch {
		case errors.Is(err, entity.ErrInvalidOrderNumber):
			h.unprocessableEntityResponse(w, r, err)
		case errors.Is(err, entity.ErrNotEnoughBalance):
			h.paymentRequiredResponse(w, r, err)
		default:
			h.internalServerError(w, r, err)
		}
		return
	}

	if err := h.jsonResponse(w, http.StatusOK, nil); err != nil {
		h.internalServerError(w, r, err)
	}
}

// @Summary		User withdrawals
// @Description	Gets user withdrawals
// @Tags			user
// @Produce		json
// @Success		200	{object}	[]entity.UserWithdrawal
// @Failure		204	{object}	response.Error	"Data is empty"
// @Failure		401	{object}	response.Error	"User is not authenticated"
// @Failure		500	{object}	response.Error	"Server encountered a problem"
// @Router			/user/withdrawals [get]
// @Security		AuthToken
func (h *V1) getUserWithdrawals(w http.ResponseWriter, r *http.Request) {
	withdrawals, err := h.u.GetWithdrawals(r.Context())
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNoContent):
			h.noContentResponse(w, r, err)
		default:
			h.internalServerError(w, r, err)
		}
		return
	}

	if err := h.jsonResponse(w, http.StatusOK, withdrawals); err != nil {
		h.internalServerError(w, r, err)
	}
}
