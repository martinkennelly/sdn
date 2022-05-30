package unidler

import (
	"net"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	eventsv1 "k8s.io/client-go/tools/events"
	"k8s.io/kubernetes/pkg/util/iptables"
	utilexec "k8s.io/utils/exec"

	unidlingapi "github.com/openshift/api/unidling/v1alpha1"
)

// NewUnidlerProxier creates a new Proxier for the given LoadBalancer and address which fires off
// unidling signals connections and traffic.  It is intended to be used as one half of a HybridProxier.
func NewUnidlerProxier(listenIP net.IP, iptables iptables.Interface, exec utilexec.Interface, pr utilnet.PortRange,
	syncPeriod, minSyncPeriod, udpIdleTimeout time.Duration, nodePortAddresses []string,
	eventRecorder eventsv1.EventRecorder) (*Proxier, error) {
	signaler := &NeedPodsSignaler{eventRecorder}
	newFunc := func(protocol v1.Protocol, ip net.IP, port int) (ProxySocket, error) {
		return newUnidlerSocket(protocol, ip, port, signaler)
	}
	return NewProxier(listenIP, iptables, exec, pr, syncPeriod, minSyncPeriod, udpIdleTimeout, nodePortAddresses, newFunc)
}

type NeedPodsSignaler struct {
	recorder eventsv1.EventRecorder
}

// NeedPods signals that endpoint addresses are needed in order to
// service a traffic coming to the given service and port
func (sig *NeedPodsSignaler) NeedPods(serviceName types.NamespacedName, port string) error {
	// TODO: we need to fake this since upstream removed our handle to the ObjectReference
	// This *should* be sufficient for the unidling controller
	serviceRef := v1.ObjectReference{
		Kind:      "Service",
		Namespace: serviceName.Namespace,
		Name:      serviceName.Name,
	}

	// HACK: make the message different to prevent event aggregation
	sig.recorder.Eventf(&serviceRef, nil, v1.EventTypeNormal, unidlingapi.NeedPodsReason, "Unidling",
		"The service-port %s:%s needs pods.", serviceRef.Name, port)

	return nil
}
