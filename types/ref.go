package types

import (
	"strings"
)

//Ref is a type to encode XSD ST_Ref, a reference that identifies a cell or a range of cells. E.g.: N28 or B5:N10
type Ref string

//ToCellRefs returns from/to CellRef of Ref
func (r Ref) ToCellRefs() (CellRef, CellRef) {
	cellRefs := strings.Split(string(r), ":")

	var from, to CellRef

	if len(cellRefs) == 1 {
		from = CellRef("A1")
		to = CellRef(cellRefs[0])
	} else {
		from = CellRef(cellRefs[0])
		to = CellRef(cellRefs[1])
	}

	return from, to
}

//ReboundIfRequired fix ref if required. E.g.: C1:B3 to B1:C3.
func (r Ref) ReboundIfRequired() Ref {
	fromCellRef, toCellRef := r.ToCellRefs()
	fromCol, fromRow := fromCellRef.ToIndexes()
	toCol, toRow := toCellRef.ToIndexes()

	if fromCol > toCol {
		toCol, fromCol = fromCol, toCol
	}

	if fromRow > toRow {
		toRow, fromRow = fromRow, toRow
	}

	return RefFromCellRefs(CellRefFromIndexes(fromCol, fromRow), CellRefFromIndexes(toCol, toRow))
}

//RefFromCellRefs returns Ref for from/to CellRefs
func RefFromCellRefs(from CellRef, to CellRef) Ref {
	return Ref(string(from) + ":" + string(to))
}