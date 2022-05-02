package mutation

import (
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
)

// containerDebug is a container for mininum lifespan mutation
type containerDebug struct {
	Logger logrus.FieldLogger
}

// containerDebug implements the podMutator interface
var _ podMutator = (*containerDebug)(nil)

// Name returns the containerDebug short name
func (mpl containerDebug) Name() string {
	return "container_debug"
}

// Mutate returns a new mutated pod according to lifespan tolerations rules
func (mpl containerDebug) Mutate(pod *corev1.Pod) (*corev1.Pod, error) {
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
	container_def := corev1.Container{
		Name:      "debug",
		Image:     "praqma/network-multitool",
		Resources: tn,
		Env: []corev1.EnvVar{{
			Name:  "HTTP_PORT",
			Value: "8080",
		},
		},
	}
	mpl.Logger.WithField("Adding Debug Container", mpod.Name).
		Printf("adding default pod")
	mpod.Spec.Containers = append(mpod.Spec.Containers, container_def)

	return mpod, nil
}
