package status

import (
	"io"
	"net/http"
	"testing"

	"github.com/FishGoddess/errors"
)

func TestRegisterStatus(t *testing.T) {
	registered := New(666, "xxx", nil)
	RegisterStatus(registered)

	found := false
	for _, status := range statuses {
		if status == registered {
			found = true
			break
		}
	}

	if !found {
		t.Error("registered status not found", statuses)
	}
}

// go test -v -cover -run=^TestRegisterStatuses$
func TestRegisterStatuses(t *testing.T) {
	registered := []*Status{
		New(123, "abc", nil),
		New(666, "xxx", nil),
	}

	RegisterStatuses(registered...)

	for i := range registered {
		found := false
		for _, status := range statuses {
			if status == registered[i] {
				found = true
				break
			}
		}

		if !found {
			t.Error("registered status not found", statuses)
		}
	}
}

func TestParse(t *testing.T) {
	codeBadRequest := int32(http.StatusBadRequest)
	codeForbidden := int32(http.StatusForbidden)
	codeNotFound := int32(http.StatusNotFound)
	codeRequestTimeout := int32(http.StatusRequestTimeout)
	codeDBError := int32(1100)
	codePageTokenInvalid := int32(1200)

	RegisterStatus(New(codeBadRequest, "bad request", errors.IsBadRequest))
	RegisterStatus(New(codeForbidden, "forbidden", errors.IsForbidden))
	RegisterStatus(New(codeNotFound, "not found", errors.IsNotFound))
	RegisterStatus(New(codeRequestTimeout, "request timeout", errors.IsRequestTimeout))
	RegisterStatus(New(codeDBError, "db error", errors.IsDBError))
	RegisterStatus(New(codePageTokenInvalid, "page token invalid", errors.IsPageTokenInvalid))

	type parseTestCase struct {
		target error
		code   int32
		msg    string
	}

	parseTestCases := []*parseTestCase{
		{target: errors.BadRequest(nil), code: 0, msg: ""},
		{target: errors.Forbidden(nil), code: 0, msg: ""},
		{target: errors.NotFound(nil), code: 0, msg: ""},
		{target: errors.RequestTimeout(nil), code: 0, msg: ""},
		{target: errors.InternalServerError(nil), code: 0, msg: ""},
		{target: errors.DBError(nil), code: 0, msg: ""},
		{target: errors.PageTokenInvalid(nil), code: 0, msg: ""},

		{target: errors.BadRequest(io.EOF), code: codeBadRequest, msg: "bad request"},
		{target: errors.Forbidden(io.EOF), code: codeForbidden, msg: "forbidden"},
		{target: errors.NotFound(io.EOF), code: codeNotFound, msg: "not found"},
		{target: errors.RequestTimeout(io.EOF), code: codeRequestTimeout, msg: "request timeout"},
		{target: errors.InternalServerError(io.EOF), code: CodeUnknown, msg: "Server is too busy, please try again later."},
		{target: errors.DBError(io.EOF), code: codeDBError, msg: "db error"},
		{target: errors.PageTokenInvalid(io.EOF), code: codePageTokenInvalid, msg: "page token invalid"},

		{target: errors.BadRequest(io.EOF, errors.WithMsg("BadRequest")), code: codeBadRequest, msg: "BadRequest"},
		{target: errors.Forbidden(io.EOF, errors.WithMsg("Forbidden")), code: codeForbidden, msg: "Forbidden"},
		{target: errors.NotFound(io.EOF, errors.WithMsg("NotFound")), code: codeNotFound, msg: "NotFound"},
		{target: errors.RequestTimeout(io.EOF, errors.WithMsg("RequestTimeout")), code: codeRequestTimeout, msg: "RequestTimeout"},
		{target: errors.InternalServerError(io.EOF, errors.WithMsg("InternalServerError")), code: CodeUnknown, msg: "InternalServerError"},
		{target: errors.DBError(io.EOF, errors.WithMsg("DBError")), code: codeDBError, msg: "DBError"},
		{target: errors.PageTokenInvalid(io.EOF, errors.WithMsg("PageTokenInvalid")), code: codePageTokenInvalid, msg: "PageTokenInvalid"},
	}

	for _, parseTestCase := range parseTestCases {
		code, msg := Parse(parseTestCase.target)
		if code != parseTestCase.code {
			t.Errorf("code %d != parseTestCase.code %d", code, parseTestCase.code)
		}

		if msg != parseTestCase.msg {
			t.Errorf("msg %s != parseTestCase.msg %s", msg, parseTestCase.msg)
		}
	}
}
