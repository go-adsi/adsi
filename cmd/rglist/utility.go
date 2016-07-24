package main

import "strings"

func dfsrRootPath(domain string) string {
	cn := makeDN("cn", "DFSR-GlobalSettings", "System")
	dc := makeDN("dc", strings.Split(domain, ".")...)
	dn := combineDN(cn, dc)
	protocol := "LDAP"
	return protocol + "://" + dn
}

func makeDN(attribute string, components ...string) string {
	if attribute != "" {
		for i := 0; i < len(components); i++ {
			components[i] = attribute + "=" + components[i]
		}
	}
	return strings.Join(components, ",")
}

func combineDN(components ...string) string {
	return strings.Join(components, ",")
}
