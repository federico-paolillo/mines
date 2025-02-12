package gc

type Reaper struct {
	store Store
}

func NewReaper(store Store) *Reaper {
	return &Reaper{
		store,
	}
}

func (r *Reaper) Reap() ReapStats {
	var stats ReapStats

	for defendant := range r.store.All() {
		verdict := emitVerdict(defendant)

		switch verdict {
		case Ok:
			stats.Ok++
		case Expired:
			stats.Expired++
		case Completed:
			stats.Completed++
		}

		if verdictIsUnfavourable(verdict) {
			r.store.Delete(defendant.Id)
		}
	}

	return stats
}
