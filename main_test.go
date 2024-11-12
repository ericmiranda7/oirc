package main

import (
	"strings"
	"testing"
)

type Result struct {
	origin  string
	command string
	params  []string
}

func TestParseMsg(t *testing.T) {
	tcs := map[string]Result{
		":*.freenode.net 353 CCIRC = #cc :@CCIRC":                                         {"*.freenode.net", "353", []string{"CCIRC", "=", "#cc", "@CCIRC"}},
		":*.freenode.net NOTICE CCIRC :*** Ident lookup timed out, using ~guest instead.": {"*.freenode.net", "NOTICE", []string{"CCIRC", "*** Ident lookup timed out, using ~guest instead."}},
		":CCIRC!~guest@freenode-kge.qup.pic9tt.IP MODE CCIRC :+wRix":                      {"CCIRC!~guest@freenode-kge.qup.pic9tt.IP", "MODE", []string{"CCIRC", "+wRix"}},
		":CCIRC!~guest@freenode-kge.qup.pic9tt.IP JOIN :#cc":                              {"CCIRC!~guest@freenode-kge.qup.pic9tt.IP", "JOIN", []string{"#cc"}},
		":CCIRC!~guest@freenode-kge.qup.pic9tt.IP PART :#cc":                              {"CCIRC!~guest@freenode-kge.qup.pic9tt.IP", "PART", []string{"#cc"}},
		":Guest4454!~guest@freenode-kge.qup.pic9tt.IP NICK :JohnC":                        {"Guest4454!~guest@freenode-kge.qup.pic9tt.IP", "NICK", []string{"JohnC"}},
	}

	for tc, result := range tcs {
		og, cmd, params := ParseMsg(tc)
		if og != result.origin {
			t.Fatalf("want %v, got %v", result.origin, og)
		}

		if cmd != result.command {
			t.Fatalf("want %v, got %v", result.command, cmd)
		}

		if len(params) != len(result.params) {
			t.Logf("want params %v, got params %v", strings.Join(result.params, ","), strings.Join(params, ","))
			t.Fatalf("want size %v, got %v", len(result.params), len(params))
		}

		for i, p := range params {
			if result.params[i] != p {
				t.Fatalf("want %v, got %v", result.params[i], p)
			}
		}
	}
}
