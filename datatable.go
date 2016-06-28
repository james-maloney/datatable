package datatable

import (
	"fmt"
	"time"
)

type Table struct {
	Count int `json:"count,omitempty"`

	Columns []*Column              `json:"cols"`
	Rows    []*Row                 `json:"rows"`
	P       map[string]interface{} `json:"p,omitempty"`
}

type Row struct {
	Cells []*Cell `json:"c"`
}

type Cell struct {
	Value  interface{} `json:"v"`
	Format string      `json:"f,omitempty"`
}

type ColumnType string

var (
	String    ColumnType = "string"
	Number    ColumnType = "number"
	Bool      ColumnType = "boolean"
	Date      ColumnType = "date"
	DateTime  ColumnType = "datetime"
	TimeOfDay ColumnType = "timeofday"
)

type Role string

var (
	Annotation     Role = "annotation"
	AnnotationText Role = "annotationText"
	Certinty       Role = "certinty"
	Emphasis       Role = "emphasis"
	Interval       Role = "interval"
	Scope          Role = "scope"
	ToolTip        Role = "tooltip"
	Domain         Role = "domain"
	Data           Role = "data"
)

var RoleType = map[Role]ColumnType{
	Annotation:     String,
	AnnotationText: String,
	Certinty:       Bool,
	Emphasis:       Bool,
	Interval:       Number,
	Scope:          Bool,
	ToolTip:        String,
}

type Column struct {
	Type    ColumnType `json:"type"`              // Required
	Label   string     `json:"label,omitempty"`   // Optional - Used for display
	ID      string     `json:"id,omitempty"`      // Optional
	Role    Role       `json:"role,omitempty"`    // Optional
	Pattern string     `json:"pattern,omitempty"` // Optional - String pattern that was used by a data source to format numeric, date, or time column values
}

func NewTable() *Table {
	return &Table{
		Columns: []*Column{},
		Rows:    []*Row{},
	}
}

func (t *Table) AddColumn(c *Column) *Table {
	if t.Columns == nil {
		t.Columns = []*Column{}
	}
	if t.Rows == nil {
		t.Rows = []*Row{}
	}

	if c.Role != "" {
		t, ok := RoleType[c.Role]
		if ok {
			c.Type = t
		}
	}

	if len(c.Type) == 0 {
		c.Type = String
	}

	t.Columns = append(t.Columns, c)

	return t
}

// FormatDate this is the format that the Google Visualization Library takes
func FormatDate(t time.Time) string {
	return fmt.Sprintf("Date(%v,%v,%v,%v,%v)", t.Year(), int(t.Month())-1, t.Day(), t.Hour(), t.Minute())
}

func (t *Table) AddRow(items []*Cell) {
	if t.Rows == nil {
		t.Rows = []*Row{}
	}

	r := &Row{}
	r.Cells = items

	t.Rows = append(t.Rows, r)
}
