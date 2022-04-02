# Plentymarkets-API
self written plentymarkets api client

## Generator manual configs

#### Before generating generating api code following steps need to changed in openapi.json
- rename field "Configuration" to "ApiConfiguration"

#### After generating api code following steps are need to made manual:
- rename SetId() function of ItemSetComponent struct to SetIdentifikator()
