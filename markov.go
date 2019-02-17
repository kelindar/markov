package markov

import (
	"math/rand"
	"sort"
	"strings"
)

// Builder represents a builder for a Markov model.
type Builder struct {
	order   int
	start   map[string]int64
	product map[string]map[string]int64
}

// New constructs a new Markov model builder.
func New(order int) *Builder {
	return &Builder{
		order:   order,
		start:   make(map[string]int64),
		product: make(map[string]map[string]int64),
	}
}

// Build builds the model which can be used to generate strings.
func (b *Builder) Build() *Model {
	product := make(map[string][]state)
	for k, v := range b.product {
		product[k] = normalize(v)
	}

	return newModel(b.order, normalize(b.start), product)
}

// Teach adds one or multiple observations to the builder.
func (b *Builder) Teach(observations ...string) {
	for _, v := range observations {
		b.teach(v)
	}
}

// Teach adds a single sample
func (b *Builder) teach(observation string) {
	observation = strings.ToLower(observation)

	// if the sample is shorter than the order, add a production that this sample leads to nil
	if len(observation) <= b.order {
		addOrUpdate(b.start, observation, 1, func(x int64) int64 { return x + 1 })
		b.increment(observation, "")
		return
	}

	// chomp string into "order" length parts, and the single letter which follows it
	for i := 0; i < len(observation)-b.order+1; i++ {
		key := observation[i : i+b.order]
		if i == 0 {
			addOrUpdate(b.start, key, 1, func(x int64) int64 { return x + 1 })
		}

		var sub string
		if i+b.order != len(observation) {
			sub = observation[i+b.order : i+b.order+1]
		}
		b.increment(key, sub)
	}
	return
}

// Increment adds one to the observation count.
func (b *Builder) increment(observation, v string) {
	set, found := b.product[observation]
	if !found {
		set = make(map[string]int64)
		b.product[observation] = set
	}

	addOrUpdate(set, v, 1, func(x int64) int64 { return x + 1 })
	return
}

// ------------------------------------------------------------------------------------

// Model ...
type Model struct {
	order   int
	start   []state
	product map[string][]state
}

// String with a weight
type state struct {
	String string
	Weight float64
}

// newModel constructs a new Markov model.
func newModel(order int, starting []state, product map[string][]state) *Model {
	return &Model{
		order:   order,
		start:   starting,
		product: product,
	}
}

// Generate generates a string.
func (m *Model) Generate(r *rand.Rand) (builder string) {
	last := weightedRandom(r, m.start)
	for last != "" {

		// extend the string
		builder += last
		if len(builder) < m.order {
			return
		}

		// key to use to find the next production
		key := builder[len(builder)-m.order:]
		if v, ok := m.product[key]; ok {
			last = weightedRandom(r, v)
			continue
		}
		return
	}
	return
}

// ------------------------------------------------------------------------------------

// WeightedRandom picks a state randomly, based on its weight.
func weightedRandom(r *rand.Rand, states []state) string {
	random := r.Float64()
	for i := 0; i < len(states); i++ {
		if random -= states[i].Weight; random <= 0 {
			return states[i].String
		}
	}

	return ""
}

// addOrUpdate adds or updates a state.
func addOrUpdate(state map[string]int64, k string, v int64, update func(int64) int64) {
	if existing, ok := state[k]; ok {
		state[k] = update(existing)
		return
	}

	state[k] = v
	return
}

// Create a normalized histogram from string counts
func normalize(stringCounts map[string]int64) []state {
	var total float64
	for _, count := range stringCounts {
		total += float64(count)
	}

	dist := make([]state, 0, len(stringCounts))
	for key, count := range stringCounts {
		dist = append(dist, state{
			String: key,
			Weight: float64(count) / total,
		})
	}

	// Sort is not necessary, but makes testing easier
	sort.SliceStable(dist, func(i, j int) bool { return dist[i].String < dist[j].String })
	return dist
}
