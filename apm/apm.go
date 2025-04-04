package apm

import (
	"net/http"
	"os"
	"sync"
)

// agentType is used to identify APM process supported by Skorlife.
type agentType string

const (
	// Elastic Application monitoring supported by monitoring library
	Elastic agentType = "Elastic"
)

// Operation sets options for database tracing.
type Operation struct {
	// Instance holds the database instance name.
	Instance string
	// Statement holds the statement executed in the span,
	// e.g. "SELECT * FROM foo".
	Statement string
	// Type holds the database type, e.g. "sql".
	Type string
	// User holds the username used for database access.
	User string
}

// handler contains all APM regarding the request context
type handler interface {
	StartTransaction(name string) (transaction interface{}, err error)
	StartWebTransaction(name string, rw http.ResponseWriter, req *http.Request) (transaction interface{}, err error)
	EndTransaction(transaction interface{}) (err error)
	StartSegment(segmentName string, Transaction interface{}) (segment interface{}, err error)
	EndSegment(segment interface{}) (err error)
	StartDataStoreSegment(segmentName string, transaction interface{}, operation string, collectionName string, operations ...Operation) (segment interface{}, err error)
	EndDataStoreSegment(segment interface{}) (err error)
	StartExternalSegment(transaction interface{}, url string) (externalSegment interface{}, err error)
	StartExternalWebSegment(transaction interface{}, req *http.Request) (externalSegment interface{}, err error)
	EndExternalSegment(segment interface{}) error
	NoticeError(transaction interface{}, err error) error
	AddAttribute(transaction interface{}, key, val string) error
}

type Agent interface {
	// Enable method allows if monitoring is disable data will not be pushed to APM
	Enable(monitoring bool)
	// StartTransaction returns a new Transaction with the specified
	// name and type, and with the start time set to the current time.
	// This is equivalent to calling StartTransactionOptions with a
	// zero TransactionOptions.
	StartTransaction(name string) (interface{}, error)

	// StartWebTransaction begins a web transaction.
	// * The Transaction is considered a web transaction if an http.Request
	//   is provided.
	// * The transaction returned implements the http.ResponseWriter
	//   interface.  Provide your ResponseWriter as a parameter and
	//   then use the Transaction in its place to instrument the response
	//   code and response headers.
	StartWebTransaction(name string, rw http.ResponseWriter, req *http.Request) (interface{}, error)

	// EndTransaction function finishes the transaction, stopping all further
	// instrumentation.
	//
	// Calling End will set trans TransactionData field to nil, so callers
	// must ensure tx is not updated after End returns.
	EndTransaction(trans interface{}, err error) error

	// StartSegment function instrument segments to the particular transaction.
	//
	// trans: Transaction object which returned from StartTransaction function.
	// name: The name of the segment
	//
	StartSegment(trans interface{}, name string) (interface{}, error)
	// EndSegment function finishes the Segment.
	//
	// Segment: Segment object which returned from StartSegment function.
	EndSegment(segment interface{}) error

	// StartDataStoreSegment function  is used to instrument calls to databases
	// and object stores.
	//
	//  trans: Transaction object which returned from StartTransaction function.
	//  name: The name of the segment
	//  operations: Operation is the relevant action, e.g. "SELECT" or "GET".
	//  table: CollectionName is the table name or group name.
	//
	StartDataStoreSegment(trans interface{}, name string, operation string, table string, operations ...Operation) (interface{}, error)
	// EndDataStoreSegment function finishes the datastore segment.
	//
	// segment: Segment object which returned from StartDataStoreSegment function.
	EndDataStoreSegment(segment interface{}) (err error)

	// StartExternalSegment function is used to instrument external
	// calls.StartExternalSegment is recommended when you do not have
	// access to an http.Request.
	//
	//    trans	: Transaction object which returned from StartTransaction function.
	//    url	: URL field should be used to indicate the endpoint.
	StartExternalSegment(trans interface{}, url string) (interface{}, error)
	// StartExternalWebSegment function is used to instrument external
	// calls.StartExternalWebSegment is recommended when you have access to an http.Request.
	//
	//    trans: Transaction object which returned from StartTransaction function.
	//    request: URL field should be used to indicate the endpoint.
	StartExternalWebSegment(trans interface{}, request *http.Request) (interface{}, error)

	// EndExternalSegment function finishes the external segment.
	//
	//    Segment: Segment object which returned from StartExternalSegment function.
	EndExternalSegment(segment interface{}) error

	// NoticeError function records an error.
	NoticeError(trans interface{}, err error) error

	// AddAttribute function adds a key value pair information to the current transaction.
	//
	//  trans: Transaction object which returned from StartTransaction function.
	//	key: key information specifies the name of the information attribute holds.
	//	val: This attribute specifies the value of the information attribute holds.
	AddAttribute(trans interface{}, key, val string) error
}

