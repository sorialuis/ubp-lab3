package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUserOK(t *testing.T) {
	ur, err := ParseUser(`{
		"id":123,
		"nickname":"facu",
		"first_name":"facundo",
		
		"last_name":"soria",
		"email":"facundo.cba2@gmail.com",
		"site_id":"mla"
	}`)

	if err != nil {
		t.Error(err)
		return
	}

	assert.EqualValues(t, 123, ur.ID)
	assert.EqualValues(t, "facu", ur.Nickname)
	assert.EqualValues(t, "facundo", ur.FirstName)
	assert.EqualValues(t, "soria", ur.LastName)
	assert.EqualValues(t, "facundo.cba2@gmail.com", ur.Email)
	assert.EqualValues(t, "mla", ur.SiteID)

}

func TestParseUserFail(t *testing.T) {
	_, err := ParseUser("{]")
	assert.NotNil(t, err)
}

func TestParseEmailCheckResponseOK(t *testing.T) {
	ur, err := ParseEmailCheckResponse(`{
		"email":{
			"user":"luis",
			"domain": "gmail.com"
		},
		"email_exists": true
	}`)

	if err != nil {
		t.Error(err)
		return
	}

	assert.EqualValues(t, "luis", ur.Email.User)
	assert.EqualValues(t, "gmail.com", ur.Email.Domain)

	assert.EqualValues(t, true, ur.EmailExists)
}

/*
@authors
TODO!!
*/
