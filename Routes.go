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
	//mock of Avaya APCS snow callback
	Route{
		"AvayaSNOWCallback",
		"POST",
		"/api/asiam/starfish_response",
		AvayaSNOWCallback,
	},

	Route{
		"AddPrimePhone",
		"POST",
		"/provisioning/v1/addprimephone",
		AddPrimePhone,
	},
	//GetTelephoneNumbersByGUID
	Route{
		"GetTelephoneNumbersByGUID",
		"POST",
		"/tips/GetTelephoneNumbersByGUID",
		GetTelephoneNumbersByGUID,
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
		"ChangePhone",
		"POST",
		"/api/x_lmlmc_ent_phone/change_phone_post",
		ChangePhoneResult,
	},

	Route{
		"DellBulkReport",
		"POST",
		"/api/dusal/starfish_inbound/bulk_load",
		DellBulkHandler,
	},

	Route{
		"DellBulkReport",
		"POST",
		"/api/dusal/starfish_inbound/bulk_load_summary",
		DellBulkSummaryHandler,
	},

	// - /api/sn_sc/servicecatalog/items/d92de7a4db120cd035453efd7c96191f/order_now
	Route{
		"TermCB",
		"POST",
		"/api/sn_sc/servicecatalog/items/d92de7a4db120cd035453efd7c96191f/order_now",
		TerminationCB,
	},

	Route{
		"NewHireCB",
		"POST",
		"/api/sn_sc/servicecatalog/items/19682786db3408d0f101307f7c96191e/order_now",
		NewHireCB,
	},

	Route{
		"AXL",
		"POST",
		"/axl/",
		ShowAxl,
	},

	Route{
		"MegaGetEmployees",
		"GET",
		"/tenants/{tenantId}/employees",
		GetMegaEmployees,
	},

	Route{
		"MegaAddEmployee",
		"POST",
		"/tenants/{tenantId}/employees",
		AddMegaEmployee,
	},
}
