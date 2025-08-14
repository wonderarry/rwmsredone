package templates

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
	"gopkg.in/yaml.v3"
)

type FSProvider struct {
	root  string
	mu    sync.RWMutex
	cache map[domain.TemplateKey]domain.CompiledTemplate
}

func NewFSProvider(root string) *FSProvider {
	return &FSProvider{root: root, cache: make(map[domain.TemplateKey]domain.CompiledTemplate)}
}

func (p *FSProvider) List(ctx context.Context) ([]domain.TemplateKey, error) {
	entries, err := os.ReadDir(p.root)
	if err != nil {
		return nil, err
	}
	var keys []domain.TemplateKey
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		ext := strings.ToLower(filepath.Ext(name))
		if ext == ".yaml" || ext == ".yml" {
			key := strings.TrimSuffix(name, ext)
			keys = append(keys, domain.TemplateKey(key))
		}
	}
	return keys, nil
}

func (p *FSProvider) Load(ctx context.Context, key domain.TemplateKey) (domain.CompiledTemplate, error) {
	p.mu.RLock()
	if ct, ok := p.cache[key]; ok {
		p.mu.RUnlock()
		return ct, nil
	}
	p.mu.RUnlock()

	path := filepath.Join(p.root, fmt.Sprintf("%s.yaml", key))
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			path = filepath.Join(p.root, fmt.Sprintf("%s.yml", key))
			b, err = os.ReadFile(path)
			if err != nil {
				return domain.CompiledTemplate{}, err
			}
		} else {
			return domain.CompiledTemplate{}, err
		}
	}

	var spec rawSpec
	if err := yaml.Unmarshal(b, &spec); err != nil {
		return domain.CompiledTemplate{}, err
	}

	ct, err := compileSpec(spec)
	if err != nil {
		return domain.CompiledTemplate{}, err
	}
	ct.TemplateKey = key

	p.mu.Lock()
	p.cache[key] = ct
	p.mu.Unlock()

	return ct, nil
}

var _ contract.TemplateProvider = (*FSProvider)(nil)
