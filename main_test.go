package main

import "testing"

type Result struct {
	origin  string
	command string
	params  []string
}

func TestParseMsg(t *testing.T) {
	tcs := map[string]Result{
		":*.freenode.net 353 CCIRC = #cc :@CCIRC":                                         Result{"*.freenode.net", "353", nil},
		":*.freenode.net NOTICE CCIRC :*** Ident lookup timed out, using ~guest instead.": Result{"*.freenode.net", "NOTICE", nil},
		":CCIRC!~guest@freenode-kge.qup.pic9tt.IP MODE CCIRC :+wRix":                      Result{"CCIRC!~guest@freenode-kge.qup.pic9tt.IP", "MODE", nil},
		":CCIRC!~guest@freenode-kge.qup.pic9tt.IP JOIN :#cc":                              Result{"CCIRC!~guest@freenode-kge.qup.pic9tt.IP", "JOIN", nil},
		":CCIRC!~guest@freenode-kge.qup.pic9tt.IP PART :#cc":                              Result{"CCIRC!~guest@freenode-kge.qup.pic9tt.IP", "PART", nil},
		":Guest4454!~guest@freenode-kge.qup.pic9tt.IP NICK :JohnC":                        Result{"Guest4454!~guest@freenode-kge.qup.pic9tt.IP", "NICK", nil},
	}

	for tc, result := range tcs {
		og, cmd, _ := ParseMsg(tc)
		if og != result.origin {
			t.Fatalf("want %v, got %v", result.origin, og)
		}

		if cmd != result.command {
			t.Fatalf("want %v, got %v", result.command, cmd)
		}
	}
}
