package cron

import (
	"log"
	"runtime"
	"sort"
	"sync"
	"time"
)

var (
	entriesMutex = new(sync.RWMutex)
)

type Cron struct {
	entries  map[string]*Entry
	stop     chan struct{}
	add      chan *Entry
	snapshot chan []*Entry
	running  bool
	ErrorLog *log.Logger
	location *time.Location
}

type Job interface {
	Run()
}

type Schedule interface {
	Next(time.Time) time.Time
}

type Entry struct {
	Name     string
	Schedule Schedule
	Next     time.Time
	Prev     time.Time
	Job      Job
}

type byTime []*Entry

func (s byTime) Len() int {
	return len(s)
}
func (s byTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byTime) Less(i, j int) bool {
	if s[i].Next.IsZero() {
		return false
	}
	if s[j].Next.IsZero() {
		return true
	}
	return s[i].Next.Before(s[j].Next)
}

func New() *Cron {
	return NewWithLocation(time.Now().Location())
}

func NewWithLocation(location *time.Location) *Cron {
	return &Cron{
		entries:  make(map[string]*Entry),
		add:      make(chan *Entry),
		stop:     make(chan struct{}),
		snapshot: make(chan []*Entry),
		running:  false,
		ErrorLog: nil,
		location: location,
	}
}

type FuncJob func()

func (f FuncJob) Run() {
	f()
}

func (c *Cron) AddFunc(name, spec string, cmd func()) error {
	return c.AddJob(name, spec, FuncJob(cmd))
}

func (c *Cron) AddJob(name, spec string, cmd Job) error {
	schedule, err := Parse(spec)
	if err != nil {
		return err
	}
	c.Schedule(name, schedule, cmd)
	return nil
}

func (c *Cron) Remove(name string) {
	if _, ok := c.entries[name]; ok {
		delete(c.entries, name)
	}
}

func (c *Cron) Schedule(name string, schedule Schedule, cmd Job) {
	entry := &Entry{
		Name:     name,
		Schedule: schedule,
		Job:      cmd,
	}
	if !c.running {
		entriesMutex.Lock()
		c.entries[entry.Name] = entry
		entriesMutex.Unlock()
		return
	}

	c.add <- entry
}

func (c *Cron) Entries() []*Entry {
	if c.running {
		c.snapshot <- nil
		x := <-c.snapshot
		return x
	}
	return c.entrySnapshot()
}

func (c *Cron) Location() *time.Location {
	return c.location
}

func (c *Cron) Start() {
	if c.running {
		return
	}
	c.running = true
	go c.run()
}

func (c *Cron) Run() {
	if c.running {
		return
	}
	c.running = true
	c.run()
}

func (c *Cron) runWithRecovery(j Job) {
	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			c.logf("cron: panic running job: %v\n%s", r, buf)
		}
	}()
	j.Run()
}

func (c *Cron) run() {
	now := c.now()
	entriesMutex.RLock()
	for _, entry := range c.entries {
		entry.Next = entry.Schedule.Next(now)
	}
	entriesMutex.RUnlock()

	for {
		entries := c.getEntries()

		var timer *time.Timer
		if len(entries) == 0 {
			timer = time.NewTimer(100000 * time.Hour)
		} else {
			timer = time.NewTimer(entries[0].Next.Sub(now))
		}

		for {
			select {
			case now = <-timer.C:
				now = now.In(c.location)
				entries = c.getEntries()
				for _, e := range entries {
					if e.Next.After(now) || e.Next.IsZero() {
						break
					}
					go c.runWithRecovery(e.Job)
					e.Prev = e.Next
					e.Next = e.Schedule.Next(now)
				}

			case newEntry := <-c.add:
				timer.Stop()
				now = c.now()
				newEntry.Next = newEntry.Schedule.Next(now)
				entriesMutex.Lock()
				c.entries[newEntry.Name] = newEntry
				entriesMutex.Unlock()

			case <-c.snapshot:
				c.snapshot <- c.entrySnapshot()
				continue

			case <-c.stop:
				timer.Stop()
				return
			}

			break
		}
	}
}

func (c *Cron) logf(format string, args ...interface{}) {
	if c.ErrorLog != nil {
		c.ErrorLog.Printf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func (c *Cron) Stop() {
	if !c.running {
		return
	}
	c.stop <- struct{}{}
	c.running = false
}

func (c *Cron) entrySnapshot() []*Entry {
	return c.getEntries()
}

func (c *Cron) getEntries() []*Entry {
	sortedEntries := make([]*Entry, 0)
	entriesMutex.RLock()
	for _, e := range c.entries {
		sortedEntries = append(sortedEntries, e)
	}
	entriesMutex.RUnlock()
	sort.Sort(byTime(sortedEntries))
	return sortedEntries
}

func (c *Cron) now() time.Time {
	return time.Now().In(c.location)
}
