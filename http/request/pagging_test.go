package request_test

import (
	"net/http"
	"testing"

	"github.com/infraboard/mcube/http/request"
	"github.com/stretchr/testify/assert"
)

func TestOffSet(t *testing.T) {
	shoud := assert.New(t)

	req, _ := http.NewRequest("GET", "/", nil)
	pr := request.NewPageRequestFromHTTP(req)
	shoud.Equal(int64(0), pr.Offset())
	shoud.Equal(uint(20), pr.PageSize)
	shoud.Equal(uint(1), pr.PageNumber)

	req, _ = http.NewRequest("GET", "/?page_size=20&page_number=1", nil)
	pr = request.NewPageRequestFromHTTP(req)
	shoud.Equal(int64(0), pr.Offset())

	req, _ = http.NewRequest("GET", "/?page_size=20&page_number=2", nil)
	pr = request.NewPageRequestFromHTTP(req)
	shoud.Equal(int64(20), pr.Offset())

	req, _ = http.NewRequest("GET", "/?page_size=20&page_number=2&offset=18", nil)
	pr = request.NewPageRequestFromHTTP(req)
	shoud.Equal(int64(18), pr.Offset())

	req, _ = http.NewRequest("GET", "/?page_size=20&offset=18", nil)
	pr = request.NewPageRequestFromHTTP(req)
	shoud.Equal(int64(18), pr.Offset())
	shoud.Equal(uint(20), pr.PageSize)

	req, _ = http.NewRequest("GET", "/?offset=18", nil)
	pr = request.NewPageRequestFromHTTP(req)
	shoud.Equal(int64(18), pr.Offset())
	shoud.Equal(uint(20), pr.PageSize)

}
