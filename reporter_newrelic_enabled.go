// +build heroku cloudfoundry

package newrelic

import (
	"github.com/sky-uk/newrelic/sdk"
)

type NRTxReporter struct{}

func (r *NRTxReporter) ReportError(txnID int64, exceptionType, errorMessage, stackTrace, stackFrameDelim string) (int, error) {
	return sdk.TransactionNoticeError(txnID, exceptionType, errorMessage, stackTrace, stackFrameDelim)
}
