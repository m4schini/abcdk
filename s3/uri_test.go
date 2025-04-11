package s3

import "testing"

func TestConnectionString_String(t *testing.T) {
	conn := ConnectionString{
		Endpoint:  "s3.abc",
		AccessKey: "access",
		SecretKey: "secret",
		Secure:    true,
		Bucket:    "bucket",
	}
	t.Log(conn.String())
}

func TestParseConn_Secure(t *testing.T) {
	secure := true
	connStr := ConnectionString{
		Endpoint:  "s3.abc",
		AccessKey: "access",
		SecretKey: "secret",
		Secure:    secure,
		Bucket:    "bucket",
	}.String()
	t.Log(connStr)

	conn, err := ParseConn(connStr)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if conn.Secure != secure {
		t.Log("Secure flag is wrong")
		t.Fail()
	}
}

func TestParseConn_SecureNotSet(t *testing.T) {
	secure := true
	connStr := "s3://access:secret@s3.abc/bucket"
	t.Log(connStr)

	conn, err := ParseConn(connStr)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if conn.Secure != secure {
		t.Log("Secure flag is wrong")
		t.Fail()
	}
}

func TestParseConn_Insecure(t *testing.T) {
	secure := false
	connStr := ConnectionString{
		Endpoint:  "s3.abc",
		AccessKey: "access",
		SecretKey: "secret",
		Secure:    secure,
		Bucket:    "bucket",
	}.String()
	t.Log(connStr)

	conn, err := ParseConn(connStr)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if conn.Secure != secure {
		t.Log("Secure flag is wrong")
		t.Fail()
	}
}

func TestParseConn_BucketName(t *testing.T) {
	bucketName := "bucket"
	connStr := ConnectionString{
		Endpoint:  "s3.abc",
		AccessKey: "access",
		SecretKey: "secret",
		Secure:    false,
		Bucket:    bucketName,
	}.String()
	t.Log(connStr)

	conn, err := ParseConn(connStr)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if conn.Bucket != bucketName {
		t.Logf("%#v", conn)
		t.Log("Bucket name is wrong:", conn.Bucket)
		t.Fail()
	}
}

func TestParseConn_BucketNameEmpty(t *testing.T) {
	bucketName := ""
	connStr := ConnectionString{
		Endpoint:  "s3.abc",
		AccessKey: "access",
		SecretKey: "secret",
		Secure:    false,
		Bucket:    bucketName,
	}.String()
	t.Log(connStr)

	conn, err := ParseConn(connStr)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if conn.Bucket != bucketName {
		t.Logf("%#v", conn)
		t.Log("Bucket name is wrong:", conn.Bucket)
		t.Fail()
	}
}
