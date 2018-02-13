package operator

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	chaosv1alpha1 "github.com/spotahome/kooper/examples/pod-terminator-operator/apis/chaos/v1alpha1"
	"github.com/spotahome/kooper/examples/pod-terminator-operator/log"
	"github.com/spotahome/kooper/examples/pod-terminator-operator/operator/service"
)

// Handler  is the pod terminator handler that will handle the
// events received from kubernetes.
type handler struct {
	chaosService service.ChaosSyncer
	logger       log.Logger
}

// newHandler returns a new handler.
func newHandler(k8sCli kubernetes.Interface, logger log.Logger) *handler {
	return &handler{
		chaosService: service.NewChaos(k8sCli, logger),
		logger:       logger,
	}
}

// Add will ensure that the required pod terminator is running.
func (h *handler) Add(obj runtime.Object) error {
	pt, ok := obj.(*chaosv1alpha1.PodTerminator)
	if !ok {
		return fmt.Errorf("%v is not a pod terminator object", obj.GetObjectKind())
	}

	return h.chaosService.EnsurePodTerminator(pt)
}

// Delete will ensure the reuited pod terminator is not running.
func (h *handler) Delete(name string) error {
	return h.chaosService.DeletePodTerminator(name)
}
