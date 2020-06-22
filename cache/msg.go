package cache

import (
	"github.com/shyptr/jianshu/model"
	"strconv"
)

const (
	MsgNum_Cache = "MsgNum"
	Msg_Cache    = "Msg"
)

type MsgNum struct {
	Uid int
}

func (m MsgNum) GetCacheKey() string {
	return MsgNum_Cache + "_" + strconv.Itoa(m.Uid)
}

func (m MsgNum) GetCachesKey() string {
	return MsgNum_Cache + "_" + strconv.Itoa(m.Uid)
}

type Msg struct {
	Uid int
	Typ model.MsgType
}

func (m Msg) GetCacheKey() string {
	return Msg_Cache + "_" + strconv.Itoa(m.Uid) + "_" + strconv.Itoa(int(m.Typ))
}

func (m Msg) GetCachesKey() string {
	return Msg_Cache + "_" + strconv.Itoa(m.Uid) + "_" + strconv.Itoa(int(m.Typ))
}
