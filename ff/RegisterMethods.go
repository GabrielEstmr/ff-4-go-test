package ff

import ff_resources "github.com/GabrielEstmr/ff-4-go-test/ff/resources"

type RegisterMethods interface {
	RegisterFeatures(features ff_resources.Features) error
}
