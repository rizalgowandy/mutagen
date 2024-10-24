package core

import (
	"testing"
)

// TestEntryKindSynchronizable tests EntryKind.synchronizable.
func TestEntryKindSynchronizable(t *testing.T) {
	// Define test cases.
	tests := []struct {
		kind     EntryKind
		expected bool
	}{
		{EntryKind_Directory, true},
		{EntryKind_File, true},
		{EntryKind_SymbolicLink, true},
		{EntryKind_Untracked, false},
		{EntryKind_Problematic, false},
		{EntryKind_PhantomDirectory, false},
	}

	// Process test cases.
	for i, test := range tests {
		if synchronizable := test.kind.synchronizable(); synchronizable != test.expected {
			t.Errorf("test case %d: synchronizability does not match expected: %t != %t",
				i, synchronizable, test.expected,
			)
		}
	}
}

// TestEntryKindUnmarshal tests that unmarshaling from a string specification
// succeeeds for EntryKind.
func TestEntryKindUnmarshal(t *testing.T) {
	// Set up test cases.
	testCases := []struct {
		text          string
		expected      EntryKind
		expectFailure bool
	}{
		{"", EntryKind_Directory, true},
		{"asdf", EntryKind_Directory, true},
		{"directory", EntryKind_Directory, false},
		{"file", EntryKind_File, false},
		{"symlink", EntryKind_SymbolicLink, false},
		{"untracked", EntryKind_Untracked, false},
		{"problematic", EntryKind_Problematic, false},
		{"phantom-directory", EntryKind_PhantomDirectory, false},
	}

	// Process test cases.
	for _, testCase := range testCases {
		var kind EntryKind
		if err := kind.UnmarshalText([]byte(testCase.text)); err != nil {
			if !testCase.expectFailure {
				t.Errorf("unable to unmarshal text (%s): %s", testCase.text, err)
			}
		} else if testCase.expectFailure {
			t.Error("unmarshaling succeeded unexpectedly for text:", testCase.text)
		} else if kind != testCase.expected {
			t.Errorf(
				"unmarshaled entry kind (%s) does not match expected (%s)",
				kind,
				testCase.expected,
			)
		}
	}
}

func init() {
	// Enable wildcard problem matching for tests.
	entryEqualWildcardProblemMatch = true
}

