package integration

// MbaEvent for event register
func MbaEvent(mba *Mba) ([]*Request, map[string]interface{}) {

	funMbaMap := map[string]interface{}{
		"PublishPulsa": mba.PublishPulsa,
	}

	mbaRouting := []*Request{
		&Request{Code: "Refill-xl25", PublishCode: "xl25", Fn: "PublishPulsa"},
	}

	return mbaRouting, funMbaMap
}
