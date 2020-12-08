/**
*  @file
*  @copyright defined in go-scdo/LICENSE
 */

package core

import (
	"github.com/scdoproject/go-scdo/common"
)

func (c *core) handleFinalCommitted() error {
	c.logger.Debug("Received a final committed proposal")
	c.startNewRound(common.Big0)
	return nil
}
