// Package search is the Search togo plugin: Bilingual Postgres full-text site search: to_tsvector('simple') + prefix tsquery for match-as-you-type + ILIKE fallback, ranked across registered content tables.
//
// It self-registers a provider on blank-import and mounts its routes onto the
// kernel. The concrete implementation is ported from the fadymondy.com app under
// internal/server — this scaffold wires the provider + a health route.
package search

import (
	"net/http"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("search", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/search/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"plugin":"search","status":"ok"}`))
		})
		return nil
	})
}
