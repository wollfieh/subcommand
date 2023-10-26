package med

import (
	"fmt"

	"github.com/golang-module/carbon"
)

type Med struct {
	name      string
	count     int
	date      carbon.Carbon
	rate      int
	threshold int
}

// NewMed like this is pretty useless. I can do foo:=&Med{ instead of calling this?
// TODO what might be cool would be to have standard values?
func NewMed(name string, count int, date carbon.Carbon, rate int, threshold int) *Med {
	m := Med{
		name:      name,
		count:     count,
		date:      date,
		rate:      rate,
		threshold: threshold,
	}

	m.update()

	return &m
}

// updates the count of med to current day
func (m *Med) update() {
	// calculate the number of days since last update

	tdiff := carbon.Now().DiffAbsInDays(m.date)

	m.date = m.date.AddDays(int(tdiff))
	m.count -= int(tdiff) * m.rate // TODO: check for negative count

}
func (m *Med) String() string {
	return fmt.Sprintf("%s count %d/%d on %s. ", m.name, m.count, m.rate, m.date)
}

// Calc calculates when (days from now) the med will be available for threshold days
// e.g.
// threshold 0 -> how many days until med runs out
// threshold 5 -> how many days until I have 5 days left
func (m *Med) Calc(threshold int) int {
	return m.count/m.rate - threshold
}
func (m *Med) Ends() carbon.Carbon {
	days := m.Calc(0)

	enddate := m.date.AddDays(days)
	return enddate
}

func (m *Med) Monday() carbon.Carbon {
	return m.Ends().SetWeekStartsAt(carbon.Monday).StartOfWeek()
}

// Alert if one med Calc threshold hits zero
func (m *Med) Alert() bool {
	t := m.Calc(m.threshold)
	return t <= 0
}

func (m *Med) Dump() {
	fmt.Printf("%s (%d) ends in %d days on %s\nMonday is %s\n",
		m.name,
		m.count,
		m.Calc(0),
		m.Ends().ToUnixDateString(), m.Monday().ToUnixDateString())

}
