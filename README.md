[![ci](https://github.com/kakeru-lab/progressive-delivery-playground/actions/workflows/ci.yml/badge.svg)](https://github.com/kakeru-lab/progressive-delivery-playground/actions/workflows/ci.yml)

# progressive-delivery-playground
Local-only Terraform starter (no cloud required).

✅ Canary release + automated rollback on local Kind cluster  
✅ GitHub Actions CI, Prometheus metrics, Argo Rollouts

## Quick demo
make up
make deploy
make canary-bad   # will auto-rollback
