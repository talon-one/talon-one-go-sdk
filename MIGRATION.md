# Migration guide

This document provides guidance on how to migrate from the [legacy version](https://github.com/talon-one/talon_go)
of the SDK to the latest version. Follow the steps below to ensure a smooth transition.

## Summary

The new Go SDK includes numerous improvements under the hood. The most visible changes
involve function renaming to better reflect their functionality and align with Go naming
conventions.

The following changes were required to update our example code:

```diff
import (
-  talon "github.com/talon-one/talon_go/v25"
+  talon "github.com/talon-one/talon-one-go-sdk"
)

//...

- integrationState, _, err := integrationClient.IntegrationApi.
+ integrationState, _, err := integrationClient.IntegrationAPI.
        UpdateCustomerSessionV2(integrationAuthContext, "deetdoot_2").
-       Body(integrationRequest).
+       IntegrationRequest(integrationRequest).
        Execute()
```

As you can see, the migration requires minimal changes:

- The import path has been updated.
- `IntegrationApi` has been renamed to `IntegrationAPI`.
- The `Body` function has been renamed to `IntegrationRequest`.
- Similarly, `ManagementApi` has been renamed to `ManagementAPI`.

## Steps

Follow these steps to migrate your code from the legacy Go SDK to the new version:

1. Replace all imports of `github.com/talon-one/talon_go/v25` with `github.com/talon-one/talon-one-go-sdk`.
2. Run `go mod tidy && go mod vendor` to update your dependencies.
3. Review and address any compilation errors by updating function names according to the new naming conventions.
