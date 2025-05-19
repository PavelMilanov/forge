package utils

import "testing"

func TestGenerateAppConfig(t *testing.T) {
	tags := map[string]string{
		"test_alpine":   "test",
		"test_nginx":    "test",
		"test_postgres": "test",
	}
	if err := GenerateAppConfig("../docker/test/docker-compose.yaml", tags); err != nil {
		t.Error(err)
	}
}
