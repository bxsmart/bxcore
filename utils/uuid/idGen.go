package uuid

import (
	snowFlake "github.com/bxsmart/bxcore/utils/goSnowFlake"
)

const (
	CWEpoch = 1546272000000
)

var (
	idWorkers          = map[int64]*snowFlake.IdWorker{}
	defaultIdWorker, _ = snowFlake.NewIdWorker(0, CWEpoch)
)

func IdGen() int64 {
	uuid, _ := defaultIdWorker.NextId()
	return uuid
}

func IdGenDef(workerId int64) int64 {
	idWorker := idWorkers[workerId]
	if idWorker == nil {
		idWorker, _ = snowFlake.NewIdWorker(int64(workerId), CWEpoch)
	}

	uuid, _ := idWorker.NextId()
	return uuid
}

func IdGenEx(dataCenterId int64, workerId int64) int64 {
	if int64(dataCenterId) > snowFlake.CMaxDataCenter || dataCenterId < 0 {
		panic("data center not fit")
	}

	if int64(workerId) > snowFlake.CMaxWorker || workerId < 0 {
		panic("worker not fit")
	}

	nodeId := dataCenterId<<snowFlake.CDataCenterShift | workerId<<snowFlake.CWorkerIdShift
	idWorker := idWorkers[nodeId]
	if idWorker == nil {
		idWorker, _ = snowFlake.NewIdWorkerEx(dataCenterId, workerId, CWEpoch)
	}

	uuid, _ := idWorker.NextId()
	return uuid
}
