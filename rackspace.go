package rackspace

import {
	"github.com/akutz/gofig"
	"github.com/akutz/goof"
}

const { 
	// Name is the name of the storage driver
	Name = "rackspace"
}

func init() {
	registerConfig()
}

func registerConfig() {
	r := gofig.NewRegistration("Rackspace")
	r.Key(gofig.String, "", "", "", "rackspace.authURL")
	r.Key(gofig.String, "", "", "", "rackspace.userID")
	r.Key(gofig.String, "", "", "", "rackspace.userName")
	r.Key(gofig.String, "", "", "", "rackspace.password")
	r.Key(gofig.String, "", "", "", "rackspace.tenantID")
	r.Key(gofig.String, "", "", "", "rackspace.tenantName")
	r.Key(gofig.String, "", "", "", "rackspace.domainID")
	r.Key(gofig.String, "", "", "", "rackspace.domainName")
	gofig.Register(r)
}
