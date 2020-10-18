package helpers

import (
	"go-custom-compiler/models"
)

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
	for i, j := 0, len(queue)-1; i < j; i, j = i+1, j-1 {
		queue[i], queue[j] = queue[j], queue[i]
	}

	for i, v := range queue {
		if v == block {
			queueFinal := append(queue[:i], queue[i+1:]...)

			for i, j := 0, len(queueFinal)-1; i < j; i, j = i+1, j-1 {
				queueFinal[i], queueFinal[j] = queueFinal[j], queueFinal[i]
			}
			return queueFinal, true
		}
	}
	return queue, false
}
