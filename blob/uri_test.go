package blob

import "testing"

func TestConnectionString_String(t *testing.T) {
	conn := ConnectionString{
		Endpoint: "s3.auroraborealis.cloud",
		Key:      "key",
		Secret:   "secret",
	}

	t.Log(conn.String())
}

func TestParseConn(t *testing.T) {
	expected := ConnectionString{
		Endpoint:   "s3.auroraborealis.cloud",
		Key:        "key",
		Secret:     "YlpBxFi2hggMmtgISS7AhiIW_/:4GlmHZNsBZoIALEI",
		BucketName: "bucketName",
	}
	str := expected.String()

	actual, err := ParseConn(str)
	if err != nil {
		t.FailNow()
	}

	if actual.Endpoint != expected.Endpoint {
		t.Logf("expected: %v - actual: %v", expected.Endpoint, actual.Endpoint)
		t.FailNow()
	}
	if actual.Key != expected.Key {
		t.Logf("expected: %v - actual: %v", expected.Key, actual.Key)
		t.FailNow()
	}
	if actual.Secret != expected.Secret {
		t.Logf("expected: %v - actual: %v", expected.Secret, actual.Secret)
		t.FailNow()
	}
	if actual.BucketName != expected.BucketName {
		t.Logf("expected: %v - actual: %v", expected.BucketName, actual.BucketName)
		t.FailNow()
	}
	t.Log(str)
}

func TestParseConn_EmptyBucketName(t *testing.T) {
	expected := ConnectionString{
		Endpoint:   "s3.auroraborealis.cloud",
		Key:        "key",
		Secret:     "YlpBxFi2hggMmtgISS7AhiIW_/:4GlmHZNsBZoIALEI",
		BucketName: "",
	}
	str := expected.String()

	actual, err := ParseConn(str)
	if err != nil {
		t.FailNow()
	}

	if actual.Endpoint != expected.Endpoint {
		t.Logf("expected: %v - actual: %v", expected.Endpoint, actual.Endpoint)
		t.FailNow()
	}
	if actual.Key != expected.Key {
		t.Logf("expected: %v - actual: %v", expected.Key, actual.Key)
		t.FailNow()
	}
	if actual.Secret != expected.Secret {
		t.Logf("expected: %v - actual: %v", expected.Secret, actual.Secret)
		t.FailNow()
	}
	if actual.BucketName != expected.BucketName {
		t.Logf("expected: %v - actual: %v", expected.BucketName, actual.BucketName)
		t.FailNow()
	}
	t.Log(str)
}

func TestParseConn_EmptyEndpoint(t *testing.T) {
	expected := ConnectionString{
		Endpoint:   "",
		Key:        "key",
		Secret:     "YlpBxFi2hggMmtgISS7AhiIW_/:4GlmHZNsBZoIALEI",
		BucketName: "",
	}
	str := expected.String()

	actual, err := ParseConn(str)
	if err != nil {
		t.FailNow()
	}

	//TODO default value
	if actual.Endpoint != expected.Endpoint {
		t.Logf("expected: %v - actual: %v", expected.Endpoint, actual.Endpoint)
		t.FailNow()
	}
	if actual.Key != expected.Key {
		t.Logf("expected: %v - actual: %v", expected.Key, actual.Key)
		t.FailNow()
	}
	if actual.Secret != expected.Secret {
		t.Logf("expected: %v - actual: %v", expected.Secret, actual.Secret)
		t.FailNow()
	}
	if actual.BucketName != expected.BucketName {
		t.Logf("expected: %v - actual: %v", expected.BucketName, actual.BucketName)
		t.FailNow()
	}
	t.Log(str)
}