// Agent holds all application monitoring
type agent struct {
	// application name is used by New Relic to link data across servers.
	appName string
	// enabled controls whether the agent will communicate with the APM servers or not
	// Setting this to be false is useful in testing and staging situations.
	enabled bool
	monitor handler
	mutex   *sync.Mutex
}

// New creates an Application and spawns goroutines to manage the
// aggregation and harvesting of data.  On success, a non-nil application and a
// nil error are returned. On failure, a nil Application and a non-nil error
// are returned. Applications do not share global state, therefore it is safe
// to create multiple applications.
func New(options ...Option) (Agent, error) {
	opt := &option{
		enable:      true,
		agentType:   Elastic,
		serviceName: os.Getenv("ELASTIC_APM_SERVICE_NAME"),
		serverURL:   os.Getenv("ELASTIC_APM_SERVER_URL"),
		secretToken: os.Getenv("ELASTIC_APM_SECRET_TOKEN"),
	}
	if !opt.enable {
		return &agent{appName: os.Getenv("CONFIG_APP"), enabled: false, mutex: &sync.Mutex{}}, nil
	}

	for _, option := range options {
		option(opt)
	}
	agent := &agent{enabled: true, mutex: &sync.Mutex{}}
	switch opt.agentType {
	case Elastic:
		monitor, _ := newElastic(opt.serviceName, opt.secretToken, opt.serverURL)
		agent.monitor = monitor
	default:
		return nil, ErrUnsupported
	}
	return agent, nil
}

func (agent *agent) Enable(monitoring bool) {
	if agent != nil || agent.monitor != nil {
		agent.mutex.Lock()
		agent.enabled = monitoring
		agent.mutex.Unlock()
	}
}

// isMonitoringEnabled determines whether the APM need enabled or not
// if monitoring is disable data will not be pushed to APM
func (agent *agent) isMonitoringEnabled() bool {
	if agent == nil || agent.monitor == nil {
		return false
	}
	return agent.enabled
}

func (agent *agent) StartTransaction(transactionName string) (interface{}, error) {
	if !agent.isMonitoringEnabled() {
		return nil, nil
	}
	return agent.monitor.StartTransaction(transactionName)
}

func (agent *agent) StartWebTransaction(transactionName string, rw http.ResponseWriter, req *http.Request) (interface{}, error) {
	if !agent.isMonitoringEnabled() {
		return nil, nil
	}
	return agent.monitor.StartWebTransaction(transactionName, rw, req)
}

func (agent *agent) EndTransaction(trans interface{}, err error) error {
	if !agent.isMonitoringEnabled() {
		return nil
	}
	if err != nil {
		agent.monitor.NoticeError(trans, err)
	}
	return agent.monitor.EndTransaction(trans)
}

func (agent *agent) StartSegment(trans interface{}, segmentName string) (interface{}, error) {
	if !agent.isMonitoringEnabled() {
		return nil, nil
	}
	return agent.monitor.StartSegment(segmentName, trans)
}

func (agent *agent) EndSegment(segment interface{}) error {
	if !agent.isMonitoringEnabled() {
		return nil
	}
	return agent.monitor.EndSegment(segment)
}

func (agent *agent) StartDataStoreSegment(transaction interface{}, name string, operation string, table string, operations ...Operation) (interface{}, error) {
	if !agent.isMonitoringEnabled() {
		return nil, nil
	}
	return agent.monitor.StartDataStoreSegment(name, transaction, operation, table, operations...)
}

func (agent *agent) EndDataStoreSegment(segment interface{}) error {
	if !agent.isMonitoringEnabled() {
		return nil
	}

	return agent.monitor.EndDataStoreSegment(segment)
}

func (agent *agent) StartExternalSegment(trans interface{}, url string) (interface{}, error) {
	if !agent.isMonitoringEnabled() {
		return nil, nil
	}
	return agent.monitor.StartExternalSegment(trans, url)
}

func (agent *agent) StartExternalWebSegment(trans interface{}, request *http.Request) (interface{}, error) {
	if !agent.isMonitoringEnabled() {
		return nil, nil
	}

	return agent.monitor.StartExternalWebSegment(trans, request)
}

func (agent *agent) EndExternalSegment(segment interface{}) error {
	if !agent.isMonitoringEnabled() {
		return nil
	}

	return agent.monitor.EndExternalSegment(segment)
}

func (agent *agent) NoticeError(trans interface{}, err error) error {
	if !agent.isMonitoringEnabled() {
		return nil
	}
	return agent.monitor.NoticeError(trans, err)
}

func (agent *agent) AddAttribute(trans interface{}, key, val string) error {
	if !agent.isMonitoringEnabled() {
		return nil
	}
	return agent.monitor.AddAttribute(trans, key, val)
}
