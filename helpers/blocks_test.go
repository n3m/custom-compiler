package helpers

import (
	"go-custom-compiler/models"
	"testing"
)

func TestRemoveBlocks(t *testing.T) {
	arr := []models.BlockType{models.NULLBLOCK, models.PROCEDUREBLOCK, models.FUNCTIONBLOCK}
	t.Errorf("BEFORE > %+v", arr)

	arr, deleted := RemoveFromQueue(arr, models.PROCEDUREBLOCK)
	if deleted {
		t.Errorf("> %+v", true)
		t.Errorf("AFTER > %+v", arr)
	} else {
		t.Logf("> %+v", false)
		t.Fatal()
	}
}
