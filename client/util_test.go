package client

import "testing"

func TestFetchTimestampInfo(t *testing.T) {
	millis, currency := fetchTimestampInfo(5245541478956032)
	t.Log(millis)
	t.Log(currency)

	millis, currency = fetchTimestampInfo(5245541478957055)
	t.Log(millis)
	t.Log(currency)
}

func TestSpreadOutUidList(t *testing.T) {
	uids := spreadOutUidList(5245541478956032, 5245541478957055)
	t.Log(len(uids))
	t.Log(uids)
}