// entryEnsureValidTestCases are test cases shared between TestEntryEnsureValid,
// TestArchiveEnsureValid, and TestChangeEnsureValid.
var entryEnsureValidTestCases = []struct {
	// entry is the entry to validate.
	entry *Entry
	// synchronizable is the synchronizability condition under which to run
	// validation.
	synchronizable bool
	// expected is the expected validity.
	expected bool
}{
	// Test synchronizable content.
	{tN, false, true},
	{tN, true, true},
	{tF1, false, true},
	{tF1, true, true},
	{tF3E, false, true},
	{tF3E, true, true},
	{tSR, false, true},
	{tSR, true, true},
	{tSA, false, true},
	{tSA, true, true},
	{tD0, false, true},
	{tD0, true, true},
	{tD1, false, true},
	{tD1, true, true},
	{tD3E, false, true},
	{tD3E, true, true},
	{tDCC, false, true},
	{tDCC, true, true},
	{tDSR, false, true},
	{tDSR, true, true},
	{tDSA, false, true},
	{tDSA, true, true},
	{tDM, false, true},
	{tDM, true, true},

	// Test unsynchronizable content.
	{tU, false, true},
	{tU, true, false},
	{tP1, false, true},
	{tP1, true, false},
	{tDU, false, true},
	{tDU, true, false},
	{tDP1, false, true},
	{tDP1, true, false},
	{tPD0, false, true},
	{tPD0, true, false},
	{tDPD0, false, true},
	{tDPD0, true, false},

	// Test invalid content.
	{tIDDE, false, false},
	{tIDDE, true, false},
	{tIDD, false, false},
	{tIDD, true, false},
	{tIDE, false, false},
	{tIDE, true, false},
	{tIDT, false, false},
	{tIDT, true, false},
	{tIDP, false, false},
	{tIDP, true, false},
	{tIDCE, false, false},
	{tIDCE, true, false},
	{tIDCD, false, false},
	{tIDCD, true, false},
	{tIDCDD, false, false},
	{tIDCDD, true, false},
	{tIDCS, false, false},
	{tIDCS, true, false},
	{tIDCN, false, false},
	{tIDCN, true, false},
	{tIFCE, false, false},
	{tIFCE, true, false},
	{tIFC, false, false},
	{tIFC, true, false},
	{tIFT, false, false},
	{tIFT, true, false},
	{tIFP, false, false},
	{tIFP, true, false},
	{tIFDN, false, false},
	{tIFDN, true, false},
	{tIFDE, false, false},
	{tIFDE, true, false},
	{tISCE, false, false},
	{tISCE, true, false},
	{tISC, false, false},
	{tISC, true, false},
	{tISDE, false, false},
	{tISDE, true, false},
	{tISD, false, false},
	{tISD, true, false},
	{tISE, false, false},
	{tISE, true, false},
	{tISP, false, false},
	{tISP, true, false},
	{tISTE, false, false},
	{tISTE, true, false},
	{tIUCE, false, false},
	{tIUCE, true, false},
	{tIUC, false, false},
	{tIUC, true, false},
	{tIUDE, false, false},
	{tIUDE, true, false},
	{tIUD, false, false},
	{tIUD, true, false},
	{tIUE, false, false},
	{tIUE, true, false},
	{tIUT, false, false},
	{tIUT, true, false},
	{tIUP, false, false},
	{tIUP, true, false},
	{tIPCE, false, false},
	{tIPCE, true, false},
	{tIPC, false, false},
	{tIPC, true, false},
	{tIPDE, false, false},
	{tIPDE, true, false},
	{tIPD, false, false},
	{tIPD, true, false},
	{tIPE, false, false},
	{tIPE, true, false},
	{tIPT, false, false},
	{tIPT, true, false},
	{tIPPE, false, false},
	{tIPPE, true, false},
	{tII, false, false},
	{tII, true, false},
}

// TestEntryEnsureValid tests Entry.EnsureValid.
func TestEntryEnsureValid(t *testing.T) {
	// Process test cases.
	for i, test := range entryEnsureValidTestCases {
		// Compute a description for the test in case we need it.
		description := "without synchronizability requirement"
		if test.synchronizable {
			description = "when requiring synchronizability"
		}

		// Check validity.
		err := test.entry.EnsureValid(test.synchronizable)
		valid := err == nil
		if valid != test.expected {
			if valid {
				t.Errorf("test index %d: entry incorrectly classified as valid (%s)", i, description)
			} else {
				t.Errorf("test index %d: entry incorrectly classified as invalid (%s): %v", i, description, err)
			}
		}
	}
}

// testEntryWalkVisit encodes a visit operation from Entry.walk.
type testEntryWalkVisit struct {
	// path is the visited path.
	path string
	// entry is the visited entry.
	entry *Entry
}

// TestEntryWalk tests Entry.walk.
func TestEntryWalk(t *testing.T) {
	// Define test cases.
	tests := []struct {
		path     string
		entry    *Entry
		reverse  bool
		expected []testEntryWalkVisit
	}{
		{"", tN, false, []testEntryWalkVisit{{"", tN}}},
		{"", tF1, false, []testEntryWalkVisit{{"", tF1}}},
		{"", tSR, false, []testEntryWalkVisit{{"", tSR}}},
		{"", tD0, false, []testEntryWalkVisit{{"", tD0}}},
		{"", tD1, false, []testEntryWalkVisit{{"", tD0}, {"file", tF1}}},
		{"", tD1, true, []testEntryWalkVisit{{"file", tF1}, {"", tD0}}},
		{"base", tD1, false, []testEntryWalkVisit{{"base", tD0}, {"base/file", tF1}}},
		{"base", tD1, true, []testEntryWalkVisit{{"base/file", tF1}, {"base", tD0}}},
		{
			"",
			nested("child", tD1),
			false,
			[]testEntryWalkVisit{{"", tD0}, {"child", tD0}, {"child/file", tF1}},
		},
		{"", tU, false, []testEntryWalkVisit{{"", tU}}},
		{"", tP1, false, []testEntryWalkVisit{{"", tP1}}},
	}

	// Process test cases.
	for i, test := range tests {
		// Perform walking and record visits.
		var visits []testEntryWalkVisit
		test.entry.walk(test.path, func(path string, entry *Entry) {
			visits = append(visits, testEntryWalkVisit{path, entry.Copy(EntryCopyBehaviorSlim)})
		}, test.reverse)

		// Verify that the number of visits was correct.
		if len(visits) != len(test.expected) {
			t.Errorf("test index %d: visit count did not match expected: %d != %d",
				i, len(visits), len(test.expected),
			)
		}

		// Verify that the visits match what was expected.
		for v, visit := range visits {
			expectedVisit := test.expected[v]
			if visit.path != expectedVisit.path {
				t.Errorf("test index %d, visit index %d: visit path did not match expected: %s != %s",
					i, v, visit.path, expectedVisit.path,
				)
			}
			if !visit.entry.Equal(expectedVisit.entry, true) {
				t.Errorf("test index %d, visit index %d: visit entry did not match expected", i, v)
			}
		}
	}
}

