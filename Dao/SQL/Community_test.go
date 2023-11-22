package sql

import (
	model "bluebell/Model"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func generateCommunities(ncomm int) (community []model.Community) {
	// id should be positive... don't know why
	for i := 1; i <= ncomm; i++ {
		community = append(community, model.Community{
			ID:            int64(i),
			Name:          fmt.Sprintf("community%d", i),
			Introducation: fmt.Sprintf("community%d is a community", i),
			Create_time:   time.Now(),
			Update_time:   time.Now(),
		})
	}

	return
}

func TestCommunity1(t *testing.T) {
	ncomm := 100
	communities := generateCommunities(ncomm)

	// clear all records in table community if exist
	if db.Migrator().HasTable(&model.Community{}) {
		db.Migrator().DropTable(&model.Community{})
		db.Migrator().CreateTable(&model.Community{})
	}

	for i, community := range communities {
		log.Printf("Inserting community %d: %v", i, community)
		if err := CreateCommunity(&community); err != nil {
			t.Error(err.Error())
			return
		}
	}

	for _, community := range communities {
		comm, err := GetCommunityByID(community.ID)
		if err != nil {
			t.Errorf("Community %v not exist", community)
			return
		}

		if comm.Name != community.Name || comm.Introducation != community.Introducation {
			t.Errorf("Community %v not equal", community)
			return
		}
	}

	comms, err := GetCommunities()
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(comms) != ncomm {
		t.Errorf("GetCommunities fail, comms: %v", comms)
		return
	}

	for _, community := range comms {
		comm, err := GetCommunityByID(community.ID)
		if err != nil {
			t.Errorf("Community %v not exist", community)
			return
		}

		if comm.Name != community.Name || comm.Introducation != community.Introducation {
			t.Errorf("Community %v not equal", community)
			return
		}
	}

	// delete *.db
	os.Remove("bluebell.db")
	return
}
