package phosphor

import (
	"fmt"

	"golang.org/x/net/context"
)

func phosphorFromContext(ctx context.Context) (*Phosphor, error) {

	if p, ok := ctx.Value("phosphor").(*Phosphor); ok {
		return p, nil
	}

	return nil, fmt.Errorf("Couldn't retrieve Phosphor from Context")

}
