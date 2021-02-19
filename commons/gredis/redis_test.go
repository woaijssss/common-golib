package gredis

import (
	"github.com/woaijssss/common-golib/commons/setting"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// SetUp
	Setup()
	setting.RemoteSetting.SentryDSN = ""
	// TestCase
	exitVal := m.Run()
	// TearDown
	os.Exit(exitVal)
}

func TestIncrBy(t *testing.T) {
	k, err := IncrBy("test_incrBy", 1)
	if err != nil {
		t.Logf("IncrBy was wrong err : %s", err)
	}

	if k == 0 {
		t.Logf("IncrBy was wrong err : %d", k)
	}
}
