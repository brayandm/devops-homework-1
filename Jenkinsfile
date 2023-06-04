pipeline {
    agent any

    tools {
        go 'go'
    }

    stages {
        stage('Cleaning') {
            steps {
                sh('rm -rf app || true')
                sh('rm -rf go.mod || true')
            }
        }

        stage('Test') {
            steps {
                sh('go mod init app')
                sh('go test .')
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

        stage('Deploy to Stage') {
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'target-ssh-credentials', keyFileVariable: 'KeyFile', usernameVariable: 'userName')]) {
                    sh "ssh-keyscan 192.168.105.3 > ~/.ssh/known_hosts"

                    sh "ssh -l ${userName} -i ${KeyFile} 192.168.105.3 -C sudo systemctl stop myapp || true"

                    sh "scp -i ${KeyFile} app ${userName}@192.168.105.3:"
                    sh "scp -i ${KeyFile} myapp.service ${userName}@192.168.105.3:"

                    sh "ssh -l ${userName} -i ${KeyFile} 192.168.105.3 -C sudo mv myapp.service /etc/systemd/system/"
                    sh "ssh -l ${userName} -i ${KeyFile} 192.168.105.3 -C sudo systemctl daemon-reload"
                    sh "ssh -l ${userName} -i ${KeyFile} 192.168.105.3 -C sudo systemctl start myapp"
                    sh "ssh -l ${userName} -i ${KeyFile} 192.168.105.3 -C sudo systemctl enable myapp"
                }
            }
        }

        stage('Deploy to Production') {
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'ec2-ssh-credentials', keyFileVariable: 'KeyFile', usernameVariable: 'userName')]) {
                    sh "ssh-keyscan 13.53.131.172 > ~/.ssh/known_hosts"

                    sh "ssh -l ${userName} -i ${KeyFile} 13.53.131.172 -C sudo systemctl stop myappProd || true"

                    sh "scp -i ${KeyFile} app ${userName}@13.53.131.172:"
                    sh "scp -i ${KeyFile} myappProd.service ${userName}@13.53.131.172:"

                    sh "ssh -l ${userName} -i ${KeyFile} 13.53.131.172 -C sudo mv myappProd.service /etc/systemd/system/"
                    sh "ssh -l ${userName} -i ${KeyFile} 13.53.131.172 -C sudo systemctl daemon-reload"
                    sh "ssh -l ${userName} -i ${KeyFile} 13.53.131.172 -C sudo systemctl start myappProd"
                    sh "ssh -l ${userName} -i ${KeyFile} 13.53.131.172 -C sudo systemctl enable myappProd"
                }
            }
        }
    }
}
