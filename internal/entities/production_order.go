package entities

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	RECEIVED_STATUS       = "RECEBIDO"
	IN_PREPARATION_STATUS = "EM_PREPARACAO"
	DONE_STATUS           = "PRONTO"
	FINISHED_STATUS       = "FINALIZADO"
)

type ProductionOrder struct {
	OrderId uint32 `dynamo:"ID,hash"`
	Status  string
}

func (o *ProductionOrder) Validate() error {
	return validation.ValidateStruct(
		o,
		validation.Field(
			&o.OrderId,
			validation.Required,
		),
		validation.Field(
			&o.Status,
			validation.Required,
			validation.In(
				RECEIVED_STATUS,
				IN_PREPARATION_STATUS,
				DONE_STATUS,
				FINISHED_STATUS,
			).Error(
				fmt.Sprintf(
					"must be between %s, %s, %s or %s",
					RECEIVED_STATUS,
					IN_PREPARATION_STATUS,
					DONE_STATUS,
					FINISHED_STATUS,
				),
			),
		),
	)
}
