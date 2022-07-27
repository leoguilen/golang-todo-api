package abstractions

import (
	"context"
)

type IHealthcheckRepository interface {
	Check(ctx context.Context) error
}
