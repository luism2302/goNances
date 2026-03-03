package auth

import "testing"

func TestHashAndCheck(t *testing.T) {
	type testCase struct {
		password     string
		checkAgainst string
		shouldMatch  bool
	}
	testCases := []testCase{
		{password: "passwordtest123", checkAgainst: "passwordtest123", shouldMatch: true},
		{password: "passwordtest123", checkAgainst: "notamatch", shouldMatch: false},
		{password: "!!!!234343this))*(*shouldmatch", checkAgainst: "!!!!234343this))*(*shouldmatch", shouldMatch: true},
	}

	for _, tc := range testCases {
		hashed, err := HashPassword(tc.password)
		if err != nil {
			t.Fatalf("Got error hashing: %v", err)
		}
		check, err := CheckHashedPassword(tc.checkAgainst, hashed)
		if err != nil {
			t.Fatalf("Got error checking hash: %v", err)
		}

		if check != tc.shouldMatch {
			t.Fatalf("Password: %s, Checked Against: %s, Got: %v", tc.password, tc.checkAgainst, tc.shouldMatch)
		}
	}
}
