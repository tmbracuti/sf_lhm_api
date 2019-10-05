package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"AddPrimePhone",
		"POST",
		"/provisioning/v1/addprimephone",
		AddPrimePhone,
	},

	Route{
		"AddVoiceMail",
		"POST",
		"/provisioning/v1/addvoicemail",
		AddVoiceMail,
	},
	Route{
		"ModVoiceMail",
		"POST",
		"/provisioning/v1/modvoicemail",
		ModVoiceMail,
	},
	Route{
		"Index",
		"GET",
		"/",
		IndexHandler,
	},

	Route{
		"BOE",
		"GET",
		"/cmrest/GetBOELocationData",
		BoeHandler,
	},

	Route{
		"Authenticate",
		"POST",
		"/provisioning/v1/authenticate",
		Authenticate,
	},
	Route{
		"PatchGroups",
		"POST",
		"/provisioning/v1/patch",
		PatchIt,
	},

	Route{
		"Sponsored",
		"POST",
		"/api/x_lmlmc_ent_phone/sponsor_phone_post.do",
		Sponsored,
	},

	Route{
		"Voicemail",
		"POST",
		"/api/x_lmlmc_ent_phone/voicemail_request_post.do",
		Voicemail,
	},

	Route{
		"Incident",
		"POST",
		"/api/lmlmc/lm_create_incident",
		ReportIncident,
	},

	//api/x_lmlmc_ent_phone/prim_phone_post.do
	Route{
		"PrimePhoneRes",
		"POST",
		"/api/x_lmlmc_ent_phone/prim_phone_post.do",
		PrimePhoneResult,
	},

	Route{
		"AXL",
		"POST",
		"/axl/",
		ShowAxl,
	},
}
