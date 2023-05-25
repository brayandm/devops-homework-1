pipeline {
    agent any

    tools {
        go 'go'
    }

    environment {
        GITHUB_TOKEN = credentials('Github')
    }

    stages {

        stage('Build') {
            steps {
                sh('go build -o app main.go')
            }
        }

        stage('Archive Artifact') {
            steps {
                archiveArtifacts(artifacts: 'app', onlyIfSuccessful: true)
            }
        }

        stage('Test') {
            steps {
                sh('go mod init app')
                sh('go test .')
            }
        }
    }
}
