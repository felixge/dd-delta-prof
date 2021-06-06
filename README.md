# dd-delta-prof

A sample app to test the dd-trace-go PR: [profiler: Implement Delta Profiles #842](https://github.com/DataDog/dd-trace-go/pull/842).

## Usage

```
# Install or upgrade Go (dd-delta-prof requires 1.16)
brew install go
brew upgrade go

# Install dd-delta-prof test app
go install github.com/felixge/dd-delta-prof@latest

# Set upload destination and API Key (agentless is enabled by default)
export DD_SITE=...
export DD_API_KEY=...
export DD_ENV=prod

# Run dd-delta-prof which does lots of allocations
dd-delta-prof

# Wait 2-3 minutes until profiles show up in DD or the following log message
# appears because the backend is not accepting the payload yet:

# Datadog Tracer v1.31.0 ERROR: Failed to upload profile: 400 Bad Request, 1 additional messages skipped (first occurrence: 06 Jun 21 16:58 CEST)
```