// TestEntryCount tests Entry.Count.
func TestEntryCount(t *testing.T) {
	// Define test cases.
	tests := []struct {
		entry    *Entry
		expected uint64
	}{
		{tN, 0},
		{tF1, 1},
		{tSR, 1},
		{tD0, 1},
		{tD1, 2},
		{tDSR, 3},
		{tDCC, 3},
		{tU, 0},
		{tP1, 0},
		{tDU, 1},
		{tDP1, 1},
	}

	// Process test cases.
	for i, test := range tests {
		if count := test.entry.Count(); count != test.expected {
			t.Errorf("test index %d: count did not match expected: %d != %d", i, count, test.expected)
		}
	}
}

// TestEntryEqual tests Entry.Equal.
func TestEntryEqual(t *testing.T) {
	// Define test cases.
	tests := []struct {
		first    *Entry
		second   *Entry
		deep     bool
		expected bool
	}{
		{tN, tN, false, true},
		{tN, tN, true, true},
		{tN, tF1, false, false},
		{tN, tF1, true, false},
		{tF1, tF1, false, true},
		{tF1, tF1, true, true},
		{tF1, tF2, false, false},
		{tF1, tF2, true, false},
		{tF1, tD0, false, false},
		{tF1, tD0, true, false},
		{tSR, tSR, false, true},
		{tSR, tSR, true, true},
		{tSR, tSA, false, false},
		{tSR, tSA, true, false},
		{tD1, tD1, false, true},
		{tD1, tD1, true, true},
		{tD1, tD2, false, true},
		{tD1, tD2, true, false},
		{tD1, tDCC, false, true},
		{tD1, tDCC, true, false},
		{tF1, tU, false, false},
		{tF1, tU, true, false},
		{tF1, tP1, false, false},
		{tF1, tP1, true, false},
		{tU, tU, false, true},
		{tU, tU, true, true},
		{tP1, tP1, false, true},
		{tP1, tP1, true, true},
		{tDU, tDU, false, true},
		{tDU, tDU, true, true},
		{tDP1, tDP1, false, true},
		{tDP1, tDP1, true, true},
		{tPD0, tPD0, false, true},
		{tPD0, tPD0, true, true},
	}

	// Process test cases.
	for i, test := range tests {
		// Compute a description for the test in case we need it.
		description := "shallow"
		if test.deep {
			description = "deep"
		}

		// Check equivalence.
		equal := test.first.Equal(test.second, test.deep)
		if equal != test.expected {
			if equal {
				t.Errorf("test index %d: entries incorrectly classified as equal (%s)", i, description)
			} else {
				t.Errorf("test index %d: entries incorrectly classified as unequal (%s)", i, description)
			}
		}

		// Check that equivalence (or non-equivalence) is symmetric
		reverseEqual := test.second.Equal(test.first, test.deep)
		if reverseEqual != equal {
			t.Errorf("test index %d: (%s) entry equivalence not symmetric: %t != %t",
				i, description, reverseEqual, equal,
			)
		}
	}
}

