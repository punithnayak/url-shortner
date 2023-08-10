package routes

import (

)
type request struct{
	URL			string			`json:"url"`
	CustomShort	string			`json:"short"`
	Expiry		time.Duration	`json:"expiry"`
}

type response struct{
	URL				string			`json:"url"`
	CustomShort		string			`json:"short"`
	Expiry			time.Duration	`json:"expiry"`
	XRateRemaining	int				`json:"rate_limit"`
	XRateLimitRest	time.Duration	`json:"rate_limit_reset"`
}