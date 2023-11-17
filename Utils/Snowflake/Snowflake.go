package snowflake

import (
	log "bluebell/Log"
	"fmt"
	"sync"
	"time"
)

const (
	DURATIONBITS     = 41
	DATACENTERIDBITS = 3
	WORKERIDBITS     = 7
	SEQUENCEBITS     = 12

	MAXDURATION     = int64(1<<DURATIONBITS - 1)
	MAXDATACENTERID = int64(1<<DATACENTERIDBITS - 1)
	MAXWORKERID     = int64(1<<WORKERIDBITS - 1)
	MAXSEQUENCE     = int64(1<<SEQUENCEBITS - 1)

	SHIFTTIMESTAMP    = SEQUENCEBITS + WORKERIDBITS + DATACENTERIDBITS
	SHIFTDATACENTERID = SEQUENCEBITS + WORKERIDBITS
	SHIFTWORKERID     = SEQUENCEBITS
	SHIFTSEQUENCE     = 0
)

type Snowflake struct {
	mutex sync.Mutex

	epoch        int64 // start time, record offset time in id[SHIFTTIMESTAMP + DURATIONBITS - 1, SHIFTTIMESTAMP]
	timestamp    int64
	datacenterid int64
	workerid     int64
	sequence     int64
}

func NewSnowflake(epoch time.Time, datacenterid, workerid int64) (*Snowflake, error) {
	if datacenterid < 0 || datacenterid > MAXDATACENTERID {
		err := fmt.Errorf("datacenterid %d is out of range", datacenterid)
		log.Panic(err.Error())
		return nil, err
	}
	if workerid < 0 || workerid > MAXWORKERID {
		err := fmt.Errorf("workerid %d is out of range", workerid)
		log.Panic(err.Error())
		return nil, err
	}

	return &Snowflake{
		mutex: sync.Mutex{},

		epoch:        epoch.UnixNano() / int64(time.Millisecond),
		timestamp:    time.Now().UnixNano() / int64(time.Millisecond),
		datacenterid: datacenterid,
		workerid:     workerid,
		sequence:     0,
	}, nil
}

func (s *Snowflake) NextID() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now().UnixNano() / int64(time.Millisecond)
	if now == s.timestamp {
		nseq := s.sequence + 1
		if nseq <= MAXSEQUENCE {
			s.sequence = nseq
		} else {
			for now == s.timestamp {
				now = time.Now().UnixNano() / int64(time.Millisecond)
			}
			s.sequence = 0
		}
	} else {
		s.sequence = 0
	}

	duration := now - s.epoch
	if duration > MAXDURATION {
		log.Panic("timestamp %d out of range %d", duration, MAXDURATION)
	}

	s.timestamp = now

	return (duration << SHIFTTIMESTAMP) + (s.datacenterid << SHIFTDATACENTERID) + (s.workerid << SHIFTWORKERID) + (s.sequence << SHIFTSEQUENCE)
}
