package utils

import "testing"

func TestGenerateAppConfig(t *testing.T) {
	tags := map[string]string{
		"alpine":   "test",
		"nginx":    "test",
		"postgres": "test",
	}
	if err := GenerateAppConfig("../docker/test/docker-compose.test1.yaml", tags); err != nil {
		t.Error(err)
	}
}
