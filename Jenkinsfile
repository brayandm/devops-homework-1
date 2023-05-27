pipeline {
    agent any

    tools {
        go 'go'
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

        stage('Cleaning') {
            steps {
                sh('rm -rf app')
                sh('rm -rf go.mod')
            }
        }

        stage('Deploy') {
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'target-ssh-credentials', keyFileVariable: 'KeyFile', usernameVariable: 'userName')]) {
                    sh "scp -i ${KeyFile} main ${userName}@192.168.105.3:"
                }
            }
        }
    }
}
