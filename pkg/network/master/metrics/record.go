package metrics

// contains all controller/master metric data updates

// RecordEgressFirewallRuleCount records the number of Egress firewall rules.
// Represents the sum of all egress rules for kind EgressNetworkPolicy.
func RecordEgressFirewallRuleCount(count float64) {
	metricEgressFirewallRuleCount.Set(count)
}

// RecordEgressIPCount records the number of Egress IPs.
// This may include multiple Egress IPs for kind EgressIP.
func RecordEgressIPCount(count float64) {
	metricEgressIPCount.Set(count)
}

// RecordEgressNetworkPolicyCount records the number of EgressNetworkPolicy
func RecordEgressNetworkPolicyCount(count float64) {
	metricEgressNetworkPolicyCount.Set(count)
}
