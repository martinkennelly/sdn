package metrics

// contains all controller/master metric definitions and registration

import "github.com/prometheus/client_golang/prometheus"

const (
	metricSDNNamespace           = "sdn"
	metricSDNSubsystemController = "controller"
)

var metricEgressIPCount = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: metricSDNNamespace,
	Subsystem: metricSDNSubsystemController,
	Name:      "num_egress_ips",
	Help:      "The number of defined egress IP addresses",
})

// represents kind EgressNetworkPolicy from API version network.openshift.io/v1
var metricEgressFirewallRuleCount = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: metricSDNNamespace,
	Subsystem: metricSDNSubsystemController,
	Name:      "num_egress_firewall_rules",
	Help:      "The number of egress firewall rules defined"},
)

var metricEgressNetworkPolicyCount = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: metricSDNNamespace,
	Subsystem: metricSDNSubsystemController,
	Name: "num_egress_network_policies",
	Help: "The number of egress network policies",
})

var registry = prometheus.NewRegistry()

func register() {
	registry.MustRegister(metricEgressIPCount)
	registry.MustRegister(metricEgressFirewallRuleCount)
	registry.MustRegister(metricEgressNetworkPolicyCount)
}
