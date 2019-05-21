package integration

// KudoEvent for event register
func KudoEvent(kudo *Kudo) ([]*Request, map[string]interface{}) {

	funKudoMap := map[string]interface{}{
		"PublishPulsa": kudo.PublishPulsa,
	}

	kudoRouting := []*Request{
		&Request{Code: "Refill-xl25", PublishCode: "xl25", Fn: "PublishPulsa"},
		&Request{Code: "Refill-xl50", PublishCode: "xl150", Fn: "PublishPulsa"},
		&Request{Code: "Refill-xl100", PublishCode: "xl100", Fn: "PublishPulsa"},
		&Request{Code: "Refill-xl150", PublishCode: "xl150", Fn: "PublishPulsa"},
		&Request{Code: "Refill-xl200", PublishCode: "xl200", Fn: "PublishPulsa"},
		&Request{Code: "Refill-indosat25", PublishCode: "indosat25", Fn: "PublishPulsa"},
		&Request{Code: "Refill-indosat50", PublishCode: "indosat150", Fn: "PublishPulsa"},
		&Request{Code: "Refill-indosat100", PublishCode: "indosat100", Fn: "PublishPulsa"},
		&Request{Code: "Refill-indosat150", PublishCode: "indosat150", Fn: "PublishPulsa"},
		&Request{Code: "Refill-indosat200", PublishCode: "indosat200", Fn: "PublishPulsa"},
		&Request{Code: "Refill-simpati25", PublishCode: "simpati25", Fn: "PublishPulsa"},
		&Request{Code: "Refill-simpati50", PublishCode: "simpati150", Fn: "PublishPulsa"},
		&Request{Code: "Refill-simpati100", PublishCode: "simpati100", Fn: "PublishPulsa"},
		&Request{Code: "Refill-simpati150", PublishCode: "simpati150", Fn: "PublishPulsa"},
		&Request{Code: "Refill-simpati200", PublishCode: "simpati200", Fn: "PublishPulsa"},
		&Request{Code: "Refill-tri25", PublishCode: "tri25", Fn: "PublishPulsa"},
		&Request{Code: "Refill-tri50", PublishCode: "tri150", Fn: "PublishPulsa"},
		&Request{Code: "Refill-tri100", PublishCode: "tri100", Fn: "PublishPulsa"},
		&Request{Code: "Refill-tri150", PublishCode: "tri150", Fn: "PublishPulsa"},
		&Request{Code: "Refill-tri200", PublishCode: "tri200", Fn: "PublishPulsa"},
		&Request{Code: "Refill-axis25", PublishCode: "axis25", Fn: "PublishPulsa"},
		&Request{Code: "Refill-axis50", PublishCode: "axis150", Fn: "PublishPulsa"},
		&Request{Code: "Refill-axis100", PublishCode: "axis100", Fn: "PublishPulsa"},
		&Request{Code: "Refill-axis150", PublishCode: "axis150", Fn: "PublishPulsa"},
		&Request{Code: "Refill-axis200", PublishCode: "axis200", Fn: "PublishPulsa"},
	}

	return kudoRouting, funKudoMap

}
