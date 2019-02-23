package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestReadingConfig(t *testing.T) {
	newConf := TwuserConfig{
		accessSecret: "",
		accessToken:  "",
		appApiKey:    "",
		appApiSecret: "",
	}

	newConf.readConfig()
	assert.NotEqual(t, "", newConf.accessSecret)
	assert.NotEqual(t, "", newConf.accessToken)
	assert.NotEqual(t, "", newConf.appApiKey)
	assert.NotEqual(t, "", newConf.appApiSecret)
}
