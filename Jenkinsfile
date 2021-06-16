pipeline{
    agent {
        label 'docker-amt'
    }
    options {
        buildDiscarder(logRotator(numToKeepStr: '5', daysToKeepStr: '30'))
        timestamps()
        timeout(unit: 'HOURS', time: 2)
    }
    stages{
        stage('Cloning Repository') {
            steps{ 
                script{
                    scmCheckout {
                        clean = true
                    }
                }
            }
        }
        stage('Static Code Scan') {
            steps{
                script{
                    staticCodeScan {
                        // generic
                        scanners             = ['protex', 'snyk']
                        scannerType          = 'go'

                        protexProjectName    = 'OpenAMT - MPS Router'
                        protexBuildName      = 'rrs-generic-protex-build'

                        // checkmarxProjectName = "OpenAMT - MPS Router"

                        //snyk details
                        snykManifestFile        = ['go.mod']
                        snykProjectName         = ['openamt-mps-router']
                    }
                }
            }
        }
    }
}
