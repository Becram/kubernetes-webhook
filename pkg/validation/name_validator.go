package validation

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
)

// nameValidator is a container for validating the name of pods
type nameValidator struct {
	Logger logrus.FieldLogger
}

// nameValidator implements the podValidator interface
var _ podValidator = (*nameValidator)(nil)

// Name returns the name of nameValidator
func (n nameValidator) Name() string {
	return "image_validator"
}

// Validate inspects the name of a given pod and returns validation.
// The returned validation is only valid if the pod name does not contain some
// bad string.
func (n nameValidator) Validate(pod *corev1.Pod) (validation, error) {

	// fmt.Printf("limit %s\n", pod.Spec.Containers[0].Resources)
	// fmt.Printf("limit type: %s\n", reflect.TypeOf(pod.Spec.Containers[0].Resources))

	// return validation{Valid: true, Reason: "Limit"}, nil
	// if len(pod.Spec.Containers[0].Resources.Limits) = 0 {}

	// }

	badString := "latest"

	if strings.Contains(pod.Spec.Containers[0].Image, badString) {
		v := validation{
			Valid:  false,
			Reason: fmt.Sprintf("pod image tag contains %q", badString),
		}
		return v, nil
	}

	return validation{Valid: true, Reason: "valid image tag"}, nil
}
