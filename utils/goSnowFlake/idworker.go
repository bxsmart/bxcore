// This package provides unique id in distribute system
// the algorithm is inspired by Twitter's famous snowflake
// its link is: https://github.com/twitter/snowflake/releases/tag/snowflake-2010
//

// +---------------+----------------+----------------+
// |timestamp(ms)42  | worker id(10) | sequence(12)	 |
// +---------------+----------------+----------------+

// Copyright (C) 2016 by bx-smart.info

package goSnowFlake

import (
	"errors"
	"sync"
	"time"
)

const (
	CEpoch          = 1546272000000 // Start timestamp 2019-01-01 00:00:00
	CWorkerIdBits   = 5             // Num of WorkerId Bits
	CDataCenterBits = 5             // Num of DataCentId Bits
	CSequenceBits   = 12            // Num of Sequence Bits

	CWorkerIdShift   = CSequenceBits
	CDataCenterShift = CWorkerIdShift + CWorkerIdBits
	CTimeStampShift  = CDataCenterShift + CDataCenterBits

	CMaxWorker     = -1 ^ (-1 << CWorkerIdBits)
	CMaxDataCenter = -1 ^ (-1 << CDataCenterBits)
	CSequenceMask  = -1 ^ (-1 << CSequenceBits)
)

// IdWorker struct
type IdWorker struct {
	twEpoch       int64
	lastTimeStamp int64
	workerId      int64
	dataCenterId  int64
	sequence      int64
	lock          *sync.Mutex
}

// NewIdWorker Func: Generate NewIdWorker with Given workerId
func NewIdWorkerEx(dataCenterId int64, workerId int64, twepoch int64) (iw *IdWorker, err error) {
	if int64(dataCenterId) > CMaxDataCenter || dataCenterId < 0 {
		return nil, errors.New("data center not fit")
	}

	iw = new(IdWorker)

	if int64(workerId) > CMaxWorker || workerId < 0 {
		return nil, errors.New("worker not fit")
	}

	if twepoch <= 0 {
		iw.twEpoch = CEpoch
	}

	iw.dataCenterId = dataCenterId
	iw.workerId = workerId
	iw.lastTimeStamp = -1
	iw.sequence = 0
	iw.lock = new(sync.Mutex)
	return iw, nil
}

// NewIdWorker Func: Generate NewIdWorker with Given workerId
func NewIdWorker(workerId_, twEpoch int64) (iw *IdWorker, err error) {
	workerId := workerId_ & CMaxWorker
	dataCenterId := (workerId_ >> CDataCenterBits) & CMaxDataCenter

	return NewIdWorkerEx(dataCenterId, workerId, twEpoch)
}

// return in ms
func (iw *IdWorker) timeGen() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (iw *IdWorker) timeReGen(last int64) int64 {
	ts := time.Now().UnixNano() / 1000 / 1000
	for {
		if ts <= last {
			ts = iw.timeGen()
		} else {
			break
		}
	}
	return ts
}

// NewId Func: Generate next id
func (iw *IdWorker) NextId() (ts int64, err error) {
	iw.lock.Lock()
	defer iw.lock.Unlock()
	ts = iw.timeGen()
	if ts == iw.lastTimeStamp {
		iw.sequence = (iw.sequence + 1) & CSequenceMask
		if iw.sequence == 0 {
			ts = iw.timeReGen(ts)
		}
	} else {
		iw.sequence = 0
	}

	if ts < iw.lastTimeStamp {
		err = errors.New("clock moved backwards, Refuse gen id")
		return -1, err
	}
	iw.lastTimeStamp = ts
	ts = (ts-CEpoch)<<CTimeStampShift | iw.dataCenterId<<CDataCenterShift | iw.workerId<<CWorkerIdShift | iw.sequence
	return ts, nil
}

// ParseId Func: reverse uuid to timestamp, dataCenterId, workId, seq
func ParseId(id int64) (t time.Time, ts int64, dataCenterId int64, workerId int64, seq int64) {
	seq = id & CSequenceMask
	workerId = (id >> CWorkerIdShift) & CMaxWorker
	dataCenterId = (id >> CDataCenterShift) & CMaxDataCenter
	ts = (id >> CTimeStampShift) + CEpoch
	t = time.Unix(ts/1000, (ts%1000)*1000000)
	return
}
