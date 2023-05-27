pipeline {
    agent any

    tools {
        go 'go'
    }

    stages {

          stage('Cleaning1') {
            steps {
                sh('rm -rf app')
                sh('rm -rf go.mod')
            }
        }


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

        stage('Deploy') {
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'target-ssh-credentials', keyFileVariable: 'KeyFile', usernameVariable: 'userName')]) {
                    sh "ssh-keyscan 192.168.105.3 > ~/.ssh/known_hosts"

                    sh "ssh -l ${userName} -i ${KeyFile} 192.168.105.3 -C sudo systemctl stop myapp"

                    sh "scp -i ${KeyFile} app ${userName}@192.168.105.3:"
                    sh "scp -i ${KeyFile} myapp.service ${userName}@192.168.105.3:"

                    sh "ssh -l ${userName} -i ${KeyFile} 192.168.105.3 -C sudo mv myapp.service /etc/systemd/system/"
                    sh "ssh -l ${userName} -i ${KeyFile} 192.168.105.3 -C sudo systemctl daemon-reload"
                    sh "ssh -l ${userName} -i ${KeyFile} 192.168.105.3 -C sudo systemctl start myapp"
                    sh "ssh -l ${userName} -i ${KeyFile} 192.168.105.3 -C sudo systemctl enable myapp"
                }
            }
        }
        
        stage('Cleaning') {
            steps {
                sh('rm -rf app')
                sh('rm -rf go.mod')
            }
        }
    }
}
