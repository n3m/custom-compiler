package helpers

import "go-custom-compiler/models"

//QueueContainsBlock ...
func QueueContainsBlock(queue []models.BlockType, block models.BlockType) bool {
	for _, a := range queue {
		if a == block {
			return true
		}
	}
	return false
}

//RemoveFromQueue ...
func RemoveFromQueue(queue []models.BlockType, block models.BlockType) ([]models.BlockType, bool) {
	for i, v := range queue {
		if v == block {
			return append(queue[:i], queue[i+1:]...), true
		}
	}
	return queue, false
}
