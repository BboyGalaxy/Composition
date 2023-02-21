package invoice

import (
	"github.com/BboyGalaxy/Composition/pkg/customer"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
}
