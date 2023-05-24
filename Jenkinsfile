pipeline {
    agent any

    environment {
        GITHUB_TOKEN = credentials('Github')
    }

    stages {
        stage('Clone Repository') {
            steps {
                sh('rm -rf go-test && git clone https://${GITHUB_TOKEN}@github.com/brayandm/devops-homework-1.git')
            }
        }
    }
}
