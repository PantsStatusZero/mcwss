package event

// BookEdited is sent by the client when it edits a book and closes the entire display. This means that this
// event is sent as soon as a player clicks the cross in the top right corner or signs the book.
type BookEdited struct {
	// Sequence is the sequence in which the book was edited.
	Sequence int `json:"Seq"`
	// Type is the type of the book after it was edited. If the book was signed, this means that the type is
	// 387 (written book), whereas if the book was only edited, the type is 386. (writable book)
	Type int
	// NetworkType no longer appears to be used. It is always 0.
	NetworkType int
	// PageCount is the amount of pages that the book had after editing.
	PageCount int
}
