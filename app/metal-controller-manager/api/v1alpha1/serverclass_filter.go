// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package v1alpha1

import "sort"

// FilterAcceptedServers returns a new slice of Servers that are accepted and qualify.
//
// Returned Servers are always sorted by name for stable results.
func FilterAcceptedServers(servers []Server, q Qualifiers) []Server {
	servers = filterAccepted(servers)
	servers = filterCPU(servers, q.CPU)
	servers = filterSysInfo(servers, q.SystemInformation)
	servers = filterLabels(servers, q.LabelSelectors)

	sort.Slice(servers, func(i, j int) bool { return servers[i].Name < servers[j].Name })

	return servers
}

func filterAccepted(servers []Server) []Server {
	res := make([]Server, 0, len(servers))

	for _, server := range servers {
		if server.Spec.Accepted {
			res = append(res, server)
		}
	}

	return res
}

func filterCPU(servers []Server, filters []CPUInformation) []Server {
	if len(filters) == 0 {
		return servers
	}

	res := make([]Server, 0, len(servers))

	for _, server := range servers {
		var match bool

		for _, cpu := range filters {
			if server.Spec.CPU != nil && cpu.PartialEqual(server.Spec.CPU) {
				match = true
				break
			}
		}

		if match {
			res = append(res, server)
		}
	}

	return res
}

func filterSysInfo(servers []Server, filters []SystemInformation) []Server {
	if len(filters) == 0 {
		return servers
	}

	res := make([]Server, 0, len(servers))

	for _, server := range servers {
		var match bool

		for _, sysInfo := range filters {
			if server.Spec.SystemInformation != nil && sysInfo.PartialEqual(server.Spec.SystemInformation) {
				match = true
				break
			}
		}

		if match {
			res = append(res, server)
		}
	}

	return res
}

func filterLabels(servers []Server, filters []map[string]string) []Server {
	if len(filters) == 0 {
		return servers
	}

	res := make([]Server, 0, len(servers))

	for _, server := range servers {
		var match bool

		for _, label := range filters {
			for labelKey, labelVal := range label {
				if val, ok := server.ObjectMeta.Labels[labelKey]; ok {
					if labelVal == val {
						match = true
						break
					}
				}
			}
		}

		if match {
			res = append(res, server)
		}
	}

	return res
}
