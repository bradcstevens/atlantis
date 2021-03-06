package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsPullRequest() models.PullRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.PullRequest))(nil)).Elem()))
	var nullValue models.PullRequest
	return nullValue
}

func EqModelsPullRequest(value models.PullRequest) models.PullRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.PullRequest
	return nullValue
}
