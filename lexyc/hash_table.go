package lexyc

import (
	"fmt"
	"go-custom-compiler/models"

	"github.com/golang-collections/collections/stack"
)

//HashTable ...
type HashTable struct {
	LineIndex  int
	LabelIndex int

	Lines  []string
	Labels map[string]string

	ActiveLabels *stack.Stack
	CurrentOp    string
	CurrentBlock string
	Statements   int
}

//NewHashTable ...
func NewHashTable() (*HashTable, error) {
	return &HashTable{Labels: make(map[string]string), ActiveLabels: stack.New()}, nil
}

//GetLine ...
func (h *HashTable) GetLine() string {
	return fmt.Sprintf("%v", h.LineIndex+1)
}

//GetLabel ...
func (h *HashTable) GetLabel() string {
	return fmt.Sprintf("_E%v", h.LabelIndex)
}

//GetPreviousLabel ...
func (h *HashTable) GetPreviousLabel() string {
	return fmt.Sprintf("_E%v", h.LabelIndex-1)
}

//GetNextLine ...
func (h *HashTable) GetNextLine() string {
	h.LineIndex++
	return fmt.Sprintf("%v", h.LineIndex)
}

//GetNextLabel ...
func (h *HashTable) GetNextLabel() string {
	h.LabelIndex++
	return fmt.Sprintf("_E%v", h.LabelIndex)
}

//AddNextLine ...
func (h *HashTable) AddNextLine(line string) {
	h.Lines = append(h.Lines, h.GetNextLine()+" "+line)
}

//AddNextLabel ...
func (h *HashTable) AddNextLabel(line string) {
	h.Labels[h.GetNextLabel()] = line
}

//PopLabelInLine ...
func (h *HashTable) PopLabelInLine() {
	h.Labels[h.ActiveLabels.Pop().(string)] = h.GetLine()
}

//AddLabelInLine ...
func (h *HashTable) AddLabelInLine() {
	h.Labels[h.GetLabel()] = h.GetLine()
}

//AddPreviousLabelInLine ...
func (h *HashTable) AddPreviousLabelInLine() {
	h.Labels[h.GetPreviousLabel()] = h.GetLine()
}

//AddNextOp ...
func (h *HashTable) AddNextOp() {
	h.Lines = append(h.Lines, h.GetNextLine()+" "+h.CurrentOp)
}

//AddNextBlock ...
func (h *HashTable) AddNextBlock() {
	h.Lines = append(h.Lines, h.GetNextLine()+" "+h.CurrentBlock)
}

//GetOperationFromOperator ...
func (h *HashTable) GetOperationFromOperator(operator string) string {
	return fmt.Sprintf("OPR 0, %v", models.ObjectCodeOperations[operator])
}
