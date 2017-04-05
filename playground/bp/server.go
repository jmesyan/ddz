package bp

import (
	"io"
	"net/http"
)

type Handler struct {
	routerMap map[string]func(http.ResponseWriter, *http.Request)
}

func (h *Handler) RegisterRouter(path string, handler func(http.ResponseWriter, *http.Request)) {
	if h.routerMap == nil {
		h.routerMap = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	h.routerMap[path] = handler
}

func (h *Handler) RemoveRouter(path string) {
	if h.routerMap == nil {
		return
	}
	delete(h.routerMap, path)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if h, ok := h.routerMap[path]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "No router for:"+path)
}
