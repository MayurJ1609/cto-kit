package apm

import (
	"net/http"
	"net/http/httptest"
)

func Example_elastic() {
	apm, err := New()
	if err != nil {
		//Handler error
	}
	txn, _ := apm.StartWebTransaction("/users", nil, httptest.NewRequest(http.MethodGet, "/", nil))
	apm.AddAttribute(txn, "IAMFeature", "iam.manage.user.r")
	defer apm.EndTransaction(txn, nil)
	dataSegment, _ := apm.StartDataStoreSegment(txn, "Mongo", "find", "tblUsers")
	apm.EndSegment(dataSegment)
	segment, _ := apm.StartSegment(txn, "opt-service")
	apm.EndSegment(segment)
	externalSegment, _ := apm.StartExternalSegment(txn, "https://services.skorlife.com/")
	apm.EndExternalSegment(externalSegment)
}
