package utils

import "testing"

func TestGenerateAppConfig(t *testing.T) {
	tags := map[string]string{
		"alpine":   "test",
		"nginx":    "test",
		"postgres": "test",
	}
	_, err := GenerateAppConfig("../docker/test/docker-compose.test1.yaml", "test", tags)
	if err != nil {
		t.Error(err)
	}
}