// TestEntryCopy tests Entry.Copy.
func TestEntryCopy(t *testing.T) {
	// Define test cases.
	tests := []struct {
		entry        *Entry
		copyBehavior EntryCopyBehavior
		expected     *Entry
	}{
		{tN, EntryCopyBehaviorDeep, tN},
		{tN, EntryCopyBehaviorDeepPreservingLeaves, tN},
		{tN, EntryCopyBehaviorShallow, tN},
		{tN, EntryCopyBehaviorSlim, tN},

		{tF1, EntryCopyBehaviorDeep, tF1},
		{tF1, EntryCopyBehaviorDeepPreservingLeaves, tF1},
		{tF1, EntryCopyBehaviorShallow, tF1},
		{tF1, EntryCopyBehaviorSlim, tF1},

		{tF3E, EntryCopyBehaviorDeep, tF3E},
		{tF3E, EntryCopyBehaviorDeepPreservingLeaves, tF3E},
		{tF3E, EntryCopyBehaviorShallow, tF3E},
		{tF3E, EntryCopyBehaviorSlim, tF3E},

		{tSA, EntryCopyBehaviorDeep, tSA},
		{tSA, EntryCopyBehaviorDeepPreservingLeaves, tSA},
		{tSA, EntryCopyBehaviorShallow, tSA},
		{tSA, EntryCopyBehaviorSlim, tSA},

		{tD0, EntryCopyBehaviorDeep, tD0},
		{tD0, EntryCopyBehaviorDeepPreservingLeaves, tD0},
		{tD0, EntryCopyBehaviorShallow, tD0},
		{tD0, EntryCopyBehaviorSlim, tD0},

		{tD1, EntryCopyBehaviorDeep, tD1},
		{tD1, EntryCopyBehaviorDeepPreservingLeaves, tD1},
		{tD1, EntryCopyBehaviorShallow, tD1},
		{tD1, EntryCopyBehaviorSlim, tD0},

		{tU, EntryCopyBehaviorDeep, tU},
		{tU, EntryCopyBehaviorDeepPreservingLeaves, tU},
		{tU, EntryCopyBehaviorShallow, tU},
		{tU, EntryCopyBehaviorSlim, tU},

		{tP1, EntryCopyBehaviorDeep, tP1},
		{tP1, EntryCopyBehaviorDeepPreservingLeaves, tP1},
		{tP1, EntryCopyBehaviorShallow, tP1},
		{tP1, EntryCopyBehaviorSlim, tP1},
	}

	// Process test cases.
	for i, test := range tests {
		// Perform copying and verify that the result matches what's expected.
		result := test.entry.Copy(test.copyBehavior)
		if !result.Equal(test.expected, true) {
			t.Errorf("test index %d: copy result does not match expected", i)
		}
	}
}

// TestEntrySynchronizable tests Entry.synchronizable.
func TestEntrySynchronizable(t *testing.T) {
	// Define test cases.
	tests := []struct {
		entry    *Entry
		expected *Entry
	}{
		{tN, tN},
		{tF1, tF1},
		{tSR, tSR},
		{tD0, tD0},
		{tD1, tD1},
		{tU, tN},
		{tP1, tN},
		{tDU, tD0},
		{tDP1, tD0},
	}

	// Process test cases.
	for i, test := range tests {
		if synchronizable := test.entry.synchronizable(); !synchronizable.Equal(test.expected, true) {
			t.Errorf("test index %d: synchronizable subentry does not match expected", i)
		}
	}
}

// TestEntryProblems tests Entry.Problems.
func TestEntryProblems(t *testing.T) {
	// Define test cases.
	tests := []struct {
		entry    *Entry
		expected []*Problem
	}{
		{tN, nil},
		{tF1, nil},
		{tSR, nil},
		{tD0, nil},
		{tD1, nil},
		{tU, nil},
		{tP1, []*Problem{{Error: tP1.Problem}}},
		{tDU, nil},
		{tDP1, []*Problem{{Path: "problematic", Error: tP1.Problem}}},
	}

	// Process test cases.
	for i, test := range tests {
		if problems := test.entry.Problems(); !testingProblemListsEqual(problems, test.expected) {
			t.Errorf("test index %d: entry problems do not match expected", i)
		}
	}
}
