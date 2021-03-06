/**
 * donnie4w@gmail.com  tim server
 */
package impl

import (
	"github.com/zhangjunfang/im/protocol"
	"github.com/zhangjunfang/im/utils"
)

func newTid(name string, domain, resource *string) *protocol.Tid {
	tid := protocol.NewTid()
	tid.Domain = domain
	tid.Name = name
	tid.Resource = resource
	return tid
}

func OnlinePBean(tid *protocol.Tid) (pbean *protocol.TimPBean) {
	pbean = protocol.NewTimPBean()
	pbean.ThreadId = utils.TimeMills()
	pbean.FromTid = tid
	show, status := "online", "probe"
	pbean.Show, pbean.Status = &show, &status
	return
}

func OfflinePBean(tid *protocol.Tid) (pbean *protocol.TimPBean) {
	pbean = protocol.NewTimPBean()
	pbean.ThreadId = utils.TimeMills()
	pbean.FromTid = tid
	show, status := "offline", "unavailable"
	pbean.Show, pbean.Status = &show, &status
	return
}
