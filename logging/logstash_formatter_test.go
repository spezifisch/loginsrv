package logging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/sirupsen/logrus"
	. "github.com/stretchr/testify/assert"
)

// Taken from github.com/bshuster-repo/logrus-logstash-hook
// MIT License (MIT)
// Copyright (c) 2016 Boaz Shuster

func TestLogstashFormatter(t *testing.T) {
	lf := LogstashFormatter{Type: "abc"}

	someErr := &url.Error{Op: "Get", URL: "http://example.com", Err: fmt.Errorf("The error")}

	fields := logrus.Fields{
		"message": "def",
		"level":   "ijk",
		"type":    "lmn",
		"one":     1,
		"pi":      3.14,
		"bool":    true,
		"error":   someErr,
	}

	entry := logrus.WithFields(fields)
	entry.Message = "msg"
	entry.Level = logrus.InfoLevel

	b, _ := lf.Format(entry)

	var data map[string]interface{}
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	err := dec.Decode(&data)
	NoError(t, err)

	// base fields
	if data["@timestamp"] == "" {
		t.Error("expected @timestamp to be not empty")
	}
	tt := []struct {
		expected string
		key      string
	}{
		// base fields
		{"1", "@version"},
		{"abc", "type"},
		{"msg", "message"},
		{"info", "level"},
		{someErr.Error(), "error"},
		// substituted fields
		{"def", "fields.message"},
		{"ijk", "fields.level"},
		{"lmn", "fields.type"},
	}
	for _, te := range tt {
		if te.expected != data[te.key] {
			t.Errorf("expected data[%s] to be '%s' but got '%s'", te.key, te.expected, data[te.key])
		}
	}

	// formats
	if json.Number("1") != data["one"] {
		t.Errorf("expected one to be '%v' but got '%v'", json.Number("1"), data["one"])
	}
	if json.Number("3.14") != data["pi"] {
		t.Errorf("expected pi to be '%v' but got '%v'", json.Number("3.14"), data["pi"])
	}
	if true != data["bool"] {
		t.Errorf("expected bool to be '%v' but got '%v'", true, data["bool"])
	}
}
