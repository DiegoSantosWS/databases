package elasticsearch

import (
	"context"

	"github.com/DiegoSantosWS/databases/elasticsearch/es01db"
)

func Insert(ctx context.Context) {
	clients := es01db.Client{
		ID:    "2",
		Name:  "Diego Dos Santos",
		Email: "santos.diedo@gmail.com",
		Address: es01db.Address{
			Address: "Rua Osvaldo Lima e Silva",
			Number:  "1190",
		},

		Phone: es01db.Phone{
			NumType: "Celphone",
			Number:  "319984008582",
		},
	}
	es01db.LoadObjectID(ctx, "test-diego", clients)
}

func Update(ctx context.Context) {

	clients := es01db.Client{
		ID:    "1",
		Name:  "Naiane Cristina Santos",
		Email: "naiane@gmail.com",
		Address: es01db.Address{
			Address: "Rua Osvaldo Lima e Silva",
			Number:  "1190",
		},
		Phone: es01db.Phone{
			NumType: "Celphone",
			Number:  "11988897721",
		},
	}
	es01db.Update(ctx, "test-diego", "1", clients)
}
