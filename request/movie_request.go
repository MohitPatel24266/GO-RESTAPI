package request
import(
	"github.com/go-playground/validator/v10"

)
var validate = validator.New()

type Request struct{
	ID    int     `json:"id"`
	Name  string  `json:"name" validate:"required,min=1,max=100"`
	Genre string  `json:"genre" validate:"required,min=1,max=50"`
	Price float64 `json:"price" validate:"required,gt=0"`
}

func (r *Request) Validate() error {
	return validate.Struct(r)
}