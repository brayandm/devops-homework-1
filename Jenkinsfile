pipeline {
    agent any

    tools {
        go 'go'
    }

    environment {
        GITHUB_TOKEN = credentials('Github')
    }

    stages {
        stage('Clone Repository') {
            steps {
                sh('rm -rf devops-homework-1 && git clone https://${GITHUB_TOKEN}@github.com/brayandm/devops-homework-1.git')
            }
        }

        stage('Build') {
            steps {
                dir('devops-homework-1') {
                    sh('go build -o app main.go')
                }
            }
        }

        stage('Archive Artifact') {
            steps {
                archiveArtifacts(artifacts: 'devops-homework-1/app', onlyIfSuccessful: true)
            }
        }

        stage('Test') {
            steps {
                dir('devops-homework-1') {
                    sh('go mod init app')
                    sh('go test .')
                }
            }
        }
    }
}
