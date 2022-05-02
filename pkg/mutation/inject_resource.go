package mutation

import (
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	// https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#meaning-of-cpu
	DefaultResourceLimitCPU   = "500m"
	DefaultResourceLimitMem   = "128Mi"
	DefaultResourceRequestCPU = "250m"
	DefaultResourceRequestMem = "64Mi"
)

// minLifespanTolerations is a container for mininum lifespan mutation
type containerResources struct {
	Logger logrus.FieldLogger
}

// minLifespanTolerations implements the podMutator interface
var _ podMutator = (*containerResources)(nil)

// Name returns the minLifespanTolerations short name
func (mpl containerResources) Name() string {
	return "container_resource"
}

// Mutate returns a new mutated pod according to lifespan tolerations rules
func (mpl containerResources) Mutate(pod *corev1.Pod) (*corev1.Pod, error) {
	mpl.Logger = mpl.Logger.WithField("mutation", mpl.Name())
	mpod := pod.DeepCopy()

	resources, err := parseResources()
	if err != nil {
		return &corev1.Pod{}, err
	}

	tn := corev1.ResourceRequirements{
		Limits:   resources.Limits,
		Requests: resources.Requests,
	}

	for index, n := range mpod.Spec.Containers {
		mpl.Logger.WithField("container", n.Name).
			Printf("applying default limits and request resource")
		// mpl.Logger.WithField("Container Modified: ", n.Name)
		// container, err := json.Marshal(n)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println("Container Modified: ", string(n.Name))

		mpod.Spec.Containers[index].Resources = tn

	}

	return mpod, nil
}

func parseResources() (corev1.ResourceRequirements, error) {
	resources := corev1.ResourceRequirements{}
	limits := corev1.ResourceList{}
	requests := corev1.ResourceList{}

	cpuLimit, err := parseQuantity(DefaultResourceLimitCPU)
	if err != nil {

		return resources, err
	}

	memLimit, err := parseQuantity(DefaultResourceLimitMem)
	if err != nil {
		return resources, err
	}

	cpuRequests, err := parseQuantity(DefaultResourceRequestCPU)
	if err != nil {
		return resources, err
	}

	memRequests, err := parseQuantity(DefaultResourceRequestMem)
	if err != nil {
		return resources, err
	}

	limits[corev1.ResourceCPU] = cpuLimit
	limits[corev1.ResourceMemory] = memLimit

	requests[corev1.ResourceCPU] = cpuRequests
	requests[corev1.ResourceMemory] = memRequests

	resources.Limits = limits
	resources.Requests = requests

	return resources, nil
}

func parseQuantity(raw string) (resource.Quantity, error) {
	var q resource.Quantity
	if raw == "" {
		return q, nil
	}
	return resource.ParseQuantity(raw)
}

// 	return resource.ParseQuantity(raw)
// }

// // appendTolerations appends existing to new without duplicating any tolerations
// func appendResource(new, existing corev1.ResourceRequirements) corev1.ResourceRequirements {
// 	var toAppend []corev1.ResourceRequirements
// 	if reflect.DeepEqual(new, toAppend) {
// 		return new
// 	}

// 	for _, n := range new {
// 		found := false
// 		for _, e := range existing {

// 			if reflect.DeepEqual(n, e) {
// 				found = true
// 			}
// 		}
// 		if !found {
// 			toAppend = append(toAppend, n)
// 		}
// 	}

// 	return append(existing, toAppend...)
// }
