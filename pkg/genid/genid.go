package genid

import (
	"context"
	"errors"
	"sync"

	"gitlab.com/vk-golang/lectures/05_web_app/99_hw/redditclone/pkg/strrand"
)

const (
	maxIter  = 1000
	stdIDLen = 20
)

var (
	stdContext         = context.Background()
	stdGenerator       = NewIDGenerator(stdContext, stdIDLen)
	stdTracker         = NewIDTracker(stdContext)
	stdGeneratorFacade = NewIDGeneratorFacade(stdContext, stdGenerator, stdTracker)
)

type idGenerator struct {
	idLen uint64
}

func NewIDGenerator(_ context.Context, idLen uint64) *idGenerator {
	return &idGenerator{
		idLen: idLen,
	}
}

func (gen *idGenerator) Generate(_ context.Context) string {
	id := strrand.RandStringBytes(gen.idLen)
	return id
}

type idTracker struct {
	ids map[string]bool
}

func NewIDTracker(_ context.Context) *idTracker {
	return &idTracker{
		ids: make(map[string]bool),
	}
}

func (tracker *idTracker) Check(_ context.Context, id string) bool {
	_, exists := tracker.ids[id]
	return exists
}

func (tracker *idTracker) Track(_ context.Context, id string) {
	tracker.ids[id] = true

}

// Facade + thread-safety
type idGeneratorFacade struct {
	mu        sync.Mutex
	generator *idGenerator
	tracker   *idTracker
}

func NewIDGeneratorFacade(_ context.Context, gen *idGenerator, track *idTracker) *idGeneratorFacade {
	return &idGeneratorFacade{
		mu:        sync.Mutex{},
		generator: gen,
		tracker:   track,
	}
}

func (genTrack *idGeneratorFacade) Generate(ctx context.Context) (string, error) {
	for i := 0; i < maxIter; i++ {
		id := genTrack.generator.Generate(ctx)
		genTrack.mu.Lock()

		if !genTrack.tracker.Check(ctx, id) {
			genTrack.tracker.Track(ctx, id)
			genTrack.mu.Unlock()
			return id, nil
		}
		genTrack.mu.Unlock()
	}
	return "", errors.New("error")
}

func Generate(ctx context.Context) (string, error) {

	id, err := stdGeneratorFacade.Generate(ctx)
	if err == nil {
		return id, nil
	}

	return "", err
}
