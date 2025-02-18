package main_test

import (
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Fila de Produção - Pedidos", func() {

	Context("Usuário altera o status de um pedido de produção", func() {
		When("o pedido está na fila", func() {
			BeforeEach(func() {
				req, _ := http.NewRequest(http.MethodPost, "http://localhost:8000/production/order/send", strings.NewReader(`{"order_id": 99}`))
				req.Header.Set("Content-Type", "application/json")
				_, err := http.DefaultClient.Do(req)
				assert.NoError(GinkgoT(), err)
			})

			It("o status do pedido deve ser alterado para novo status", func() {
				req, _ := http.NewRequest(http.MethodPut, "http://localhost:8000/production/order/99/status", strings.NewReader(`{"status": "EM_PREPARACAO"}`))
				req.Header.Set("Content-Type", "application/json")
				res, err := http.DefaultClient.Do(req)

				assert.NoError(GinkgoT(), err)
				assert.Equal(GinkgoT(), http.StatusOK, res.StatusCode)
			})
		})
	})
})
