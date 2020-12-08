/**
*  @file
*  @copyright defined in go-scdo/LICENSE
 */

package core

import (
	"github.com/scdoproject/go-scdo/consensus/istanbul"
)

type backlogEvent struct {
	src istanbul.Validator
	msg *message
}

type timeoutEvent struct{}
