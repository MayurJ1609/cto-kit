package apm

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"go.elastic.co/apm"
	"go.elastic.co/apm/transport"
)

// elastic holds all the tracer manages the sampling and sending of transactions to
// Elastic APM.
type elastic struct {
	tracer *apm.Tracer
}

func newElastic(name, token string, urls ...string) (*elastic, error) {
	transport, err := transport.NewHTTPTransport()
	if err != nil {
		return nil, err
	}
	for _, u := range urls {
		u = strings.TrimSpace(u)
		if url, err := url.Parse(u); err == nil {
			transport.SetServerURL(url)
		}
	}
	transport.SetSecretToken(token)
	tracer := apm.DefaultTracer
	tracer.Transport = transport
	tracer.Service.Name = name
	tracer.Service.Environment = os.Getenv("ENV")
	return &elastic{
		tracer: tracer,
	}, nil
}

func (n *elastic) StartTransaction(name string) (transaction interface{}, err error) {
	opts := apm.TransactionOptions{
		Start: time.Now(),
	}
	txn := n.tracer.StartTransactionOptions(name, "request", opts)
	txn.Result = "Success"
	return txn, nil
}

func (n *elastic) StartWebTransaction(name string, w http.ResponseWriter, req *http.Request) (transaction interface{}, err error) {
	if req != nil {
		txn := n.tracer.StartTransaction(name, req.Method)
		txn.Context.SetHTTPRequest(req)
		txn.Outcome = "success"
		return txn, nil
	}
	return nil, ErrInvalidRequest
}

func (n *elastic) EndTransaction(transaction interface{}) error {
	if transaction == nil {
		return ErrInvalidTrans
	}
	if txn, ok := transaction.(*apm.Transaction); ok && txn != nil {
		txn.End()
		return nil
	}
	return ErrInvalidDataType
}

func (n *elastic) StartSegment(name string, transaction interface{}) (interface{}, error) {
	if transaction == nil {
		return nil, ErrInvalidTrans
	}
	if txn, ok := transaction.(*apm.Transaction); ok && txn.TransactionData != nil {
		opts := apm.SpanOptions{
			Start: time.Now(),
		}
		span := txn.StartSpanOptions(name, "", opts)
		return span, nil
	}
	return nil, ErrInvalidDataType
}

func (n *elastic) EndSegment(segment interface{}) error {
	if segment == nil {
		return ErrInvalidSegment
	}
	if span, ok := segment.(*apm.Span); ok && span != nil {
		span.End()
		return nil
	}
	return ErrInvalidDataType
}

func (n *elastic) StartDataStoreSegment(name string, transaction interface{}, operation string, collectionName string, operations ...Operation) (interface{}, error) {
	if txn, ok := transaction.(*apm.Transaction); ok && txn.TransactionData != nil {
		opt := Operation{}
		for _, o := range operations {
			opt = o
		}
		opts := apm.SpanOptions{
			Start: time.Now(),
		}
		span := txn.StartSpanOptions(name, fmt.Sprintf("db.%s.%s", collectionName, operation), opts)
		if !span.Dropped() {
			span.Context.SetDatabase(apm.DatabaseSpanContext{
				Instance:  opt.Instance,
				Statement: opt.Statement,
				Type:      name,
				User:      opt.User,
			})
		}
		return span, nil
	}
	return nil, ErrInvalidDataType
}

func (n *elastic) EndDataStoreSegment(segment interface{}) error {
	return n.EndSegment(segment)
}

func (n *elastic) StartExternalSegment(transaction interface{}, url string) (interface{}, error) {
	if txn, ok := transaction.(*apm.Transaction); ok && txn.TransactionData != nil {
		opts := apm.SpanOptions{
			Start: time.Now(),
		}
		span := txn.StartSpanOptions(url, "", opts)
		return span, nil
	}
	return nil, ErrInvalidDataType
}

func (n *elastic) StartExternalWebSegment(transaction interface{}, req *http.Request) (interface{}, error) {
	if txn, ok := transaction.(*apm.Transaction); ok && txn.TransactionData != nil {
		opts := apm.SpanOptions{
			Start: time.Now(),
		}
		if req != nil && req.URL != nil {
			span := txn.StartSpanOptions(req.URL.Path, "", opts)
			return span, nil
		}
		return nil, ErrInvalidRequest
	}
	return nil, ErrInvalidDataType
}

func (n *elastic) EndExternalSegment(segment interface{}) error {
	return n.EndSegment(segment)
}

func (n *elastic) NoticeError(transaction interface{}, err error) error {
	if transaction == nil {
		return ErrInvalidTrans
	}
	if txn, ok := transaction.(*apm.Transaction); ok && txn != nil {
		if err != nil {
			txn.Outcome = "failure"
			txn.Result = err.Error()
			n.tracer.NewError(err).Send()
		}
		return nil
	}
	return ErrInvalidDataType
}

func (n *elastic) AddAttribute(transaction interface{}, key, value string) error {
	if transaction == nil {
		return ErrInvalidTrans
	}
	if txn, ok := transaction.(*apm.Transaction); ok && txn.TransactionData != nil {
		txn.Context.SetLabel(key, value)
		return nil
	}
	return ErrInvalidDataType
}
