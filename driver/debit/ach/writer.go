package ach

import (
	"bufio"
	"io"
	"strings"
)

// Writer writes a File to an io.Writer.
// The File is validated against Nacha guidelines unless BypassValidation is enabled.
type Writer struct {
	w          *bufio.Writer
	lineNum    int    //current line being written
	LineEnding string // configurable line ending to support different consumer requirements
	// BypassValidation can be set to skip file validation and will allow non-compliant Nacha files to be written.
	BypassValidation bool
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w:          bufio.NewWriter(w),
		LineEnding: "\n", //set default line ending
	}
}

// Writer writes a single ach.file record to w
func (w *Writer) Write(file *File) error {
	if !w.BypassValidation {
		if err := file.Validate(); err != nil {
			return err
		}
	}

	w.lineNum = 0
	// Iterate over all records in the file
	if _, err := w.w.WriteString(file.Header.String() + w.LineEnding); err != nil {
		return err
	}
	w.lineNum++

	isADV := file.IsADV()

	if err := w.writeBatch(file, isADV); err != nil {
		return err
	}

	if err := w.writeIATBatch(file); err != nil {
		return err
	}

	if !isADV {
		if _, err := w.w.WriteString(file.Control.String() + w.LineEnding); err != nil {
			return err
		}
	} else {
		if _, err := w.w.WriteString(file.ADVControl.String() + w.LineEnding); err != nil {
			return err
		}
	}
	w.lineNum++

	// pad the final block
	for i := 0; i < (10-(w.lineNum%10)) && w.lineNum%10 != 0; i++ {
		if _, err := w.w.WriteString(strings.Repeat("9", 94) + w.LineEnding); err != nil {
			return err
		}
	}

	return w.w.Flush()
}

// Flush writes any buffered data to the underlying io.Writer.
func (w *Writer) Flush() error {
	return w.w.Flush()
}

func (w *Writer) writeBatch(file *File, isADV bool) error {
	for _, batch := range file.Batches {
		if _, err := w.w.WriteString(batch.GetHeader().String() + w.LineEnding); err != nil {
			return err
		}
		w.lineNum++
		if !isADV {
			for _, entry := range batch.GetEntries() {
				if _, err := w.w.WriteString(entry.String() + w.LineEnding); err != nil {
					return err
				}
				w.lineNum++

				if entry.Addenda02 != nil {
					if _, err := w.w.WriteString(entry.Addenda02.String() + w.LineEnding); err != nil {
						return err
					}
					w.lineNum++
				}
				for _, addenda05 := range entry.Addenda05 {
					if _, err := w.w.WriteString(addenda05.String() + w.LineEnding); err != nil {
						return err
					}
					w.lineNum++
				}
				if entry.Addenda98 != nil {
					if _, err := w.w.WriteString(entry.Addenda98.String() + w.LineEnding); err != nil {
						return err
					}
					w.lineNum++
				}
				if entry.Addenda99 != nil {
					if _, err := w.w.WriteString(entry.Addenda99.String() + w.LineEnding); err != nil {
						return err
					}
					w.lineNum++
				}
				if entry.Addenda99Dishonored != nil {
					if _, err := w.w.WriteString(entry.Addenda99Dishonored.String() + w.LineEnding); err != nil {
						return err
					}
					w.lineNum++
				}
				if entry.Addenda99Contested != nil {
					if _, err := w.w.WriteString(entry.Addenda99Contested.String() + w.LineEnding); err != nil {
						return err
					}
					w.lineNum++
				}
			}
		} else {
			for _, entry := range batch.GetADVEntries() {
				if _, err := w.w.WriteString(entry.String() + w.LineEnding); err != nil {
					return err
				}
				w.lineNum++
				if entry.Addenda99 != nil {
					if _, err := w.w.WriteString(entry.Addenda99.String() + w.LineEnding); err != nil {
						return err
					}
					w.lineNum++
				}
			}
		}

		if batch.GetHeader().StandardEntryClassCode != ADV {
			if _, err := w.w.WriteString(batch.GetControl().String() + w.LineEnding); err != nil {
				return err
			}
		} else {
			if _, err := w.w.WriteString(batch.GetADVControl().String() + w.LineEnding); err != nil {
				return err
			}
		}
		w.lineNum++
	}
	return nil
}

func (w *Writer) writeIATBatch(file *File) error {
	for _, iatBatch := range file.IATBatches {
		if _, err := w.w.WriteString(iatBatch.GetHeader().String() + w.LineEnding); err != nil {
			return err
		}
		w.lineNum++
		for _, entry := range iatBatch.GetEntries() {
			if _, err := w.w.WriteString(entry.String() + w.LineEnding); err != nil {
				return err
			}
			w.lineNum++
			if _, err := w.w.WriteString(entry.Addenda10.String() + w.LineEnding); err != nil {
				return err
			}
			w.lineNum++
			if _, err := w.w.WriteString(entry.Addenda11.String() + w.LineEnding); err != nil {
				return err
			}
			w.lineNum++
			if _, err := w.w.WriteString(entry.Addenda12.String() + w.LineEnding); err != nil {
				return err
			}
			w.lineNum++
			if _, err := w.w.WriteString(entry.Addenda13.String() + w.LineEnding); err != nil {
				return err
			}
			w.lineNum++
			if _, err := w.w.WriteString(entry.Addenda14.String() + w.LineEnding); err != nil {
				return err
			}
			w.lineNum++
			if _, err := w.w.WriteString(entry.Addenda15.String() + w.LineEnding); err != nil {
				return err
			}
			w.lineNum++
			if _, err := w.w.WriteString(entry.Addenda16.String() + w.LineEnding); err != nil {
				return err
			}
			w.lineNum++
			// IAT Addenda17
			for _, addenda17 := range entry.Addenda17 {
				if _, err := w.w.WriteString(addenda17.String() + w.LineEnding); err != nil {
					return err
				}
				w.lineNum++
			}
			// IAT Addenda18
			for _, addenda18 := range entry.Addenda18 {
				if _, err := w.w.WriteString(addenda18.String() + w.LineEnding); err != nil {
					return err
				}
				w.lineNum++
			}
			if entry.Addenda98 != nil {
				if _, err := w.w.WriteString(entry.Addenda98.String() + w.LineEnding); err != nil {
					return err
				}
				w.lineNum++
			}
			if entry.Addenda99 != nil {
				if _, err := w.w.WriteString(entry.Addenda99.String() + w.LineEnding); err != nil {
					return err
				}
				w.lineNum++
			}
		}
		if _, err := w.w.WriteString(iatBatch.GetControl().String() + w.LineEnding); err != nil {
			return err
		}
		w.lineNum++
	}
	return nil
}
