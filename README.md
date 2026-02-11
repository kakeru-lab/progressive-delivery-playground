[![ci](https://github.com/kakeru-lab/progressive-delivery-playground/actions/workflows/ci.yml/badge.svg)](https://github.com/kakeru-lab/progressive-delivery-playground/actions/workflows/ci.yml)

CI: gofmt check + go vet + go test + docker build (on push/PR)

# progressive-delivery-playground

Local-only **Progressive Delivery** playground (no cloud required).

✅ Canary release + automated rollback on local Kind cluster  
✅ GitHub Actions CI + Argo Rollouts (+ Prometheus metrics in next step)

## What you can demo
- Canary (10% → 50% → 100%)
- Intentional failure → automated rollback
- CI checks (test/build)

## Quick demo
```bash
make up
make deploy
make canary-bad  # will auto-rollback
