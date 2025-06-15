package v1

import "net/http"

func (h *V1) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	h.l.Error("internal error", "method", r.Method, "path", r.URL.Path, "error", err)
	writeJSONError(w, http.StatusInternalServerError, "server encountered a problem")
}

func (h *V1) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.l.Warn("bad request error", "method", r.Method, "path", r.URL.Path, "error", err)
	writeJSONError(w, http.StatusBadRequest, "invalid request body")
}

func (h *V1) unprocessableEntityResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.l.Error("unprocessable entity error", "method", r.Method, "path", r.URL.Path, "error", err)
	writeJSONError(w, http.StatusUnprocessableEntity, "unprocessable entity")
}

func (h *V1) paymentRequiredResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.l.Error("payment required error", "method", r.Method, "path", r.URL.Path, "error", err)
	writeJSONError(w, http.StatusPaymentRequired, "payment required")
}

func (h *V1) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.l.Warn("not found error", "method", r.Method, "path", r.URL.Path, "error", err)
	writeJSONError(w, http.StatusNotFound, "not found")
}

func (h *V1) noContentResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.l.Warn("no content error", "method", r.Method, "path", r.URL.Path, "error", err)
	writeJSONError(w, http.StatusNoContent, "no content")
}

func (h *V1) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.l.Error("conflict error", "method", r.Method, "path", r.URL.Path, "error", err)
	writeJSONError(w, http.StatusConflict, "resource already exists")
}

func (h *V1) unauthorizedResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.l.Error("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err)
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}
