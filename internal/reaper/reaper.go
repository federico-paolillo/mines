package reaper

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

	toReap := make([]string, 0)

	// We first collect and THEN delete because .All() acquires a lock that will not release until finished
	// Calling .Delete while looping will dead-lock

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
			toReap = append(toReap, defendant.Id)
		}
	}

	r.store.Delete(toReap...)

	return stats
}
