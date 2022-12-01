package flights

import "context"

type Flights interface {
	Update(ctx context.Context) error
}
