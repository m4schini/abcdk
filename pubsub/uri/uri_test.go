package uri

import (
	"testing"
	"time"
)

func TestConnectionString_String(t *testing.T) {
	expected := ConnectionString{
		Protocol:    ProtocolTcp,
		Address:     "emqx.auroraborealis.cloud",
		ClientId:    "test",
		KeepAlive:   10 * time.Millisecond,
		PingTimeout: 20 * time.Millisecond,
	}
	str := expected.String()
	t.Log(str)

	actual, err := ParseConn(str)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if actual.Protocol != expected.Protocol {
		t.Log("failed")
		t.FailNow()
	}
	if actual.Address != expected.Address {
		t.Log("failed", expected.Address, actual.Address)
		t.FailNow()
	}
	if actual.Port != expected.Port {
		t.Log("failed", expected.Port, actual.Port)
		t.FailNow()
	}
	if actual.ClientId != expected.ClientId {
		t.Log("failed")
		t.FailNow()
	}
	if actual.KeepAlive != expected.KeepAlive {
		t.Log("failed")
		t.FailNow()
	}
	if actual.PingTimeout != expected.PingTimeout {
		t.Log("failed")
		t.FailNow()
	}
}
