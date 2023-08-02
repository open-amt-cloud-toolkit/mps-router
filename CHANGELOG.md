<a name="2.1.1"></a>
## [2.1.1] - 2022-10-05
### Build
- **deps:** bump github.com/lib/pq from 1.10.6 to 1.10.7 (#7e3f3a3) 
- **deps:** bump github.com/stretchr/testify from 1.7.1 to 1.8.0 (#bf8cd4f) 
- **deps:** bump github.com/lib/pq from 1.10.5 to 1.10.6 (#e4d6def) 
- **docker:** remove GOARCH build flag (#f590d82) 

### Ci
- add junit test output (#3233976) 

### Fix
- release updated dependencies (#f3e5830) 

### Test
- **proxy:** refactors parseguid to use table driven (#d59687b) 


<a name="v2.1.0"></a>
## [v2.1.0] - 2022-05-03
### Build
- **deps:** bump github.com/lib/pq from 1.10.4 to 1.10.5 (#5f91e2f) 
- **deps:** bump github.com/stretchr/testify from 1.7.0 to 1.7.1 (#795b20d) 
- **deps:** bump github.com/lib/pq from 1.10.1 to 1.10.4 (#adebc41) 

### Ci
- **jenkinsfile:** removes protex scan (#89c4ecc) 
- **lint:** adds semantic checks to PRs (#a42845c) 
- **release:** adds semantic release to repo (#891195c) 

### Docs
- **github:** add pull request template (#36ff9bf) 

### Feat
- **env:** add option to override default mps host (#0b5fdd9) 
- **healthcheck:** adds flag for checking db status (#b3360c7) 


<a name="v2.0.0"></a>
## [v2.0.0] - 2021-09-15
### Ci
- **changelog:** add automation (#4cc5126) 

### Docs
- **security:** added SECURITY.md file (#777962e) 
- **security:** added security.md file (#3d8be20) 


<a name="v1.4.0"></a>
## v1.4.0 - 2021-06-23
### Build
- **changelog:** add config (#625fb46) 
- **docker:** use non root user (#df60f17) 
- **scan:** fixed protex and checkmarx scan (#2b0313a) 
- **scan:** fixed MPS-Router checkmarx scan (#a9633cc) 
- **scans:** enabled checkmarx (#f34d530) 
- **scans:** enabled Checkmarx (#355035f) 

### Ci
- add jenkinsfile (#5b25763) 
- **changelog:** add automation for changelog generation (#6f6a354) 
- **jenkins:** fix protex project name (#7b3fa88) 

### Docs
- **changelog:** fix version (#e47c569) 
- **changelog:** add changelog (#7d622f1) 
- **copyright:** add missing header (#8d88116) 

### Feat
- **build:** added git workflows (#7ce7e80) 

### Fix
- guid parse now supports v1-4 (#393dc3f) 
- **proxy:** Updated mps server and mps router ports as env variables (#4bb454b) 
- **test:** added unit tests for db (#bd816f6) 

### Test
- **proxy:** test forward and backward functions (#d02b571) 


[Unreleased]: https://github.com/open-amt-cloud-toolkit/mps/compare/2.0.0...HEAD
[2.0.0]: https://github.com/open-amt-cloud-toolkit/mps/compare/v2.1.0...2.0.0
[v2.1.0]: https://github.com/open-amt-cloud-toolkit/mps/compare/v2.0.0...v2.1.0
[v2.0.0]: https://github.com/open-amt-cloud-toolkit/mps/compare/v1.4.0...v2.0.0
