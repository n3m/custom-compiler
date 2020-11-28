package lexyc

import "fmt"

//HashTable ...
type HashTable struct {
	LineIndex  int
	LabelIndex int

	Lines  []string
	Labels map[string]string

	CurrentOp string
}

//NewHashTable ...
func NewHashTable() (*HashTable, error) {
	return &HashTable{Labels: make(map[string]string)}, nil
}

//GetLine ...
func (h *HashTable) GetLine() string {
	return fmt.Sprintf("%v", h.LineIndex+1)
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

//AddNextOp ...
func (h *HashTable) AddNextOp() {
	h.Lines = append(h.Lines, h.GetNextLine()+" "+h.CurrentOp)
}
